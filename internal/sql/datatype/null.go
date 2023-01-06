package datatype

import (
	"fmt"

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

func (n Null) Compare(v sql.Value) (sql.CompareType, error) {
	switch v.Raw().(type) {
	case nil:
		return sql.Equal, nil
	default:
		return sql.Less, nil
	}
}

func (n Null) UnaryPlus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (n Null) UnaryMinus() (sql.Value, error) {
	return nil, fmt.Errorf("unsupported operation")
}

func (n Null) Add(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Sub(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Mul(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Div(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Pow(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Mod(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) Equal(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) NotEqual(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) GreaterThan(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) LessThan(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) GreaterOrEqual(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) LessOrEqual(v sql.Value) (sql.Value, error) {
	switch v.Raw().(type) {
	case nil:
		return nil, fmt.Errorf("unsupported operation")
	default:
		return Null{}, nil
	}
}

func (n Null) And(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case nil:
		return Null{}, nil
	case bool:
		if value {
			return Null{}, nil
		}

		return Boolean{}, nil
	default:
		return nil, fmt.Errorf("unsupported operation")
	}
}

func (n Null) Or(v sql.Value) (sql.Value, error) {
	switch value := v.Raw().(type) {
	case nil:
		return Null{}, nil
	case bool:
		if value {
			return Boolean{value: true}, nil
		}

		return Null{}, nil
	default:
		return nil, fmt.Errorf("unsupported operation")
	}
}
