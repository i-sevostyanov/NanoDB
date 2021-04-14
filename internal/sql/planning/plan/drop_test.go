package plan_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestDropDatabase_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		dropper := NewMockDatabaseDropper(ctrl)
		dropper.EXPECT().DropDatabase(name).Return(nil)

		dropPlan := plan.NewDropDatabase(dropper, name)
		iter, err := dropPlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on drop database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		expectedErr := fmt.Errorf("something went wrong")

		dropper := NewMockDatabaseDropper(ctrl)
		dropper.EXPECT().DropDatabase(name).Return(expectedErr)

		dropPlan := plan.NewDropDatabase(dropper, name)
		iter, err := dropPlan.RowIter()

		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})
}

func TestDropTable_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "users"

		dropper := NewMockTableDropper(ctrl)
		dropper.EXPECT().DropTable(name).Return(nil)

		dropPlan := plan.NewDropTable(dropper, name)
		iter, err := dropPlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on drop table", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		expectedErr := fmt.Errorf("something went wrong")

		dropper := NewMockTableDropper(ctrl)
		dropper.EXPECT().DropTable(name).Return(expectedErr)

		createPlan := plan.NewDropTable(dropper, name)
		iter, err := createPlan.RowIter()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})
}
