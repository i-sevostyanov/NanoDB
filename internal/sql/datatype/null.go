package datatype

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Null struct{}

func NewNull() Null {
	return Null{}
}

func (n Null) Raw() any {
	return nil
}

func (n Null) String() string {
	return "null"
}

func (n Null) DataType() sql.DataType {
	return sql.Null
}
