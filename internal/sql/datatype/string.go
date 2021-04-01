package datatype

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type String struct {
	value string
}

func NewString(v string) String {
	return String{value: v}
}

func (s String) Raw() interface{} {
	return s.value
}

func (s String) Compare(v sql.Value) (sql.CompareType, error) {
	switch value := v.Raw().(type) {
	case string:
		switch {
		case s.value < value:
			return sql.Less, nil
		case s.value > value:
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

func (s String) UnaryPlus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) UnaryMinus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Add(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return String{value: s.value + value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) Sub(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Mul(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Div(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Pow(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Mod(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Equal(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) NotEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value != value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) GreaterThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value > value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) LessThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value < value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value > value || s.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case string:
		return Boolean{value: s.value < value || s.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (s String) And(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (s String) Or(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}
