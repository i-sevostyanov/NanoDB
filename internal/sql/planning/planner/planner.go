package planner

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

type Planner struct {
	catalog sql.Catalog
}

func New(catalog sql.Catalog) *Planner {
	return &Planner{
		catalog: catalog,
	}
}

func (p *Planner) Plan(database string, node ast.Node) (plan.Node, error) {
	switch stmt := node.(type) {
	// DDL
	case *ast.CreateDatabaseStatement:
		return p.planCreateDatabase(stmt)
	case *ast.DropDatabaseStatement:
		return p.planDropDatabase(stmt)
	case *ast.CreateTableStatement:
		return p.planCreateTable(database, stmt)
	case *ast.DropTableStatement:
		return p.planDropTable(database, stmt)
	// DML
	case *ast.SelectStatement:
		return p.planSelect(database, stmt)
	case *ast.InsertStatement:
		return p.planInsert(database, stmt)
	case *ast.UpdateStatement:
		return p.planUpdate(database, stmt)
	case *ast.DeleteStatement:
		return p.planDelete(database, stmt)
	case nil:
		return plan.NewRows(), nil
	default:
		return nil, fmt.Errorf("unexpected statement %T", stmt)
	}
}

func (p *Planner) planSelect(database string, stmt *ast.SelectStatement) (plan.Node, error) {
	var (
		table sql.Table
		node  plan.Node
		err   error
	)

	if table, node, err = p.planScan(database, stmt.From); err != nil {
		return nil, fmt.Errorf("failed to plan scan: %w", err)
	}

	if node, err = p.planFilter(table, stmt.Where, node); err != nil {
		return nil, fmt.Errorf("failed to plan filter: %w", err)
	}

	if node, err = p.planSort(table, stmt.OrderBy, node); err != nil {
		return nil, fmt.Errorf("failed to plan sort: %w", err)
	}

	if node, err = p.planProject(table, stmt.Result, node); err != nil {
		return nil, fmt.Errorf("failed to plan project: %w", err)
	}

	if node, err = p.planOffset(stmt.Offset, node); err != nil {
		return nil, fmt.Errorf("failed to plan offset: %w", err)
	}

	if node, err = p.planLimit(stmt.Limit, node); err != nil {
		return nil, fmt.Errorf("failed to plan limit: %w", err)
	}

	return node, nil
}

func (p *Planner) planScan(database string, stmt *ast.FromStatement) (sql.Table, plan.Node, error) {
	if stmt == nil {
		return nil, plan.NewRows(sql.Row{}), nil
	}

	table, err := p.getTable(database, stmt.Table)
	if err != nil {
		return nil, nil, err
	}

	return table, plan.NewScan(table), nil
}

func (p *Planner) planInsert(database string, stmt *ast.InsertStatement) (plan.Node, error) {
	table, err := p.getTable(database, stmt.Table)
	if err != nil {
		return nil, err
	}

	if len(stmt.Columns) != len(stmt.Values) {
		return nil, fmt.Errorf("number of expressions should be equal to the number of columns")
	}

	var key int64

	scheme := table.Scheme()
	row := make(sql.Row, len(scheme))

	for i := range scheme {
		value := scheme[i].Default

		if scheme[i].PrimaryKey {
			key = table.Sequence().Next()
			value = datatype.NewInteger(key)
		}

		row[scheme[i].Position] = value
	}

	for idx, columnName := range stmt.Columns {
		column, ok := scheme[columnName]
		if !ok {
			return nil, fmt.Errorf("column %q not found", columnName)
		}

		valueExpr, err := expr.New(stmt.Values[idx], scheme)
		if err != nil {
			return nil, err
		}

		value, err := valueExpr.Eval(nil)
		if err != nil {
			return nil, err
		}

		if column.DataType != value.DataType() {
			switch {
			case value.DataType() != sql.Null:
				return nil, fmt.Errorf("invalid value for column %q", column.Name)
			case value.DataType() == sql.Null && !column.Nullable:
				return nil, fmt.Errorf("null value in column %q violates not-null constraint", column.Name)
			}
		}

		if column.PrimaryKey {
			if key, ok = value.Raw().(int64); !ok {
				return nil, fmt.Errorf("unsupported primary key type %T", value.Raw())
			}
		}

		row[column.Position] = value
	}

	return plan.NewInsert(table, key, row), nil
}

func (p *Planner) planUpdate(database string, stmt *ast.UpdateStatement) (plan.Node, error) {
	var (
		table   sql.Table
		node    plan.Node
		columns map[uint8]expr.Node
		err     error
	)

	if table, err = p.getTable(database, stmt.Table); err != nil {
		return nil, err
	}

	if node, err = p.planFilter(table, stmt.Where, plan.NewScan(table)); err != nil {
		return nil, fmt.Errorf("failed to plan filter: %w", err)
	}

	if columns, err = p.planUpdateColumns(table.Scheme(), stmt.Set); err != nil {
		return nil, fmt.Errorf("failed to plan columns for update: %w", err)
	}

	return plan.NewUpdate(table, table.PrimaryKey().Position, columns, node), nil
}

