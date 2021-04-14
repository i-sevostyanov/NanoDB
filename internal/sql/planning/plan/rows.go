package plan

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Rows struct {
	rows []sql.Row
}

func NewRows(row ...sql.Row) Rows {
	return Rows{
		rows: row,
	}
}

func (e Rows) RowIter() (sql.RowIter, error) {
	return sql.RowsIter(e.rows...), nil
}
