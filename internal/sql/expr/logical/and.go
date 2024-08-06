package logical

import (
	"errors"
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func And(left, right sql.Value) (sql.Value, error) {
	if left.DataType() != right.DataType() || left.DataType() != sql.Boolean {
		return nil, fmt.Errorf("and: unsupported operation for %T and %T values", left.Raw(), right.Raw())
	}

	lvalue, ok := left.Raw().(bool)
	if !ok {
		return nil, errors.New("and: left operand should be bool")
	}

	rvalue, ok := right.Raw().(bool)
	if !ok {
		return nil, errors.New("and: right operand should be bool")
	}

	return datatype.NewBoolean(lvalue && rvalue), nil
}
