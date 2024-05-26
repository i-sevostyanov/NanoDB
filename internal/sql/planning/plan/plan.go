package plan

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate go run go.uber.org/mock/mockgen -typed -source=plan.go -destination ./plan_mock.go -package plan

type Node interface {
	Columns() []string
	RowIter() (sql.RowIter, error)
}
