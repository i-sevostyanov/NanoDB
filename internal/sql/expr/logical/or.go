package logical

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func Or(left, right sql.Value) (sql.Value, error) {
	if left.DataType() != right.DataType() || left.DataType() != sql.Boolean {
		return nil, fmt.Errorf("and: unsupported operation for %T and %T values", left.Raw(), right.Raw())
	}

	lvalue := left.Raw().(bool)
	rvalue := right.Raw().(bool)

	return datatype.NewBoolean(lvalue || rvalue), nil
}
