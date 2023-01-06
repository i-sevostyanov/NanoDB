package datatype

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Text struct {
	value string
}

func NewText(v string) Text {
	return Text{value: v}
}

func (t Text) Raw() any {
	return t.value
}

func (t Text) String() string {
	return t.value
}

func (t Text) DataType() sql.DataType {
	return sql.Text
}
