package datatype_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

func TestNull_Raw(t *testing.T) {
	t.Parallel()

	null := datatype.NewNull()

	switch value := null.Raw().(type) {
	case nil:
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestNull_String(t *testing.T) {
	t.Parallel()

	n := datatype.NewNull()
	assert.Equal(t, "null", n.String())
}

func TestNull_DataType(t *testing.T) {
	t.Parallel()

	n := datatype.NewNull()
	assert.Equal(t, sql.Null, n.DataType())
}
