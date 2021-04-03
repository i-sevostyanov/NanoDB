package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Column struct {
	Position uint8
}

func (c Column) Eval(row sql.Row) (sql.Value, error) {
	if int(c.Position) > len(row) {
		return nil, fmt.Errorf("column position out of range: %d > %d", c.Position, len(row))
	}

	return row[c.Position], nil
}
