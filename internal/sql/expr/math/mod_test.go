package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/math"
)

func TestMod(t *testing.T) {
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
			name:     "10.0 % 10.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(10),
			expected: datatype.NewFloat(0),
		},
		{
			name:     "10.0 % -10.0",
			a:        datatype.NewFloat(10),
			b:        datatype.NewFloat(-10),
			expected: datatype.NewFloat(0),
		},
		{
			name:     "10.0 % 3",
			a:        datatype.NewFloat(10.0),
			b:        datatype.NewInteger(3),
			expected: datatype.NewFloat(1),
		},
		{
			name:     "10.0 % null",
			a:        datatype.NewFloat(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
		},
		{
			name: "10.0 % 0.0",
			a:    datatype.NewFloat(10),
			b:    datatype.NewFloat(0),
			err:  true,
		},
		{
			name: "10.0 % 0",
			a:    datatype.NewFloat(10),
			b:    datatype.NewInteger(0),
			err:  true,
		},
		{
			name: "float % text",
			a:    datatype.NewFloat(10),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "float % boolean",
			a:    datatype.NewFloat(10),
			b:    datatype.NewBoolean(true),
			err:  true,
		},
		// Integer
		{
			name:     "10 % 10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(10),
			expected: datatype.NewInteger(0),
		},
		{
			name:     "10 % -10",
			a:        datatype.NewInteger(10),
			b:        datatype.NewInteger(-10),
			expected: datatype.NewInteger(0),
		},
		{
			name:     "10 % 3.0",
			a:        datatype.NewInteger(10.0),
			b:        datatype.NewFloat(3),
			expected: datatype.NewFloat(1),
		},
		{
			name:     "10 % null",
			a:        datatype.NewInteger(10),
			b:        datatype.NewNull(),
			expected: datatype.NewNull(),
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

			actual, err := math.Mod(test.a, test.b)
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
