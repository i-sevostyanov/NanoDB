package datatype

import (
	"fmt"
	"math"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Integer struct {
	value int64
}

func NewInteger(v int64) Integer {
	return Integer{value: v}
}

func (i Integer) Raw() interface{} {
	return i.value
}

func (i Integer) DataType() sql.DataType {
	return sql.Integer
}

func (i Integer) Compare(v sql.Value) (sql.CompareType, error) {
	switch value := v.Raw().(type) {
	case int64:
		switch {
		case i.value < value:
			return sql.Less, nil
		case i.value > value:
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

func (i Integer) UnaryPlus() (sql.Value, error) {
	return Integer{value: i.value}, nil
}

func (i Integer) UnaryMinus() (sql.Value, error) {
	return Integer{value: -i.value}, nil
}

func (i Integer) Add(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Integer{value: i.value + value}, nil
	case float64:
		return Float{value: float64(i.value) + value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Sub(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Integer{value: i.value - value}, nil
	case float64:
		return Float{value: float64(i.value) - value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Mul(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Integer{value: i.value * value}, nil
	case float64:
		return Float{value: float64(i.value) * value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Div(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Integer{value: i.value / value}, nil
	case float64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: float64(i.value) / value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Pow(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Float{value: math.Pow(float64(i.value), float64(value))}, nil
	case float64:
		return Float{value: math.Pow(float64(i.value), value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Mod(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Integer{value: i.value % value}, nil
	case float64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: math.Mod(float64(i.value), value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) Equal(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value == value}, nil
	case float64:
		return Boolean{value: float64(i.value) == value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) NotEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value != value}, nil
	case float64:
		return Boolean{value: float64(i.value) != value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) GreaterThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value > value}, nil
	case float64:
		return Boolean{value: float64(i.value) > value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) LessThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value < value}, nil
	case float64:
		return Boolean{value: float64(i.value) < value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value >= value}, nil
	case float64:
		return Boolean{value: float64(i.value) >= value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case int64:
		return Boolean{value: i.value <= value}, nil
	case float64:
		return Boolean{value: float64(i.value) <= value}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("unexptected arg type: %T", value)
	}
}

func (i Integer) And(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (i Integer) Or(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}
