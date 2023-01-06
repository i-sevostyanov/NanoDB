package logical_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/logical"
)

func TestAnd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        sql.Value
		b        sql.Value
		expected sql.Value
		err      bool
	}{
		{
			name:     "true AND true",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "false AND false",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "true AND false",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "false AND true",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(false),
		},
		{
			name: "true AND null",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewNull(),
			err:  true,
		},
		{
			name: "false AND null",
			a:    datatype.NewBoolean(false),
			b:    datatype.NewNull(),
			err:  true,
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

			actual, err := logical.And(test.a, test.b)
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
