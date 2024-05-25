package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/math"
)

func TestPow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        sql.Value
		b        sql.Value
		expected sql.Value
		err      bool
	}{
		// Float
		{
			name:     "10.0 ^ 10.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(10),
			expected: datatype.NewFloat(10000000000.0),
		},
		{
			name:     "10.0 ^ -10.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(-10),
			expected: datatype.NewFloat(0.0000000001),
		},
		{
			name:     "10.0 ^ 2",
			a:        datatype.NewFloat(10.0),
			b:        datatype.NewInteger(2),
			expected: datatype.NewFloat(100),
		},
		{
			name:     "10.0 ^ null",
			a:        datatype.NewFloat(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name:     "10.0 ^ 0.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(0),
			expected: datatype.NewFloat(1),
		},
		{
			name: "float ^ text",
			a:    datatype.NewFloat(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float ^ boolean",
			a:    datatype.NewFloat(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
		// Integer
		{
			name:     "10 ^ 10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(10),
			expected: datatype.NewFloat(10000000000),
		},
		{
			name:     "10 ^ -10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(-10),
			expected: datatype.NewFloat(0.0000000001),
		},
		{
			name:     "10 ^ 2.0",
			a:        datatype.NewInteger(10.0),
			b:        datatype.NewFloat(2),
			expected: datatype.NewFloat(100),
		},
		{
			name:     "10.0 ^ null",
			a:        datatype.NewInteger(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name:     "10 ^ 0",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(0),
			expected: datatype.NewFloat(1),
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := math.Pow(test.a, test.b)
			if test.err {
				require.Error(t, err)
				assert.Equal(t, test.expected, actual)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}
