package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestBoolean_Raw(t *testing.T) {
	t.Parallel()

	fn := func(expected bool) {
		b := datatype.NewBoolean(expected)

		switch value := b.Raw().(type) {
		case bool:
			assert.Equal(t, expected, value)
		default:
			assert.Failf(t, "fail", "unexpected type %T", value)
		}
	}

	fn(true)
	fn(false)
}

func TestBoolean_String(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	assert.Equal(t, "true", b.String())

	b = datatype.NewBoolean(false)
	assert.Equal(t, "false", b.String())
}

func TestBoolean_DataType(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	assert.Equal(t, sql.Boolean, b.DataType())
}

func TestBoolean_Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    sql.Value
		b    sql.Value
		cmp  sql.CompareType
		err  bool
	}{
		{
			name: "true vs true",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewBoolean(true),
			cmp:  sql.Equal,
		},
		{
			name: "false vs false",
			a:    datatype.NewBoolean(false),
			b:    datatype.NewBoolean(false),
			cmp:  sql.Equal,
		},
		{
			name: "true vs false",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewBoolean(false),
			cmp:  sql.Greater,
		},
		{
			name: "false vs true",
			a:    datatype.NewBoolean(false),
			b:    datatype.NewBoolean(true),
			cmp:  sql.Less,
		},
		{
			name: "true vs null",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "false vs null",
			a:    datatype.NewBoolean(false),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "boolean vs integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "boolean vs text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "boolean vs float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			cmp:  sql.Equal,
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cmp, err := test.a.Compare(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.cmp, cmp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.cmp, cmp)
			}
		})
	}
}

func TestBoolean_UnaryPlus(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.UnaryPlus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_UnaryMinus(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.UnaryMinus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Add(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Add(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Sub(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Sub(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Mul(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Mul(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Div(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Div(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Pow(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Pow(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Mod(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	v, err := b.Mod(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestBoolean_Equal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true == true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false == false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true == false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false == true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true == null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false == null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean == integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean == text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean == float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Equal(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_NotEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true != true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false != false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true != false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false != true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true != null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false != null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean != integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean != text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean != float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.NotEqual(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_GreaterThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true > true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false > false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true > false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false > true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true > null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false > null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean > integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean > text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean > float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.GreaterThan(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_LessThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true < true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false < false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true < false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false < true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true < null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false < null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean < integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean < text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean < float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.LessThan(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_GreaterOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true >= true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false >= false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true >= false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false >= true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true >= null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false >= null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean >= integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean >= text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean >= float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.GreaterOrEqual(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_LessOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true <= true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false <= false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true <= false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false <= true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true <= null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false <= null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean <= integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean <= text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean <= float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.LessOrEqual(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_And(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true AND true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false AND false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true AND false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "false AND true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true AND null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "false AND null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewBoolean(false),
		},
		{
			name: "boolean AND integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean AND text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean AND float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.And(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}

func TestBoolean_Or(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "true OR true",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false OR false",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "true OR false",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewBoolean(false),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false OR true",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewBoolean(true),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "true OR null",
			a:      datatype.NewBoolean(true),
			b:      datatype.NewNull(),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "false OR null",
			a:      datatype.NewBoolean(false),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "boolean OR integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean OR text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean OR float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Or(test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.result, value)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.result, value)
			}
		})
	}
}
