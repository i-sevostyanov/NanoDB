package plan_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestScan_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scheme := sql.Scheme{
		"id": sql.Column{
			Position:   0,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
		"name": sql.Column{
			Position:   1,
			Name:       "name",
			DataType:   sql.Text,
			PrimaryKey: false,
			Nullable:   false,
			Default:    nil,
		},
	}

	columns := []string{"id", "name"}

	table := sql.NewMockTable(ctrl)
	table.EXPECT().Scheme().Return(scheme)

	scan := plan.NewScan(table)
	assert.Equal(t, columns, scan.Columns())
}

func TestScan_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		table := sql.NewMockTable(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			table.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Close().Return(nil),
		)

		scan := plan.NewScan(table)
		iter, err := scan.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("something went wrong")

		table := sql.NewMockTable(ctrl)
		table.EXPECT().RowIter().Return(nil, expectedErr)

		scan := plan.NewScan(table)
		iter, err := scan.RowIter()
		require.Equal(t, expectedErr, err)
		require.Nil(t, iter)
	})
}
