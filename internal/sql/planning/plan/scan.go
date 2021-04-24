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

func (s *Scan) Columns() []string {
	scheme := s.table.Scheme()
	columns := make([]string, len(scheme))

	for name := range scheme {
		position := int(scheme[name].Position)
		columns[position] = name
	}

	return columns
}

func (s *Scan) RowIter() (sql.RowIter, error) {
	return s.table.RowIter()
}
