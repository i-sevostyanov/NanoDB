package plan

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate go run go.uber.org/mock/mockgen -typed -source=insert.go -destination ./insert_mock_test.go -package plan_test

type TableInserter interface {
	Insert(key int64, row sql.Row) error
}

type Insert struct {
	inserter TableInserter
	key      int64
	row      sql.Row
}

func NewInsert(inserter TableInserter, key int64, row sql.Row) *Insert {
	return &Insert{
		inserter: inserter,
		key:      key,
		row:      row,
	}
}

func (i *Insert) Columns() []string {
	return nil
}

func (i *Insert) RowIter() (sql.RowIter, error) {
	if err := i.inserter.Insert(i.key, i.row); err != nil {
		return nil, fmt.Errorf("insert row: %w", err)
	}

	return sql.RowsIter(), nil
}
