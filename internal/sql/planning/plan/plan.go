package plan

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate mockgen -source=plan.go -destination ./plan_mock.go -package plan

type Node interface {
	RowIter() (sql.RowIter, error)
}
