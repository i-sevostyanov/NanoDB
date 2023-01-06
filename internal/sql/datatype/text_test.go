package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestText_Raw(t *testing.T) {
	t.Parallel()

	expected := "xyz"
	s := datatype.NewText(expected)

	switch value := s.Raw().(type) {
	case string:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestText_String(t *testing.T) {
	t.Parallel()

	value := "xyz"
	text := datatype.NewText(value)
	assert.Equal(t, value, text.String())
}

func TestText_DataType(t *testing.T) {
	t.Parallel()

	n := datatype.NewText("xyz")
	assert.Equal(t, sql.Text, n.DataType())
}
