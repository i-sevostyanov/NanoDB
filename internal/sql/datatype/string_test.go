package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestString_Raw(t *testing.T) {
	t.Parallel()

	expected := "xyz"
	s := datatype.NewString(expected)

	switch value := s.Raw().(type) {
	case string:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestString_Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    sql.Value
		b    sql.Value
		cmp  sql.CompareType
		err  bool
	}{
		{
			name: "xyz vs xyz",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewString("xyz"),
			cmp:  sql.Equal,
		},
		{
			name: "qwerty vs xyz",
			a:    datatype.NewString("qwerty"),
			b:    datatype.NewString("xyz"),
			cmp:  sql.Less,
		},
		{
			name: "xyz vs qwerty",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewString("qwerty"),
			cmp:  sql.Greater,
		},
		{
			name: "xyz vs null",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "string vs float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(20),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "string vs boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_UnaryPlus(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.UnaryPlus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_UnaryMinus(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.UnaryMinus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Add(t *testing.T) {
	t.Parallel()

	t.Run("string + string -> string", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewString("123")
		b := datatype.NewString("456")

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewString("123456"), value)
	})

	t.Run("string + null -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewString("123")
		b := datatype.NewNull()

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("string + integer -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewString("123")
		b := datatype.NewInteger(10)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("string + float -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewString("123")
		b := datatype.NewFloat(10)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("string + boolean -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewString("123")
		b := datatype.NewBoolean(true)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}

func TestString_Sub(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.Sub(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Mul(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.Mul(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Div(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.Div(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Pow(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.Pow(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Mod(t *testing.T) {
	t.Parallel()

	s := datatype.NewString("xyz")
	v, err := s.Mod(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Equal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz == xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz == qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz == null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string == integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string == float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string == boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_NotEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz != xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz != qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz != null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string != integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string != float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string != boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_GreaterThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz > xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz > qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz > null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string > integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string > float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string > boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_LessThan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz < xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz < qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz < null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string < integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string < float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string < boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_GreaterOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz >= xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz >= qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "qwerty >= xyz",
			a:      datatype.NewString("qwerty"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz >= null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string >= integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string >= float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string >= boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_LessOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      sql.Value
		b      sql.Value
		result sql.Value
		err    bool
	}{
		{
			name:   "xyz <= xyz",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz <= qwerty",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewString("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "qwerty <= xyz",
			a:      datatype.NewString("qwerty"),
			b:      datatype.NewString("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz <= null",
			a:      datatype.NewString("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "string <= integer",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "string <= float",
			a:    datatype.NewString("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "string <= boolean",
			a:    datatype.NewString("xyz"),
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

func TestString_And(t *testing.T) {
	t.Parallel()

	a := datatype.NewString("xyz")
	b := datatype.NewString("xyz")
	v, err := a.And(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestString_Or(t *testing.T) {
	t.Parallel()

	a := datatype.NewString("xyz")
	b := datatype.NewString("xyz")
	v, err := a.Or(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}
