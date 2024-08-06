package math

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func UnaryPlus(value sql.Value) (sql.Value, error) {
	switch value.DataType() {
	case sql.Float:
		v := value.Raw().(float64)

		return datatype.NewFloat(v), nil
	case sql.Integer:
		v := value.Raw().(int64)

		return datatype.NewInteger(v), nil
	default:
		return nil, fmt.Errorf("unary-plus: unsupported operand %T", value.Raw())
	}
}
