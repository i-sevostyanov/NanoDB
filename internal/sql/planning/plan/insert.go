package plan

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate mockgen -source=insert.go -destination ./insert_mock_test.go -package plan_test

type TableInserter interface {
	Insert(row sql.Row) error
}

type Insert struct {
	inserter TableInserter
	row      sql.Row
}

func NewInsert(inserter TableInserter, row sql.Row) *Insert {
	return &Insert{
		inserter: inserter,
		row:      row,
	}
}

func (i *Insert) RowIter() (sql.RowIter, error) {
	if err := i.inserter.Insert(i.row); err != nil {
		return nil, fmt.Errorf("failed to insert row: %w", err)
	}

	return sql.RowsIter(), nil
}
