package comparison_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/comparison"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        sql.Value
		b        sql.Value
		expected sql.CompareType
		err      bool
	}{
		// Boolean
		{
			name:     "true vs true",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(true),
			expected: sql.Equal,
		},
		{
			name:     "false vs false",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(false),
			expected: sql.Equal,
		},
		{
			name:     "true vs false",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(false),
			expected: sql.Greater,
		},
		{
			name:     "false vs true",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(true),
			expected: sql.Less,
		},
		{
			name:     "true vs null",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewNull(),
			expected: sql.Greater,
		},
		{
			name:     "false vs null",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewNull(),
			expected: sql.Greater,
		},
		{
			name:     "boolean vs integer",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewInteger(10),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "boolean vs text",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewText("xyz"),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "boolean vs float",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewFloat(10.2),
			expected: sql.Equal,
			err:      true,
		},
		// Float
		{
			name:     "10 vs 10",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(10),
			expected: sql.Equal,
		},
		{
			name:     "20 vs 10",
			a:        datatype.NewFloat(20),
			b:        datatype.NewFloat(10),
			expected: sql.Greater,
		},
		{
			name:     "10 vs 20",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(20),
			expected: sql.Less,
		},
		{
			name:     "10 vs null",
			a:        datatype.NewFloat(10),
			b:        datatype.NewNull(),
			expected: sql.Greater,
		},
		{
			name:     "float vs integer",
			a:        datatype.NewFloat(10),
			b:        datatype.NewInteger(20),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "float vs text",
			a:        datatype.NewFloat(10),
			b:        datatype.NewText("xyz"),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "float vs boolean",
			a:        datatype.NewFloat(10),
			b:        datatype.NewBoolean(true),
			expected: sql.Equal,
			err:      true,
		},
		// Integer
		{
			name:     "10 vs 10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(10),
			expected: sql.Equal,
		},
		{
			name:     "20 vs 10",
			a:        datatype.NewInteger(20),
			b:        datatype.NewInteger(10),
			expected: sql.Greater,
		},
		{
			name:     "10 vs 20",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(20),
			expected: sql.Less,
		},
		{
			name:     "10 vs null",
			a:        datatype.NewInteger(10),
			b:        datatype.NewNull(),
			expected: sql.Greater,
		},
		{
			name:     "integer vs float",
			a:        datatype.NewInteger(10),
			b:        datatype.NewFloat(20),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "integer vs text",
			a:        datatype.NewInteger(10),
			b:        datatype.NewText("xyz"),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "integer vs boolean",
			a:        datatype.NewInteger(10),
			b:        datatype.NewBoolean(true),
			expected: sql.Equal,
			err:      true,
		},
		// Null
		{
			name:     "null vs null",
			a:        datatype.NewNull(),
			b:        datatype.NewNull(),
			expected: sql.Equal,
		},
		{
			name:     "null vs integer",
			a:        datatype.NewNull(),
			b:        datatype.NewInteger(10),
			expected: sql.Less,
		},
		{
			name:     "null vs float",
			a:        datatype.NewNull(),
			b:        datatype.NewFloat(10),
			expected: sql.Less,
		},
		{
			name:     "null vs text",
			a:        datatype.NewNull(),
			b:        datatype.NewText("xyz"),
			expected: sql.Less,
		},
		{
			name:     "null vs boolean",
			a:        datatype.NewNull(),
			b:        datatype.NewBoolean(true),
			expected: sql.Less,
		},
		// Text
		{
			name:     "xyz vs xyz",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewText("xyz"),
			expected: sql.Equal,
		},
		{
			name:     "qwerty vs xyz",
			a:        datatype.NewText("qwerty"),
			b:        datatype.NewText("xyz"),
			expected: sql.Less,
		},
		{
			name:     "xyz vs qwerty",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewText("qwerty"),
			expected: sql.Greater,
		},
		{
			name:     "xyz vs null",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewNull(),
			expected: sql.Greater,
		},
		{
			name:     "text vs float",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewFloat(20),
			expected: sql.Equal,
			err:      true,
		},
		{
			name:     "text vs boolean",
			a:        datatype.NewText("xyz"),
			b:        datatype.NewBoolean(true),
			expected: sql.Equal,
			err:      true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := comparison.Compare(test.a, test.b)
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
