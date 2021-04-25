package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestText_Raw(t *testing.T) {
	t.Parallel()

	expected := "xyz"
	s := datatype.NewText(expected)

	switch value := s.Raw().(type) {
	case string:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestText_DataType(t *testing.T) {
	t.Parallel()

	n := datatype.NewText("xyz")
	assert.Equal(t, sql.String, n.DataType())
}

func TestText_Compare(t *testing.T) {
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
			a:    datatype.NewText("xyz"),
			b:    datatype.NewText("xyz"),
			cmp:  sql.Equal,
		},
		{
			name: "qwerty vs xyz",
			a:    datatype.NewText("qwerty"),
			b:    datatype.NewText("xyz"),
			cmp:  sql.Less,
		},
		{
			name: "xyz vs qwerty",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewText("qwerty"),
			cmp:  sql.Greater,
		},
		{
			name: "xyz vs null",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewNull(),
			cmp:  sql.Greater,
		},
		{
			name: "text vs float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(20),
			cmp:  sql.Equal,
			err:  true,
		},
		{
			name: "text vs boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_UnaryPlus(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.UnaryPlus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_UnaryMinus(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.UnaryMinus()

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Add(t *testing.T) {
	t.Parallel()

	t.Run("text + text -> text", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewText("123")
		b := datatype.NewText("456")

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewText("123456"), value)
	})

	t.Run("text + null -> null", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewText("123")
		b := datatype.NewNull()

		value, err := a.Add(b)
		require.NoError(t, err)
		assert.Equal(t, datatype.NewNull(), value)
	})

	t.Run("text + integer -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewText("123")
		b := datatype.NewInteger(10)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("text + float -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewText("123")
		b := datatype.NewFloat(10)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("text + boolean -> error", func(t *testing.T) {
		t.Parallel()

		a := datatype.NewText("123")
		b := datatype.NewBoolean(true)

		value, err := a.Add(b)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}

func TestText_Sub(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.Sub(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Mul(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.Mul(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Div(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.Div(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Pow(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.Pow(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Mod(t *testing.T) {
	t.Parallel()

	s := datatype.NewText("xyz")
	v, err := s.Mod(nil)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Equal(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz == qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz == null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text == integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text == float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text == boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_NotEqual(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz != qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz != null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text != integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text != float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text != boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_GreaterThan(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz > qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz > null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text > integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text > float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text > boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_LessThan(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz < qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz < null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text < integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text < float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text < boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_GreaterOrEqual(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz >= qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "qwerty >= xyz",
			a:      datatype.NewText("qwerty"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "xyz >= null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text >= integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text >= float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text >= boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_LessOrEqual(t *testing.T) {
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
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz <= qwerty",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewText("qwerty"),
			result: datatype.NewBoolean(false),
		},
		{
			name:   "qwerty <= xyz",
			a:      datatype.NewText("qwerty"),
			b:      datatype.NewText("xyz"),
			result: datatype.NewBoolean(true),
		},
		{
			name:   "xyz <= null",
			a:      datatype.NewText("xyz"),
			b:      datatype.NewNull(),
			result: datatype.NewNull(),
		},
		{
			name: "text <= integer",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "text <= float",
			a:    datatype.NewText("xyz"),
			b:    datatype.NewFloat(10),
			err:  true,
		},
		{
			name: "text <= boolean",
			a:    datatype.NewText("xyz"),
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

func TestText_And(t *testing.T) {
	t.Parallel()

	a := datatype.NewText("xyz")
	b := datatype.NewText("xyz")
	v, err := a.And(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}

func TestText_Or(t *testing.T) {
	t.Parallel()

	a := datatype.NewText("xyz")
	b := datatype.NewText("xyz")
	v, err := a.Or(b)

	require.NotNil(t, err)
	require.Nil(t, v)
}
