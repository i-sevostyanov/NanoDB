package plan

import (
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Limit struct {
	child Node
	Value int64
}

func NewLimit(value int64, child Node) *Limit {
	return &Limit{
		child: child,
		Value: value,
	}
}

func (l *Limit) Columns() []string {
	return l.child.Columns()
}

func (l *Limit) RowIter() (sql.RowIter, error) {
	iter, err := l.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("get row iter: %w", err)
	}

	iter = &limitIter{
		limit: l.Value,
		iter:  iter,
	}

	return iter, nil
}

type limitIter struct {
	limit int64
	pos   int64
	iter  sql.RowIter
}

func (i *limitIter) Next() (sql.Row, error) {
	if i.pos >= i.limit {
		return nil, io.EOF
	}

	i.pos++

	row, err := i.iter.Next()
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (i *limitIter) Close() error {
	return i.iter.Close()
}
