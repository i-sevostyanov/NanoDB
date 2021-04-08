package plan

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Scan struct {
	table sql.Table
}

func NewScan(table sql.Table) *Scan {
	return &Scan{
		table: table,
	}
}

func (t *Scan) RowIter() (sql.RowIter, error) {
	return t.table.RowIter()
}
