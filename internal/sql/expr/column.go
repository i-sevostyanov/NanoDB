package expr

import (
	"errors"
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Column struct {
	Name     string
	Position uint8
}

func (c Column) String() string {
	return c.Name
}

func (c Column) Eval(row sql.Row) (sql.Value, error) {
	if len(row) == 0 {
		return nil, errors.New("empty row")
	}

	if int(c.Position) > len(row) {
		return nil, fmt.Errorf("column position out of range: %d > %d", c.Position, len(row))
	}

	return row[c.Position], nil
}
