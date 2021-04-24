package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate mockgen -source=delete.go -destination ./delete_mock_test.go -package plan_test

type RowDeleter interface {
	Delete(key int64) error
}

type Delete struct {
	deleter RowDeleter
	child   Node
	pkIndex uint8
}

func NewDelete(deleter RowDeleter, pkIndex uint8, child Node) *Delete {
	return &Delete{
		deleter: deleter,
		child:   child,
		pkIndex: pkIndex,
	}
}

func (d *Delete) Columns() []string {
	return nil
}

func (d *Delete) RowIter() (sql.RowIter, error) {
	iter, err := d.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("failed to get row iter: %w", err)
	}

	iter = &deleteIter{
		iter:    iter,
		deleter: d.deleter,
		pkIndex: d.pkIndex,
	}

	return iter, nil
}

type deleteIter struct {
	iter    sql.RowIter
	deleter RowDeleter
	pkIndex uint8
}

func (i *deleteIter) Next() (sql.Row, error) {
	for {
		row, err := i.iter.Next()
		switch {
		case errors.Is(err, io.EOF):
			return nil, err
		case err != nil:
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}

		key, ok := row[i.pkIndex].Raw().(int64)
		if !ok {
			return nil, fmt.Errorf("unsupported key type: %T", key)
		}

		if err = i.deleter.Delete(key); err != nil {
			return nil, fmt.Errorf("failed to delete row: %w", err)
		}
	}
}

func (i *deleteIter) Close() error {
	return i.iter.Close()
}
