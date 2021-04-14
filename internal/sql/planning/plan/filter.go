package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

type Filter struct {
	cond  expr.Node
	child Node
}

func NewFilter(cond expr.Node, child Node) *Filter {
	return &Filter{
		cond:  cond,
		child: child,
	}
}

func (f *Filter) RowIter() (sql.RowIter, error) {
	iter, err := f.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("failed to get row iter: %w", err)
	}

	iter = &filterIter{
		cond: f.cond,
		iter: iter,
	}

	return iter, nil
}

// filterIter is an iterator that filters another iterator and skips rows that don't match the given condition.
type filterIter struct {
	cond expr.Node
	iter sql.RowIter
}

func (i *filterIter) Next() (sql.Row, error) {
	for {
		row, err := i.iter.Next()
		switch {
		case errors.Is(err, io.EOF):
			return nil, err
		case err != nil:
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}

		value, err := i.cond.Eval(row)
		if err != nil {
			return nil, err
		}

		isTrue, ok := value.Raw().(bool)
		if !ok {
			return nil, fmt.Errorf("argument must be type boolean, not type %T", isTrue)
		}

		if isTrue {
			return row, nil
		}
	}
}

func (i *filterIter) Close() error {
	return i.iter.Close()
}
