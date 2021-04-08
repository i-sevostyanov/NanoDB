package plan

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Empty struct{}

func NewEmpty() Empty {
	return Empty{}
}

func (e Empty) RowIter() (sql.RowIter, error) {
	return sql.RowsIter(), nil
}
