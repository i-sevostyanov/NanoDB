package comparison

import (
	"cmp"
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

func Compare(left, right sql.Value) (sql.CompareType, error) {
	if left.DataType() == right.DataType() {
		switch left.DataType() {
		case sql.Null:
			return sql.Equal, nil
		case sql.Boolean:
			return compareBool(left, right)
		case sql.Integer:
			return compareOrdered[int64](left, right)
		case sql.Float:
			return compareOrdered[float64](left, right)
		case sql.Text:
			return compareOrdered[string](left, right)
		}
	}

	if left.DataType() == sql.Null {
		return sql.Less, nil
	}

	if right.DataType() == sql.Null {
		return sql.Greater, nil
	}

	return sql.Equal, fmt.Errorf("can't compare %T and %T values", left.Raw(), right.Raw())
}

func compareBool(left, right sql.Value) (sql.CompareType, error) {
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

	return sql.CompareType(cmp.Compare(lvalue, rvalue)), nil
}

func compareOrdered[T cmp.Ordered](left, right sql.Value) (sql.CompareType, error) {
	lvalue := left.Raw().(T)
	rvalue := right.Raw().(T)

	return sql.CompareType(cmp.Compare(lvalue, rvalue)), nil
}
