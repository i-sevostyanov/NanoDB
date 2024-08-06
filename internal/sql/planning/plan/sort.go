package plan

import (
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/comparison"
)

type Order uint8

const (
	Ascending Order = iota
	Descending
)

type Sort struct {
	columnPos uint8
	order     Order
	child     Node
}

func NewSort(columnPos uint8, order Order, child Node) *Sort {
	return &Sort{
		columnPos: columnPos,
		order:     order,
		child:     child,
	}
}

func (s *Sort) Columns() []string {
	return s.child.Columns()
}

func (s *Sort) RowIter() (sql.RowIter, error) {
	iter, err := s.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("get row iter: %w", err)
	}

	iter = &sortIter{
		columnPos: s.columnPos,
		order:     s.order,
		iter:      iter,
		index:     -1,
	}

	return iter, nil
}

type sortIter struct {
	columnPos uint8
	order     Order
	rows      []sql.Row
	iter      sql.RowIter
	index     int
}

func (i *sortIter) Next() (sql.Row, error) {
	if i.index == -1 {
		if err := i.sortRows(); err != nil {
			return nil, err
		}
	}

	if i.index >= len(i.rows)-1 {
		return nil, io.EOF
	}

	i.index++
	row := i.rows[i.index]

	return row, nil
}

func (i *sortIter) Close() error {
	i.rows = nil

	return i.iter.Close()
}

func (i *sortIter) sortRows() error {
	position := i.columnPos
	operator := sql.Equal

	switch i.order {
	case Ascending:
		operator = sql.Less
	case Descending:
		operator = sql.Greater
	default:
		return fmt.Errorf("unknown sort order: %v", i.order)
	}

	i.rows = make([]sql.Row, 0)

loop:
	for {
		row, err := i.iter.Next()
		switch {
		case errors.Is(err, io.EOF):
			break loop
		case err != nil:
			return fmt.Errorf("get next row: %w", err)
		default:
			i.rows = append(i.rows, row)
		}
	}

	sort.Slice(i.rows, func(x, y int) bool {
		a := i.rows[x][position]
		b := i.rows[y][position]
		c, _ := comparison.Compare(a, b)

		return c == operator
	})

	return nil
}
