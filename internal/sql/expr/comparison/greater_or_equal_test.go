package comparison_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/comparison"
)

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        sql.Value
		b        sql.Value
		expected sql.Value
		err      bool
	}{
		// Boolean
		{
			name:     "true >= true",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "false >= false",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "true >= false",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "false >= true",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "true >= null",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name:     "false >= null",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
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
		// Float
		{
			name:     "10.0 >= 10.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(10),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "-10.0 >= 10.0",
			a:        datatype.NewFloat(-10),
			b:        datatype.NewFloat(10),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "6.0 >= 10",
			a:        datatype.NewFloat(6.0),
			b:        datatype.NewInteger(10),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "10.0 >= null",
			a:        datatype.NewFloat(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name: "float >= text",
			a:    datatype.NewFloat(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float >= boolean",
			a:    datatype.NewFloat(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
		// Integer
		{
			name:     "10 >= 10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(10),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "10 >= -10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(-10),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "10 >= 10.0",
			a:        datatype.NewInteger(10),
			b:        datatype.NewFloat(10.0),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "10 >= 100",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(100),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "10 >= null",
			a:        datatype.NewInteger(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
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
		// Null
		{
			name:     "null > null",
			a:        datatype.NewNull(),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name:     "null > boolean",
			a:        datatype.NewNull(),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewNull(),
		},
		{
			name:     "null > float",
			a:        datatype.NewNull(),
			b:        datatype.NewFloat(10.0),
			expected: datatype.NewNull(),
		},
		{
			name:     "null > integer",
			a:        datatype.NewNull(),
			b:        datatype.NewInteger(10),
			expected: datatype.NewNull(),
		},
		// Text
		{
			name:     "xyz >= xyz",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewText("xyz"),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "xyz >= qwerty",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewText("qwerty"),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "qwerty >= xyz",
			a:        datatype.NewText("qwerty"),
			b:        datatype.NewText("xyz"),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "xyz >= null",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
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

			actual, err := comparison.GreaterOrEqual(test.a, test.b)
			if test.err {
				require.NotNil(t, err)
				assert.Equal(t, test.expected, actual)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}
