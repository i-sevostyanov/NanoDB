package datatype

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Text struct {
	value string
}

func NewText(v string) Text {
	return Text{value: v}
}

func (t Text) Raw() interface{} {
	return t.value
}

func (t Text) String() string {
	return t.value
}

func (t Text) DataType() sql.DataType {
	return sql.Text
}

func (t Text) Compare(v sql.Value) (sql.CompareType, error) {
	switch value := v.Raw().(type) {
	case string:
		switch {
		case t.value < value:
			return sql.Less, nil
		case t.value > value:
			return sql.Greater, nil
		default:
			return sql.Equal, nil
		}
	case nil:
		return sql.Greater, nil
	default:
		return sql.Equal, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) UnaryPlus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) UnaryMinus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Add(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Text{value: t.value + value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) Sub(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Mul(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Div(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Pow(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Mod(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Equal(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) NotEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value != value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) GreaterThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value > value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) LessThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value < value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value > value || t.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: t.value < value || t.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (t Text) And(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (t Text) Or(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}
