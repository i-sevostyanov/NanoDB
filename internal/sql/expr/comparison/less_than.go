package comparison

import (
	"cmp"
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func LessThan(left, right sql.Value) (sql.Value, error) {
	if left.DataType() == sql.Null || right.DataType() == sql.Null {
		return datatype.NewNull(), nil
	}

	if left.DataType() == right.DataType() {
		switch left.DataType() {
		case sql.Boolean:
			return lessThanBool(left, right), nil
		case sql.Integer:
			return lessThan[int64](left, right), nil
		case sql.Float:
			return lessThan[float64](left, right), nil
		case sql.Text:
			return lessThan[string](left, right), nil
		case sql.Null:
			return datatype.NewNull(), nil
		}
	}

	if left.DataType() == sql.Integer && right.DataType() == sql.Float {
		lvalue := float64(left.Raw().(int64))
		rvalue := right.Raw().(float64)

		return datatype.NewBoolean(lvalue < rvalue), nil
	}

	if left.DataType() == sql.Float && right.DataType() == sql.Integer {
		lvalue := left.Raw().(float64)
		rvalue := float64(right.Raw().(int64))

		return datatype.NewBoolean(lvalue < rvalue), nil
	}

	return nil, fmt.Errorf("can't compare %T and %T values", left.Raw(), right.Raw())
}

func lessThanBool(left, right sql.Value) sql.Value {
	var (
		lvalue uint8
		rvalue uint8
	)

	if leftRaw := left.Raw().(bool); leftRaw {
		lvalue = 1
	}

	if rightRaw := right.Raw().(bool); rightRaw {
		rvalue = 1
	}

	return datatype.NewBoolean(lvalue < rvalue)
}

func lessThan[T cmp.Ordered](left, right sql.Value) sql.Value {
	lvalue := left.Raw().(T)
	rvalue := right.Raw().(T)

	return datatype.NewBoolean(lvalue < rvalue)
}
