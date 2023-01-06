package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestInteger_Raw(t *testing.T) {
	t.Parallel()

	expected := int64(10)
	b := datatype.NewInteger(expected)

	switch value := b.Raw().(type) {
	case int64:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestInteger_String(t *testing.T) {
	t.Parallel()

	i := datatype.NewInteger(10)
	assert.Equal(t, "10", i.String())
}

func TestInteger_DataType(t *testing.T) {
	t.Parallel()

	i := datatype.NewInteger(10)
	assert.Equal(t, sql.Integer, i.DataType())
}
