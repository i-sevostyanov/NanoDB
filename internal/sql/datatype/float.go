package datatype

import (
	"fmt"
	"math"
	"strconv"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Float struct {
	value float64
}

func NewFloat(v float64) Float {
	return Float{value: v}
}

func (f Float) Raw() interface{} {
	return f.value
}

func (f Float) String() string {
	return strconv.FormatFloat(f.value, 'E', -1, 64)
}

func (f Float) DataType() sql.DataType {
	return sql.Float
}

func (f Float) Compare(v sql.Value) (sql.CompareType, error) {
	switch value := v.Raw().(type) {
	case float64:
		switch {
		case f.value < value:
			return sql.Less, nil
		case f.value > value:
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

func (f Float) UnaryPlus() (sql.Value, error) {
	return Float{value: f.value}, nil
}

func (f Float) UnaryMinus() (sql.Value, error) {
	return Float{value: -f.value}, nil
}

func (f Float) Add(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Float{value: f.value + value}, nil
	case int64:
		return Float{value: f.value + float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("add: unexptected arg type: %T", value)
	}
}

func (f Float) Sub(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Float{value: f.value - value}, nil
	case int64:
		return Float{value: f.value - float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("sub: unexptected arg type: %T", value)
	}
}

func (f Float) Mul(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Float{value: f.value * value}, nil
	case int64:
		return Float{value: f.value * float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("mul: unexptected arg type: %T", value)
	}
}

func (f Float) Div(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: f.value / value}, nil
	case int64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: f.value / float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("div: unexptected arg type: %T", value)
	}
}

func (f Float) Pow(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Float{value: math.Pow(f.value, value)}, nil
	case int64:
		return Float{value: math.Pow(f.value, float64(value))}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("pow: unexptected arg type: %T", value)
	}
}

func (f Float) Mod(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: math.Mod(f.value, value)}, nil
	case int64:
		if value == 0 {
			return nil, fmt.Errorf("division by zero")
		}

		return Float{value: math.Mod(f.value, float64(value))}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("mod: unexptected arg type: %T", value)
	}
}

func (f Float) Equal(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value == value}, nil
	case int64:
		return Boolean{value: f.value == float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("equal: unexptected arg type: %T", value)
	}
}

func (f Float) NotEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value != value}, nil
	case int64:
		return Boolean{value: f.value != float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("not equal: unexptected arg type: %T", value)
	}
}

func (f Float) GreaterThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value > value}, nil
	case int64:
		return Boolean{value: f.value > float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("greater-than: unexptected arg type: %T", value)
	}
}

func (f Float) LessThan(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value < value}, nil
	case int64:
		return Boolean{value: f.value < float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("less-than: unexptected arg type: %T", value)
	}
}

func (f Float) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value >= value}, nil
	case int64:
		return Boolean{value: f.value >= float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("greater-or-equal: unexptected arg type: %T", value)
	}
}

func (f Float) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case float64:
		return Boolean{value: f.value <= value}, nil
	case int64:
		return Boolean{value: f.value <= float64(value)}, nil
	case nil:
		return Null{}, nil
	default:
		return nil, fmt.Errorf("less-or-equal: unexptected arg type: %T", value)
	}
}

func (f Float) And(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("and: unsupported operation")
}

func (f Float) Or(_ sql.Value) (sql.Value, error) {
	return nil, fmt.Errorf("or: unsupported operation")
}
