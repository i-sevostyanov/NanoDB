package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestFloat_Raw(t *testing.T) {
	t.Parallel()

	expected := float64(10)
	b := datatype.NewFloat(expected)

	switch value := b.Raw().(type) {
	case float64:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestFloat_Compare(t *testing.T) {
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
			a:    datatype.NewFloat(10),
			b:    datatype.NewFloat(10),
			cmp:  sql.Equal,
		},
		{
			name: "20 vs 10",
			a:    datatype.NewFloat(20),
			b:    datatype.NewFloat(10),
			cmp:  sql.Greater,
		},
		{
			name: "10 vs 20",
			a:    datatype.NewFloat(10),
			b:    datatype.NewFloat(20),
			cmp:  sql.Less,
		},
		{
			name: "10 vs null",
			a:    datatype.NewFloat(10),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "float vs integer",
			a:    datatype.NewFloat(10),
			b:    datatype.NewInteger(20),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "float vs string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "float vs boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_UnaryPlus(t *testing.T) {
	t.Parallel()

	f := datatype.NewFloat(10)
	value, err := f.UnaryPlus()

	require.NoError(t, err)
	assert.Equal(t, f, value)
}

func TestFloat_UnaryMinus(t *testing.T) {
	t.Parallel()

	expected := datatype.NewFloat(-10)
	f := datatype.NewFloat(10)
	value, err := f.UnaryMinus()

	require.NoError(t, err)
	assert.Equal(t, expected, value)
}

func TestFloat_Add(t *testing.T) {
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
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(20),
		},
		{
			name:   "10 + -10",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(0),
		},
		{
			name:   "10 + 0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(0),
			result: datatype.NewFloat(10),
		},
		{
			name:   "10.0 + 2",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(2),
			result: datatype.NewFloat(12.0),
		},
		{
			name:   "10 + null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float + string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float + boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 - 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(0),
		},
		{
			name:   "10.0 - -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(20),
		},
		{
			name:   "10.0 - 0.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(0),
			result: datatype.NewFloat(10),
		},
		{
			name:   "10.0 - 2",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(2),
			result: datatype.NewFloat(8.0),
		},
		{
			name:   "10.0 - null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float - string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float - boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Mul(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 * 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(100),
		},
		{
			name:   "10.0 * -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(-100),
		},
		{
			name:   "10.0 * 0.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(0),
			result: datatype.NewFloat(0),
		},
		{
			name:   "10.0 * 2",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(2),
			result: datatype.NewFloat(20.0),
		},
		{
			name:   "10.0 * null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float * string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float * boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Div(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 / 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(1),
		},
		{
			name:   "10.0 / -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(-1),
		},
		{
			name:   "10.0 / 2",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(2),
			result: datatype.NewFloat(5.0),
		},
		{
			name:   "10.0 / null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "10.0 / 0.0",
			a:    datatype.NewFloat(10),
			b:    datatype.NewFloat(0),
			err:  true,
		},
		{
			name: "float / string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float / boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Pow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 ^ 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(10000000000.0),
		},
		{
			name:   "10.0 ^ -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(0.0000000001),
		},
		{
			name:   "10.0 ^ 2",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(2),
			result: datatype.NewFloat(100),
		},
		{
			name:   "10.0 ^ null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name:   "10.0 ^ 0.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(0),
			result: datatype.NewFloat(1),
		},
		{
			name: "float ^ string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float ^ boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Mod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 % 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewFloat(0),
		},
		{
			name:   "10.0 % -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewFloat(0),
		},
		{
			name:   "10.0 % 3",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(3),
			result: datatype.NewFloat(1),
		},
		{
			name:   "10.0 % null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "10.0 % 0.0",
			a:    datatype.NewFloat(10),
			b:    datatype.NewFloat(0),
			err:  true,
		},
		{
			name: "float % string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float % boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_Equal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 == 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 == -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 == 10",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 == null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float == string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float == boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_NotEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 != 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 != -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 != 10",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 != null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float != string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float != boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_GreaterThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 > 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 > -10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(-10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 > 10",
			a:      datatype.NewFloat(10.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 > null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float > string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float > boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_LessThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 < 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "-10.0 < 10.0",
			a:      datatype.NewFloat(-10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "6.0 < 10",
			a:      datatype.NewFloat(6.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 < null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float < string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float < boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_GreaterOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 >= 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "-10.0 >= 10.0",
			a:      datatype.NewFloat(-10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "6.0 >= 10",
			a:      datatype.NewFloat(6.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "10.0 >= null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float >= string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float >= boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_LessOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "10.0 <= 10.0",
			a:      datatype.NewFloat(10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "-10.0 <= 10.0",
			a:      datatype.NewFloat(-10),
			b:      datatype.NewFloat(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "6.0 <= 10",
			a:      datatype.NewFloat(6.0),
			b:      datatype.NewInteger(10),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "10.0 <= null",
			a:      datatype.NewFloat(10),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "float <= string",
			a:    datatype.NewFloat(10),
			b:    datatype.NewString("xyz"),
			err:  true,
		},
		{
			name: "float <= boolean",
			a:    datatype.NewFloat(10),
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

func TestFloat_And(t *testing.T) {
	t.Parallel()

	a := datatype.NewFloat(10)
	b := datatype.NewFloat(20)
	v, err := a.And(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestFloat_Or(t *testing.T) {
	t.Parallel()

	a := datatype.NewFloat(10)
	b := datatype.NewFloat(20)
	v, err := a.Or(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}
