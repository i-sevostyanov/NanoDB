package datatype

import (
	"fmt"
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

func (b Boolean) Compare(v sql.Value) (sql.CompareType, error) {
	switch value := v.Raw().(type) {
	case bool:
		x := b.toInt(b.value)
		y := b.toInt(value)

		switch {
		case x < y:
			return sql.Less, nil
		case x > y:
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

func (b Boolean) UnaryPlus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) UnaryMinus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Add(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Sub(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Mul(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Div(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Pow(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Mod(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (b Boolean) Equal(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		return Boolean{value: b.value == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("equal: unexptected arg type: %T", value)
	}
}

func (b Boolean) NotEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		return Boolean{value: b.value != value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("not-equal: unexptected arg type: %T", value)
	}
}

func (b Boolean) GreaterThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		x := b.toInt(b.value)
		y := b.toInt(value)

		return Boolean{value: x > y}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("greater-than: unexptected arg type: %T", value)
	}
}

func (b Boolean) LessThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		x := b.toInt(b.value)
		y := b.toInt(value)

		return Boolean{value: x < y}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("less-than: unexptected arg type: %T", value)
	}
}

func (b Boolean) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		x := b.toInt(b.value)
		y := b.toInt(value)

		return Boolean{value: x >= y}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("greater-or-equal: unexptected arg type: %T", value)
	}
}

func (b Boolean) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		x := b.toInt(b.value)
		y := b.toInt(value)

		return Boolean{value: x <= y}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("less-or-equal: unexptected arg type: %T", value)
	}
}

func (b Boolean) And(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		return Boolean{value: b.value && value}, nil
	case nil:
		if b.value {
			return Null{}, nil
		}

		return Boolean{}, nil
	default:
		return nil, fmt.Errorf("and: unexptected arg type: %T", value)
	}
}

func (b Boolean) Or(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case bool:
		return Boolean{value: b.value || value}, nil
	case nil:
		if b.value {
			return Boolean{value: true}, nil
		}

		return Null{}, nil
	default:
		return nil, fmt.Errorf("or: unexptected arg type: %T", value)
	}
}

func (b Boolean) toInt(value bool) uint {
	if value {
		return 1
	}

	return 0
}
