package memory

import (
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Table struct {
	name       string
	scheme     sql.Scheme
	keys       []int64
	rows       map[int64]sql.Row
	seq        *Sequence
	primaryKey sql.Column
}

func NewTable(name string, scheme sql.Scheme) *Table {
	var primaryKey sql.Column

	for column := range scheme {
		if scheme[column].PrimaryKey {
			primaryKey = scheme[column]
			break
		}
	}

	return &Table{
		name:       name,
		scheme:     scheme,
		seq:        &Sequence{},
		primaryKey: primaryKey,
		rows:       make(map[int64]sql.Row),
	}
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) Scheme() sql.Scheme {
	return t.scheme
}

func (t *Table) PrimaryKey() sql.Column {
	return t.primaryKey
}

func (t *Table) RowIter() (sql.RowIter, error) {
	rows := make([]sql.Row, 0, len(t.rows))

	for _, key := range t.keys {
		rows = append(rows, t.rows[key])
	}

	i := &iter{
		rows: rows,
	}

	return i, nil
}

func (t *Table) Sequence() sql.Sequence {
	return t.seq
}

func (t *Table) Insert(key int64, row sql.Row) error {
	if _, ok := t.rows[key]; ok {
		return fmt.Errorf("duplicate primary key: %d", key)
	}

	t.rows[key] = row
	t.keys = append(t.keys, key)

	if t.seq.Value() < key {
		t.seq.SetValue(key)
	}

	return nil
}

func (t *Table) Delete(key int64) error {
	if _, ok := t.rows[key]; !ok {
		return fmt.Errorf("row with key %d not found", key)
	}

	// O(n), fix it in the next release
	for index := range t.keys {
		if t.keys[index] == key {
			t.keys = append(t.keys[:index], t.keys[index+1:]...)
			break
		}
	}

	delete(t.rows, key)

	return nil
}

func (t *Table) Update(key int64, row sql.Row) error {
	_, ok := t.rows[key]
	if !ok {
		return fmt.Errorf("row with key %d not found", key)
	}

	t.rows[key] = row

	return nil
}

type iter struct {
	index int
	rows  []sql.Row
}

func (i *iter) Next() (sql.Row, error) {
	if i.index > len(i.rows)-1 {
		return nil, io.EOF
	}

	row := i.rows[i.index]
	i.index++

	return row, nil
}

func (i *iter) Close() error {
	i.rows = nil
	return nil
}
