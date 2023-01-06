package math

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func Add(left, right sql.Value) (sql.Value, error) {
	if left.DataType() == sql.Null || right.DataType() == sql.Null {
		return datatype.NewNull(), nil
	}

	if left.DataType() == right.DataType() {
		switch left.DataType() {
		case sql.Float:
			lvalue := left.Raw().(float64)
			rvalue := right.Raw().(float64)
			return datatype.NewFloat(lvalue + rvalue), nil
		case sql.Integer:
			lvalue := left.Raw().(int64)
			rvalue := right.Raw().(int64)
			return datatype.NewInteger(lvalue + rvalue), nil
		case sql.Text:
			lvalue := left.Raw().(string)
			rvalue := right.Raw().(string)
			return datatype.NewText(lvalue + rvalue), nil
		}
	}

	if left.DataType() == sql.Integer && right.DataType() == sql.Float {
		lvalue := float64(left.Raw().(int64))
		rvalue := right.Raw().(float64)

		return datatype.NewFloat(lvalue + rvalue), nil
	}

	if left.DataType() == sql.Float && right.DataType() == sql.Integer {
		lvalue := left.Raw().(float64)
		rvalue := float64(right.Raw().(int64))

		return datatype.NewFloat(lvalue + rvalue), nil
	}

	return nil, fmt.Errorf("add: unsupported operand %T and %T", left.Raw(), right.Raw())
}
