package datatype

import (
	"strconv"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Float struct {
	value float64
}

func NewFloat(v float64) Float {
	return Float{value: v}
}

func (f Float) Raw() any {
	return f.value
}

func (f Float) String() string {
	return strconv.FormatFloat(f.value, 'E', -1, 64)
}

func (f Float) DataType() sql.DataType {
	return sql.Float
}
