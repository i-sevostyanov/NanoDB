package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestBoolean_Raw(t *testing.T) {
	t.Parallel()

	fn := func(expected bool) {
		b := datatype.NewBoolean(expected)

		switch value := b.Raw().(type) {
		case bool:
			assert.Equal(t, expected, value)
		default:
			assert.Failf(t, "fail", "unexpected type %T", value)
		}
	}

	fn(true)
	fn(false)
}

func TestBoolean_String(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	assert.Equal(t, "true", b.String())

	b = datatype.NewBoolean(false)
	assert.Equal(t, "false", b.String())
}

func TestBoolean_DataType(t *testing.T) {
	t.Parallel()

	b := datatype.NewBoolean(true)
	assert.Equal(t, sql.Boolean, b.DataType())
}
