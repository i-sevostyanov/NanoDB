package sql

import (
	"fmt"
	"strconv"
)

type DataType uint

const (
	Null DataType = iota
	Integer
	Float
	Text
	Boolean
)

func (t DataType) String() string {
	switch t {
	case Integer:
		return "integer"
	case Float:
		return "float"
	case Text:
		return "text"
	case Boolean:
		return "boolean"
	case Null:
		return "null"
	default:
		return fmt.Sprintf("DataType<%s>", strconv.Itoa(int(t)))
	}
}
