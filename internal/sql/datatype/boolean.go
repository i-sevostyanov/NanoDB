package datatype

import (
	"strconv"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Boolean struct {
	value bool
}

func NewBoolean(v bool) Boolean {
	return Boolean{value: v}
}

func (b Boolean) Raw() any {
	return b.value
}

func (b Boolean) String() string {
	return strconv.FormatBool(b.value)
}

func (b Boolean) DataType() sql.DataType {
	return sql.Boolean
}
