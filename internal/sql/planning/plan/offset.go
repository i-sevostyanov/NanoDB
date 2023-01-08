package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

// Offset is a node that skips the first N rows.
type Offset struct {
	Offset int64
	child  Node
}

// NewOffset creates a new Offset node.
func NewOffset(n int64, child Node) *Offset {
	return &Offset{
		Offset: n,
		child:  child,
	}
}

func (o *Offset) Columns() []string {
	return o.child.Columns()
}

func (o *Offset) RowIter() (sql.RowIter, error) {
	iter, err := o.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("get row iter: %w", err)
	}

	iter = &offsetIter{
		skip: o.Offset,
		iter: iter,
	}

	return iter, nil
}

type offsetIter struct {
	skip int64
	iter sql.RowIter
}

func (i *offsetIter) Next() (sql.Row, error) {
	for i.skip > 0 {
		_, err := i.iter.Next()
		switch {
		case errors.Is(err, io.EOF):
			return nil, err
		case err != nil:
			return nil, fmt.Errorf("skip row: %w", err)
		default:
			i.skip--
		}
	}

	row, err := i.iter.Next()
	switch {
	case errors.Is(err, io.EOF):
		return nil, err
	case err != nil:
		return nil, fmt.Errorf("get next row: %w", err)
	default:
		return row, nil
	}
}

func (i *offsetIter) Close() error {
	return i.iter.Close()
}