func (p *Planner) planUpdateColumns(scheme sql.Scheme, stmts []ast.SetStatement) (map[uint8]expr.Node, error) {
	columns := make(map[uint8]expr.Node, len(stmts))

	for i := range stmts {
		column, ok := scheme[stmts[i].Column]
		if !ok {
			return nil, fmt.Errorf("column %q not found", stmts[i].Column)
		}

		value, err := expr.New(stmts[i].Value, scheme)
		if err != nil {
			return nil, fmt.Errorf("failed to create expr from value: %w", err)
		}

		columns[column.Position] = value
	}

	return columns, nil
}

func (p *Planner) planDelete(database string, stmt *ast.DeleteStatement) (plan.Node, error) {
	var (
		table sql.Table
		node  plan.Node
		err   error
	)

	if table, err = p.getTable(database, stmt.Table); err != nil {
		return nil, err
	}

	if node, err = p.planFilter(table, stmt.Where, plan.NewScan(table)); err != nil {
		return nil, fmt.Errorf("failed to plan filter: %w", err)
	}

	return plan.NewDelete(table, table.PrimaryKey().Position, node), nil
}

func (p *Planner) planCreateDatabase(stmt *ast.CreateDatabaseStatement) (plan.Node, error) {
	if _, err := p.catalog.GetDatabase(stmt.Database); err == nil {
		return nil, fmt.Errorf("database %q already exist", stmt.Database)
	}

	return plan.NewCreateDatabase(p.catalog, stmt.Database), nil
}

func (p *Planner) planDropDatabase(stmt *ast.DropDatabaseStatement) (plan.Node, error) {
	if _, err := p.catalog.GetDatabase(stmt.Database); err != nil {
		return nil, fmt.Errorf("failed to get database %q: %w", stmt.Database, err)
	}

	return plan.NewDropDatabase(p.catalog, stmt.Database), nil
}

func (p *Planner) planCreateTable(database string, stmt *ast.CreateTableStatement) (plan.Node, error) {
	db, err := p.catalog.GetDatabase(database)
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if _, err = db.GetTable(stmt.Table); err == nil {
		return nil, fmt.Errorf("table %q already exist", stmt.Table)
	}

	scheme, err := p.planTableScheme(stmt.Columns)
	if err != nil {
		return nil, fmt.Errorf("failed to plan table scheme: %w", err)
	}

	return plan.NewCreateTable(db, stmt.Table, scheme), nil
}

func (p *Planner) planTableScheme(columns []ast.Column) (sql.Scheme, error) {
	primaryKeys := 0
	scheme := make(sql.Scheme, len(columns))

	for i := range columns {
		column, err := p.planSchemeColumn(uint8(i), columns[i])
		if err != nil {
			return nil, err
		}

		if column.PrimaryKey {
			primaryKeys++
		}

		scheme[column.Name] = column
	}

	if primaryKeys == 0 {
		return nil, fmt.Errorf("primary key is required")
	}

	if primaryKeys > 1 {
		return nil, fmt.Errorf("multiple primary keys are not allowed")
	}

	return scheme, nil
}

func (p *Planner) planSchemeColumn(position uint8, column ast.Column) (sql.Column, error) {
	var (
		dataType sql.DataType
		value    sql.Value
	)

	switch column.Type {
	case token.Integer:
		dataType = sql.Integer
	case token.Float:
		dataType = sql.Float
	case token.Text:
		dataType = sql.Text
	case token.Boolean:
		dataType = sql.Boolean
	default:
		return sql.Column{}, fmt.Errorf("unexpected column type: %q", column.Type)
	}

	if column.Default != nil {
		defaultExpr, err := expr.New(column.Default, nil)
		if err != nil {
			return sql.Column{}, fmt.Errorf("failed to create default expr: %w", err)
		}

		value, err = defaultExpr.Eval(nil)
		if err != nil {
			return sql.Column{}, fmt.Errorf("failed to eval default expr: %w", err)
		}

		if dataType != value.DataType() {
			switch {
			case value.DataType() != sql.Null:
				return sql.Column{}, fmt.Errorf("invalid default value for column %q", column.Name)
			case value.DataType() == sql.Null && !column.Nullable:
				return sql.Column{}, fmt.Errorf("null value in column %q violates not-null constraint", column.Name)
			}
		}
	}

	return sql.Column{
		Position:   position,
		Name:       column.Name,
		DataType:   dataType,
		PrimaryKey: column.PrimaryKey,
		Nullable:   column.Nullable,
		Default:    value,
	}, nil
}

