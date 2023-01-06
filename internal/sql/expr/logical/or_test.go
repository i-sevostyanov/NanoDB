package logical_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/logical"
)

func TestOr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        sql.Value
		b        sql.Value
		expected sql.Value
		err      bool
	}{
		{
			name:     "true OR true",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "false OR false",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(false),
		},
		{
			name:     "true OR false",
			a:        datatype.NewBoolean(true),
			b:        datatype.NewBoolean(false),
			expected: datatype.NewBoolean(true),
		},
		{
			name:     "false OR true",
			a:        datatype.NewBoolean(false),
			b:        datatype.NewBoolean(true),
			expected: datatype.NewBoolean(true),
		},
		{
			name: "true OR null",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewNull(),
			err:  true,
		},
		{
			name: "false OR null",
			a:    datatype.NewBoolean(false),
			b:    datatype.NewNull(),
			err:  true,
		},
		{
			name: "boolean OR integer",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewInteger(10),
			err:  true,
		},
		{
			name: "boolean OR text",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewText("xyz"),
			err:  true,
		},
		{
			name: "boolean OR float",
			a:    datatype.NewBoolean(true),
			b:    datatype.NewFloat(10.2),
			err:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := logical.Or(test.a, test.b)
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
