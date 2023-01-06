package sql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

func TestDataType_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		datatype sql.DataType
		expected string
	}{
		{
			datatype: sql.Null,
			expected: "null",
		},
		{
			datatype: sql.Integer,
			expected: "integer",
		},
		{
			datatype: sql.Float,
			expected: "float",
		},
		{
			datatype: sql.Text,
			expected: "text",
		},
		{
			datatype: sql.Boolean,
			expected: "boolean",
		},
		{
			datatype: sql.DataType(999),
			expected: "DataType<999>",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.datatype.String())
	}
}
