package math

import (
	"fmt"
	"math"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func Pow(left, right sql.Value) (sql.Value, error) {
	if left.DataType() == sql.Null || right.DataType() == sql.Null {
		return datatype.NewNull(), nil
	}

	if left.DataType() == right.DataType() {
		switch left.DataType() {
		case sql.Float:
			lvalue := left.Raw().(float64)
			rvalue := right.Raw().(float64)

			return datatype.NewFloat(math.Pow(lvalue, rvalue)), nil
		case sql.Integer:
			lvalue := left.Raw().(int64)
			rvalue := right.Raw().(int64)

			return datatype.NewFloat(math.Pow(float64(lvalue), float64(rvalue))), nil
		default:
		}
	}

	if left.DataType() == sql.Integer && right.DataType() == sql.Float {
		lvalue := float64(left.Raw().(int64))
		rvalue := right.Raw().(float64)

		return datatype.NewFloat(math.Pow(lvalue, rvalue)), nil
	}

	if left.DataType() == sql.Float && right.DataType() == sql.Integer {
		lvalue := left.Raw().(float64)
		rvalue := float64(right.Raw().(int64))

		return datatype.NewFloat(math.Pow(lvalue, rvalue)), nil
	}

	return nil, fmt.Errorf("pow: unsupported operand %T and %T", left.Raw(), right.Raw())
}