func (p *Planner) planDropTable(database string, stmt *ast.DropTableStatement) (plan.Node, error) {
	db, err := p.catalog.GetDatabase(database)
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if _, err = db.GetTable(stmt.Table); err != nil {
		return nil, fmt.Errorf("failed to get table %q: %w", stmt.Table, err)
	}

	return plan.NewDropTable(db, stmt.Table), nil
}

func (p *Planner) planProject(table sql.Table, stmt []ast.ResultStatement, child plan.Node) (plan.Node, error) {
	var (
		scheme      sql.Scheme
		projections []plan.Projection
		err         error
	)

	if len(stmt) == 0 {
		return nil, fmt.Errorf("projections list should be not empty")
	}

	if table != nil {
		scheme = table.Scheme()
	}

	if projections, err = p.planProjections(scheme, stmt); err != nil {
		return nil, err
	}

	return plan.NewProject(projections, child), nil
}

func (p *Planner) planProjections(scheme sql.Scheme, stmt []ast.ResultStatement) ([]plan.Projection, error) {
	projections := make([]plan.Projection, 0, len(stmt))

	for i := range stmt {
		switch stmt[i].Expr.(type) {
		case *ast.AsteriskExpr:
			if scheme == nil {
				return nil, fmt.Errorf("table not specified")
			}

			columns := make([]string, len(scheme))

			for name := range scheme {
				columns[scheme[name].Position] = name
			}

			for position := range columns {
				projections = append(projections, plan.Projection{
					Expr: expr.Column{
						Name:     columns[position],
						Position: uint8(position),
					},
				})
			}
		default:
			node, err := expr.New(stmt[i].Expr, scheme)
			if err != nil {
				return nil, err
			}

			projections = append(projections, plan.Projection{
				Alias: stmt[i].Alias,
				Expr:  node,
			})
		}
	}

	return projections, nil
}

func (p *Planner) planFilter(table sql.Table, stmt *ast.WhereStatement, child plan.Node) (plan.Node, error) {
	if stmt == nil {
		return child, nil
	}

	if table == nil {
		return nil, fmt.Errorf("table not specified")
	}

	cond, err := expr.New(stmt.Expr, table.Scheme())
	if err != nil {
		return nil, err
	}

	return plan.NewFilter(cond, child), nil
}

func (p *Planner) planSort(table sql.Table, stmt *ast.OrderByStatement, child plan.Node) (plan.Node, error) {
	if stmt == nil {
		return child, nil
	}

	if table == nil {
		return nil, fmt.Errorf("table not specified")
	}

	column, ok := table.Scheme()[stmt.Column]
	if !ok {
		return nil, fmt.Errorf("column %s not exists", stmt.Column)
	}

	var order plan.Order

	switch stmt.Direction {
	case token.Asc:
		order = plan.Ascending
	case token.Desc:
		order = plan.Descending
	default:
		return nil, fmt.Errorf("unexpected sort order: %s", stmt.Direction)
	}

	return plan.NewSort(column.Position, order, child), nil
}

func (p *Planner) planOffset(stmt *ast.OffsetStatement, child plan.Node) (plan.Node, error) {
	if stmt == nil {
		return child, nil
	}

	offsetExpr, err := expr.New(stmt.Value, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create offset expr: %w", err)
	}

	limit, err := offsetExpr.Eval(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to eval offset expr: %w", err)
	}

	n, ok := limit.Raw().(int64)
	if !ok {
		return nil, fmt.Errorf("OFFSET expr must be integer type")
	}

	if n < 0 {
		return nil, fmt.Errorf("OFFSET must not be negative")
	}

	return plan.NewOffset(n, child), nil
}

func (p *Planner) planLimit(stmt *ast.LimitStatement, child plan.Node) (plan.Node, error) {
	if stmt == nil {
		return child, nil
	}

	limitExpr, err := expr.New(stmt.Value, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create limit expr: %w", err)
	}

	limit, err := limitExpr.Eval(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to eval limit expr: %w", err)
	}

	n, ok := limit.Raw().(int64)
	if !ok {
		return nil, fmt.Errorf("LIMIT expr must be integer type")
	}

	if n < 0 {
		return nil, fmt.Errorf("LIMIT must not be negative")
	}

	return plan.NewLimit(n, child), nil
}

func (p *Planner) getTable(databaseName, tableName string) (sql.Table, error) {
	if databaseName == "" {
		return nil, fmt.Errorf("database not specified")
	}

	if tableName == "" {
		return nil, fmt.Errorf("table not specified")
	}

	db, err := p.catalog.GetDatabase(databaseName)
	if err != nil {
		return nil, err
	}

	table, err := db.GetTable(tableName)
	if err != nil {
		return nil, err
	}

	return table, nil
}
