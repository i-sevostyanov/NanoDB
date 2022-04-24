package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestInteger_Raw(t *testing.T) {
	t.Parallel()

	expected := int64(10)
	b := datatype.NewInteger(expected)

	switch value := b.Raw().(type) {
	case int64:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestInteger_String(t *testing.T) {
	t.Parallel()

	i := datatype.NewInteger(10)
	assert.Equal(t, "10", i.String())
}

func TestInteger_DataType(t *testing.T) {
	t.Parallel()

	i := datatype.NewInteger(10)
	assert.Equal(t, sql.Integer, i.DataType())
}

func TestInteger_Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    sql.Value
		b    sql.Value
		cmp  sql.CompareType
		err  bool
	}{
		{
			name: "10 vs 10",
			a:    datatype.NewInteger(10),
			b:    datatype.NewInteger(10),
			cmp:  sql.Equal,
		},
		{
			name: "20 vs 10",
			a:    datatype.NewInteger(20),
			b:    datatype.NewInteger(10),
			cmp:  sql.Greater,
		},
		{
			name: "10 vs 20",
			a:    datatype.NewInteger(10),
			b:    datatype.NewInteger(20),
			cmp:  sql.Less,
		},
		{
			name: "10 vs null",
			a:    datatype.NewInteger(10),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "integer vs float",
			a:    datatype.NewInteger(10),
			b:    datatype.NewFloat(20),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "integer vs text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "integer vs boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_UnaryPlus(t *testing.T) {
	t.Parallel()

	i := datatype.NewInteger(10)
	value, err := i.UnaryPlus()

	require.NoError(t, err)
	assert.Equal(t, i, value)
}

func TestInteger_UnaryMinus(t *testing.T) {
	t.Parallel()

	expected := datatype.NewInteger(-10)
	i := datatype.NewInteger(10)
	value, err := i.UnaryMinus()

	require.NoError(t, err)
	assert.Equal(t, expected, value)
}

func TestInteger_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 + 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewInteger(20),
		},
		{
			name:   "10 + -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewInteger(0),
		},
		{
			name:   "10 + 0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(0),
			result: datatype.NewInteger(10),
		},
		{
			name:   "10 + 2.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(2),
			result: datatype.NewFloat(12.0),
		},
		{
			name:   "10 + null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "integer + text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer + boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Add(test.b)
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

func TestInteger_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 - 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewInteger(0),
		},
		{
			name:   "10 - -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewInteger(20),
		},
		{
			name:   "10 - 0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(0),
			result: datatype.NewInteger(10),
		},
		{
			name:   "10 - 2.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(2),
			result: datatype.NewFloat(8.0),
		},
		{
			name:   "10 - null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "integer - text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer - boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Sub(test.b)
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

func TestInteger_Mul(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 * 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewInteger(100),
		},
		{
			name:   "10 * -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewInteger(-100),
		},
		{
			name:   "10 * 0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(0),
			result: datatype.NewInteger(0),
		},
		{
			name:   "10 * 2.0",
			a:      datatype.NewInteger(10.0),
			b:      datatype.NewFloat(2),
			result: datatype.NewFloat(20.0),
		},
		{
			name:   "10 * null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "integer * text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer * boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Mul(test.b)
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

func TestInteger_Div(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 / 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewInteger(1),
		},
		{
			name:   "10 / -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewInteger(-1),
		},
		{
			name:   "10 / 2.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(2.0),
			result: datatype.NewFloat(5.0),
		},
		{
			name:   "10 / null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "10 / 0",
			a:    datatype.NewInteger(10),
			b:    datatype.NewInteger(0),
			err:  true,
		},
		{
			name: "10 / 0.0",
			a:    datatype.NewInteger(10),
			b:    datatype.NewFloat(0),
			err:  true,
		},
		{
			name: "integer / text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer / boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Div(test.b)
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

func TestInteger_Pow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 ^ 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewFloat(10000000000),
		},
		{
			name:   "10 ^ -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewFloat(0.0000000001),
		},
		{
			name:   "10 ^ 2.0",
			a:      datatype.NewInteger(10.0),
			b:      datatype.NewFloat(2),
			result: datatype.NewFloat(100),
		},
		{
			name:   "10.0 ^ null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "10 ^ 0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(0),
			result: datatype.NewFloat(1),
		},
		{
			name: "integer ^ text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer ^ boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Pow(test.b)
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

func TestInteger_Mod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 % 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewInteger(0),
		},
		{
			name:   "10 % -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewInteger(0),
		},
		{
			name:   "10 % 3.0",
			a:      datatype.NewInteger(10.0),
			b:      datatype.NewFloat(3),
			result: datatype.NewFloat(1),
		},
		{
			name:   "10 % null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "10 % 0",
			a:    datatype.NewInteger(10),
			b:    datatype.NewInteger(0),
			err:  true,
		},
		{
			name: "10 % 0.0",
			a:    datatype.NewInteger(10),
			b:    datatype.NewFloat(0),
			err:  true,
		},
		{
			name: "integer % text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "integer % boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			value, err := test.a.Mod(test.b)
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

func TestInteger_Equal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 == 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 == -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 == 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 == null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float == text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float == boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_NotEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 != 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 != -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 != 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 != null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float != text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float != boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_GreaterThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 > 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 > -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 > 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 > null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float > text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float > boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_LessThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 < 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 < -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 < 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 < null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float < text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float < boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_GreaterOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 >= 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 >= -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 >= 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 >= 100",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(100),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 >= null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float >= text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float >= boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_LessOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10 <= 10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 <= -10",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(-10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10 <= 10.0",
			a:      datatype.NewInteger(10),
			b:      datatype.NewFloat(10.0),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 <= 100",
			a:      datatype.NewInteger(10),
			b:      datatype.NewInteger(100),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10 <= null",
			a:      datatype.NewInteger(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float <= text",
			a:    datatype.NewInteger(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float <= boolean",
			a:    datatype.NewInteger(10),
			b:    datatype.NewBoolean(true),
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

func TestInteger_And(t *testing.T) {
	t.Parallel()

	a := datatype.NewInteger(10)
	b := datatype.NewInteger(20)
	v, err := a.And(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestInteger_Or(t *testing.T) {
	t.Parallel()

	a := datatype.NewInteger(10)
	b := datatype.NewInteger(20)
	v, err := a.Or(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}
