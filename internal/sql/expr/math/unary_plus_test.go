package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/math"
)

func TestUnaryPlus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    sql.Value
		expected sql.Value
		err      bool
	}{
		{
			name:  "Boolean",
			value: datatype.NewBoolean(true),
			err:   true,
		},
		{
			name:     "Float",
			value:    datatype.NewFloat(-10.0),
			expected: datatype.NewFloat(-10.0),
		},
		{
			name:     "Integer",
			value:    datatype.NewInteger(10),
			expected: datatype.NewInteger(10),
		},
		{
			name:  "Null",
			value: datatype.NewNull(),
			err:   true,
		},
		{
			name:  "Text",
			value: datatype.NewText("xyz"),
			err:   true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := math.UnaryPlus(test.value)
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
