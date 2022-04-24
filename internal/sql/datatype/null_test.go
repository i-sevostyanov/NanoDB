package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestNull_Raw(t *testing.T) {
	t.Parallel()

	null := datatype.NewNull()

	switch value := null.Raw().(type) {
	case nil:
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestNull_String(t *testing.T) {
	t.Parallel()

	n := datatype.NewNull()
	assert.Equal(t, "null", n.String())
}

func TestNull_DataType(t *testing.T) {
	t.Parallel()

	n := datatype.NewNull()
	assert.Equal(t, sql.Null, n.DataType())
}

func TestNull_Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    sql.Value
		b    sql.Value
		cmp  sql.CompareType
	}{
		{
			name: "null vs null",
			a:    datatype.NewNull(),
			b:    datatype.NewNull(),
			cmp:  sql.Equal,
		},
		{
			name: "null vs integer",
			a:    datatype.NewNull(),
			b:    datatype.NewInteger(10),
			cmp:  sql.Less,
		},
		{
			name: "null vs float",
			a:    datatype.NewNull(),
			b:    datatype.NewFloat(10),
			cmp:  sql.Less,
		},
		{
			name: "null vs text",
			a:    datatype.NewNull(),
			b:    datatype.NewText("xyz"),
			cmp:  sql.Less,
		},
		{
			name: "null vs boolean",
			a:    datatype.NewNull(),
			b:    datatype.NewBoolean(true),
			cmp:  sql.Less,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cmp, err := test.a.Compare(test.b)
			require.NoError(t, err)
			assert.Equal(t, test.cmp, cmp)
		})
	}
}

func TestNull_UnaryPlus(t *testing.T) {
	t.Parallel()

	a := datatype.NewNull()
	v, err := a.UnaryPlus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestNull_UnaryMinus(t *testing.T) {
	t.Parallel()

	a := datatype.NewNull()
	v, err := a.UnaryMinus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestNull_Add(t *testing.T) {
	t.Parallel()

	t.Run("null + null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null + integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null + float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null + text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Sub(t *testing.T) {
	t.Parallel()

	t.Run("null - null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Sub(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null - integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Sub(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null - float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Sub(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null - text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Sub(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Mul(t *testing.T) {
	t.Parallel()

	t.Run("null * null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Mul(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null * integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Mul(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null * float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Mul(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null * text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Mul(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Div(t *testing.T) {
	t.Parallel()

	t.Run("null / null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Div(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null / integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Div(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null / float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Div(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null / text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Div(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Pow(t *testing.T) {
	t.Parallel()

	t.Run("null ^ null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Pow(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null ^ integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Pow(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null ^ float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Pow(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null ^ text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Pow(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Mod(t *testing.T) {
	t.Parallel()

	t.Run("null % null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Mod(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null % integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Mod(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null % float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Mod(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null % text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Mod(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_Equal(t *testing.T) {
	t.Parallel()

	t.Run("null == null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Equal(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null == integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Equal(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null == float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Equal(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null == text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Equal(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_NotEqual(t *testing.T) {
	t.Parallel()

	t.Run("null != null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.NotEqual(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null != integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.NotEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null != float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.NotEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null != text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.NotEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_GreaterThan(t *testing.T) {
	t.Parallel()

	t.Run("null > null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.GreaterThan(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null > integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.GreaterThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null > float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.GreaterThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null > text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.GreaterThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_LessThan(t *testing.T) {
	t.Parallel()

	t.Run("null < null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.LessThan(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null < integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.LessThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null < float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.LessThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null < text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.LessThan(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_GreaterOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("null >= null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.GreaterOrEqual(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null >= integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.GreaterOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null >= float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.GreaterOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null >= text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.GreaterOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_LessOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("null <= null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.LessOrEqual(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null <= integer", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.LessOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null <= float", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.LessOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null <= text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.LessOrEqual(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})
}

func TestNull_And(t *testing.T) {
	t.Parallel()

	t.Run("null AND null -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.And(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null AND true -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewBoolean(true)

		value, err := a.And(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null AND false -> false", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewBoolean(false)

		value, err := a.And(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewBoolean(false), value)
	})

	t.Run("null AND integer -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.And(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null AND float -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.And(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null AND text -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.And(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}

func TestNull_Or(t *testing.T) {
	t.Parallel()

	t.Run("null OR null -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewNull()

		value, err := a.Or(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null OR true -> true", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewBoolean(true)

		value, err := a.Or(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewBoolean(true), value)
	})

	t.Run("null OR false -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewBoolean(false)

		value, err := a.Or(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("null OR integer -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewInteger(10)

		value, err := a.Or(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null OR float -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewFloat(10)

		value, err := a.Or(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("null OR text -> unsupported", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewNull()
		b := datatype.NewText("xyz")

		value, err := a.Or(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}
