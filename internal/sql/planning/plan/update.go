package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

//go:generate mockgen -source=update.go -destination ./update_mock_test.go -package plan_test

type RowUpdater interface {
	Update(key int64, row sql.Row) error
}

type Update struct {
	updater RowUpdater
	child   Node
	pkIndex uint8
	columns map[uint8]expr.Node
}

func NewUpdate(updater RowUpdater, pkIndex uint8, columns map[uint8]expr.Node, child Node) *Update {
	return &Update{
		updater: updater,
		child:   child,
		pkIndex: pkIndex,
		columns: columns,
	}
}

func (u *Update) RowIter() (sql.RowIter, error) {
	iter, err := u.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("failed to get child iter: %w", err)
	}

	iter = &updateIter{
		iter:    iter,
		updater: u.updater,
		pkIndex: u.pkIndex,
		columns: u.columns,
	}

	return iter, err
}

type updateIter struct {
	iter    sql.RowIter
	updater RowUpdater
	pkIndex uint8
	columns map[uint8]expr.Node
}

func (u *updateIter) Next() (sql.Row, error) {
	for {
		row, err := u.iter.Next()
		switch {
		case errors.Is(err, io.EOF):
			return nil, err
		case err != nil:
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}

		updatedRow := make(sql.Row, len(row))
		copy(updatedRow, row)

		for i, expression := range u.columns {
			if updatedRow[i], err = expression.Eval(row); err != nil {
				return nil, fmt.Errorf("failed to eval expr: %w", err)
			}
		}

		key, ok := row[u.pkIndex].Raw().(int64)
		if !ok {
			return nil, fmt.Errorf("unsupported key type: %T", key)
		}

		if err = u.updater.Update(key, updatedRow); err != nil {
			return nil, fmt.Errorf("failed to update row: %w", err)
		}
	}
}

func (u *updateIter) Close() error {
	return u.iter.Close()
}
