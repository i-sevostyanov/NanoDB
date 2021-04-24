package planner_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/planner"
)

func TestPlanner_CreateDatabase(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		database := "users"
		errDBNotExist := fmt.Errorf("databasse not exist")

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(database).Return(nil, errDBNotExist)

		expected := plan.NewCreateDatabase(catalog, database)
		stmt := &ast.CreateDatabaseStatement{
			Database: database,
		}

		planNode, err := planner.New(catalog).Plan("", stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("returns error if database already exist", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "users"

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)

		stmt := &ast.CreateDatabaseStatement{
			Database: databaseName,
		}

		planNode, err := planner.New(catalog).Plan("", stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})
}

func TestPlanner_DropDatabase(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		database := "users"

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(database).Return(nil, nil)

		expected := plan.NewDropDatabase(catalog, database)
		stmt := &ast.DropDatabaseStatement{
			Database: database,
		}

		planNode, err := planner.New(catalog).Plan("", stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("return error if database not exist", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		database := "users"
		expectedErr := fmt.Errorf("something went wrong")

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(database).Return(nil, expectedErr)

		stmt := &ast.DropDatabaseStatement{
			Database: database,
		}

		planNode, err := planner.New(catalog).Plan("", stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})
}

func TestPlanner_CreateTable(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.String,
						Literal: "<unknown>",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
				{
					Name: "salary",
					Type: token.Float,
					Default: &ast.ScalarExpr{
						Type:    token.Float,
						Literal: "0",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
				{
					Name:       "is_active",
					Type:       token.Boolean,
					Default:    nil,
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		scheme := sql.Scheme{
			"id": {
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
			"name": {
				Position:   1,
				Name:       "name",
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   true,
				Default:    datatype.NewString("<unknown>"),
			},
			"salary": {
				Position:   2,
				Name:       "salary",
				DataType:   sql.Float,
				PrimaryKey: false,
				Nullable:   true,
				Default:    datatype.NewFloat(0),
			},
			"is_active": {
				Position:   3,
				Name:       "is_active",
				DataType:   sql.Boolean,
				PrimaryKey: false,
				Nullable:   true,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		expected := plan.NewCreateTable(database, tableName, scheme)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("nullable column pass null as default value", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		scheme := sql.Scheme{
			"id": {
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
			"name": {
				Position:   1,
				Name:       "name",
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   true,
				Default:    datatype.NewNull(),
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		expected := plan.NewCreateTable(database, tableName, scheme)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("returns error if can't get database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		expectedErr := fmt.Errorf("something went wrong")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(databaseName).Return(nil, expectedErr)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if table already exist", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, nil)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})

	t.Run("return error if no primary key specified", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: false,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})

	t.Run("return error if more than one primary key specified", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   true,
					PrimaryKey: true,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if column default value have different type", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Integer,
						Literal: "10",
					},
					Nullable:   true,
					PrimaryKey: false,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if column default value is null but column not nullable", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		errTableNotExist := fmt.Errorf("table not exist")

		stmt := &ast.CreateTableStatement{
			Table: tableName,
			Columns: []ast.Column{
				{
					Name:       "id",
					Type:       token.Integer,
					Default:    nil,
					Nullable:   false,
					PrimaryKey: true,
				},
				{
					Name: "name",
					Type: token.String,
					Default: &ast.ScalarExpr{
						Type:    token.Null,
						Literal: "null",
					},
					Nullable:   false,
					PrimaryKey: false,
				},
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, errTableNotExist)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})
}

func TestPlanner_DropTable(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, nil)

		expected := plan.NewDropTable(database, tableName)
		stmt := &ast.DropTableStatement{
			Table: tableName,
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("returns error if can't get database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		expectedErr := fmt.Errorf("something went wrong")

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(databaseName).Return(nil, expectedErr)

		stmt := &ast.DropTableStatement{
			Table: tableName,
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if table not exist", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		expectedErr := fmt.Errorf("something went wrong")

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, expectedErr)

		stmt := &ast.DropTableStatement{
			Table: tableName,
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})
}

func TestPlanner_Select(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"salary": sql.Column{
				Position:   2,
				Name:       "salary",
				DataType:   sql.Float,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme).Times(3)

		stmt := &ast.SelectStatement{
			Result: []ast.ResultStatement{
				{
					Expr: &ast.IdentExpr{
						Name: "id",
					},
				},
				{
					Expr: &ast.IdentExpr{
						Name: "name",
					},
				},
				{
					Expr: &ast.IdentExpr{
						Name: "salary",
					},
				},
			},
			From: &ast.FromStatement{
				Table: tableName,
			},
			Where: &ast.WhereStatement{
				Expr: &ast.BinaryExpr{
					Left: &ast.IdentExpr{
						Name: "id",
					},
					Operator: token.GreaterThan,
					Right: &ast.ScalarExpr{
						Type:    token.Integer,
						Literal: "10",
					},
				},
			},
			OrderBy: &ast.OrderByStatement{
				Column:    "salary",
				Direction: token.Desc,
			},
			Limit: &ast.LimitStatement{
				Value: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
			},
			Offset: &ast.OffsetStatement{
				Value: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "2",
				},
			},
		}

		projections := []plan.Projection{
			{
				Alias: "",
				Expr:  expr.Column{Name: "id", Position: 0},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "name", Position: 1},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "salary", Position: 2},
			},
		}

		cond, err := expr.New(stmt.Where.Expr, scheme)
		require.NoError(t, err)

		expected := plan.NewLimit(
			10,
			plan.NewOffset(
				2,
				plan.NewProject(
					projections,
					plan.NewSort(
						2,
						plan.Descending,
						plan.NewFilter(
							cond,
							plan.NewScan(
								table,
							),
						),
					),
				),
			),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("select *", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"salary": sql.Column{
				Position:   2,
				Name:       "salary",
				DataType:   sql.Float,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme)

		stmt := &ast.SelectStatement{
			Result: []ast.ResultStatement{
				{
					Expr: &ast.AsteriskExpr{},
				},
			},
			From: &ast.FromStatement{
				Table: tableName,
			},
		}

		projections := []plan.Projection{
			{
				Alias: "",
				Expr:  expr.Column{Name: "id", Position: 0},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "name", Position: 1},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "salary", Position: 2},
			},
		}

		expected := plan.NewProject(
			projections,
			plan.NewScan(
				table,
			),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("select *, id", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"salary": sql.Column{
				Position:   2,
				Name:       "salary",
				DataType:   sql.Float,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme)

		stmt := &ast.SelectStatement{
			Result: []ast.ResultStatement{
				{
					Expr: &ast.AsteriskExpr{},
				},
				{
					Expr: &ast.IdentExpr{
						Name: "id",
					},
				},
			},
			From: &ast.FromStatement{
				Table: tableName,
			},
		}

		projections := []plan.Projection{
			{
				Alias: "",
				Expr:  expr.Column{Name: "id", Position: 0},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "name", Position: 1},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "salary", Position: 2},
			},
			{
				Alias: "",
				Expr:  expr.Column{Name: "id", Position: 0},
			},
		}

		expected := plan.NewProject(
			projections,
			plan.NewScan(
				table,
			),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})
}

func TestPlanner_Insert(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"salary": sql.Column{
				Position:   2,
				Name:       "salary",
				DataType:   sql.Float,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)
		seq := sql.NewMockSequence(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme)
		table.EXPECT().Sequence().Return(seq)
		seq.EXPECT().Next().Return(int64(1))

		stmt := &ast.InsertStatement{
			Table: tableName,
			Columns: []string{
				"name",
				"salary",
			},
			Values: []ast.Expression{
				&ast.ScalarExpr{
					Type:    token.String,
					Literal: "Mad Max",
				},
				&ast.ScalarExpr{
					Type:    token.Float,
					Literal: "200.2",
				},
			},
		}

		key := int64(1)
		row := sql.Row{
			datatype.NewInteger(key),
			datatype.NewString("Mad Max"),
			datatype.NewFloat(200.2),
		}

		expected := plan.NewInsert(table, key, row)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})
}

func TestPlanner_Update(t *testing.T) {
	t.Parallel()

	t.Run("update with filter", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		pkIndex := uint8(0)

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		stmt := &ast.UpdateStatement{
			Table: tableName,
			Set: []ast.SetStatement{
				{
					Column: "name",
					Value: &ast.ScalarExpr{
						Type:    token.String,
						Literal: "xyz",
					},
				},
			},
			Where: &ast.WhereStatement{
				Expr: &ast.BinaryExpr{
					Left:     &ast.IdentExpr{Name: "id"},
					Operator: token.Equal,
					Right: &ast.ScalarExpr{
						Type:    token.Integer,
						Literal: "10",
					},
				},
			},
		}

		astExpr := &ast.ScalarExpr{
			Type:    token.String,
			Literal: "xyz",
		}

		nameExpr, err := expr.New(astExpr, scheme)
		require.NoError(t, err)

		columns := map[uint8]expr.Node{
			1: nameExpr,
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme).Times(2)
		table.EXPECT().PrimaryKey().Return(scheme["id"])

		cond, err := expr.New(stmt.Where.Expr, scheme)
		require.NoError(t, err)

		expected := plan.NewUpdate(
			table,
			pkIndex,
			columns,
			plan.NewFilter(
				cond,
				plan.NewScan(table),
			),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("update without filter", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tableName := "users"
		databaseName := "playground"
		pkIndex := uint8(0)

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		stmt := &ast.UpdateStatement{
			Table: tableName,
			Set: []ast.SetStatement{
				{
					Column: "name",
					Value: &ast.ScalarExpr{
						Type:    token.String,
						Literal: "xyz",
					},
				},
			},
		}

		astExpr := &ast.ScalarExpr{
			Type:    token.String,
			Literal: "xyz",
		}

		nameExpr, err := expr.New(astExpr, scheme)
		require.NoError(t, err)

		columns := map[uint8]expr.Node{
			1: nameExpr,
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().Scheme().Return(scheme)
		table.EXPECT().PrimaryKey().Return(scheme["id"])

		expected := plan.NewUpdate(
			table,
			pkIndex,
			columns,
			plan.NewScan(table),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})
}

func TestPlanner_Delete(t *testing.T) {
	t.Parallel()

	t.Run("delete with filter", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "playground"
		tableName := "users"
		pkIndex := uint8(0)

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().PrimaryKey().Return(scheme["id"])
		table.EXPECT().Scheme().Return(scheme)

		stmt := &ast.DeleteStatement{
			Table: tableName,
			Where: &ast.WhereStatement{
				Expr: &ast.BinaryExpr{
					Left:     &ast.IdentExpr{Name: "id"},
					Operator: token.Equal,
					Right:    &ast.ScalarExpr{Type: token.Integer, Literal: "10"},
				},
			},
		}

		cond, err := expr.New(stmt.Where.Expr, scheme)
		require.NoError(t, err)

		expected := plan.NewDelete(
			table,
			pkIndex,
			plan.NewFilter(
				cond,
				plan.NewScan(table),
			),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("delete without filter", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "playground"
		tableName := "users"
		pkIndex := uint8(0)

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
				DataType:   sql.String,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)
		table := sql.NewMockTable(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(table, nil)
		table.EXPECT().PrimaryKey().Return(scheme["id"])

		stmt := &ast.DeleteStatement{
			Table: tableName,
		}

		expected := plan.NewDelete(
			table,
			pkIndex,
			plan.NewScan(table),
		)

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NoError(t, err)
		assert.Equal(t, expected, planNode)
	})

	t.Run("returns error if can't get a database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "playground"
		tableName := "users"
		expectedErr := fmt.Errorf("something went wrong")

		catalog := sql.NewMockCatalog(ctrl)
		catalog.EXPECT().GetDatabase(databaseName).Return(nil, expectedErr)

		stmt := &ast.DeleteStatement{
			Table: tableName,
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if table not exist", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "playground"
		tableName := "users"
		expectedErr := fmt.Errorf("something went wrong")

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, expectedErr)

		stmt := &ast.DeleteStatement{
			Table: tableName,
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, planNode)
	})

	t.Run("returns error if table not specified", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		databaseName := "playground"
		tableName := "users"

		catalog := sql.NewMockCatalog(ctrl)
		database := sql.NewMockDatabase(ctrl)

		catalog.EXPECT().GetDatabase(databaseName).Return(database, nil)
		database.EXPECT().GetTable(tableName).Return(nil, nil)

		stmt := &ast.DeleteStatement{
			Table: tableName,
			Where: &ast.WhereStatement{Expr: nil},
		}

		planNode, err := planner.New(catalog).Plan(databaseName, stmt)
		require.NotNil(t, err)
		assert.Nil(t, planNode)
	})
}

func TestPlanner_Empty(t *testing.T) {
	t.Parallel()

	databaseName := "playground"
	planNode, err := planner.New(nil).Plan(databaseName, nil)
	require.NoError(t, err)
	assert.Equal(t, plan.NewRows(), planNode)
}
