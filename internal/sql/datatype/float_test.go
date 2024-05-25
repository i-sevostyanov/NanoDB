package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestFloat_Raw(t *testing.T) {
	t.Parallel()

	expected := float64(10)
	b := datatype.NewFloat(expected)

	switch value := b.Raw().(type) {
	case float64:
		assert.InEpsilon(t, expected, value, 0)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestFloat_String(t *testing.T) {
	t.Parallel()

	f := datatype.NewFloat(10.0006)
	assert.Equal(t, "1.00006E+01", f.String())
}

func TestFloat_DataType(t *testing.T) {
	t.Parallel()

	f := datatype.NewFloat(10)
	assert.Equal(t, sql.Float, f.DataType())
}
