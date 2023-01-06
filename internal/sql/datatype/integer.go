package datatype

import (
	"strconv"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Integer struct {
	value int64
}

func NewInteger(v int64) Integer {
	return Integer{value: v}
}

func (i Integer) Raw() any {
	return i.value
}

func (i Integer) String() string {
	return strconv.FormatInt(i.value, 10)
}

func (i Integer) DataType() sql.DataType {
	return sql.Integer
}
