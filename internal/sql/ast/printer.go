package ast

import (
	"bytes"
	"fmt"
)

// Print "pretty-prints" an AST node to output.
func Print(node Node) string {
	if node == nil {
		return "<nil>"
	}

	switch n := node.(type) {
	case *Statements:
		return printTree(n)
	case *SelectStatement:
		return printSelect(n)
	case *ResultStatement:
		return printResult(n)
	case *FromStatement:
		return printFrom(n)
	case *WhereStatement:
		return printWhere(n)
	case *OrderByStatement:
		return printOrderBy(n)
	case *LimitStatement:
		return printLimit(n)
	case *OffsetStatement:
		return printOffset(n)
	case *InsertStatement:
		return printInsert(n)
	case *UpdateStatement:
		return printUpdate(n)
	case *SetStatement:
		return printSet(n)
	case *DeleteStatement:
		return printDelete(n)
	case *CreateDatabaseStatement:
		return printCreateDatabase(n)
	case *DropDatabaseStatement:
		return printDropDatabase(n)
	case *CreateTableStatement:
		return printCreateTable(n)
	case *DropTableStatement:
		return printDropTable(n)
	case *BinaryExpr, *IdentExpr, *UnaryExpr, *ScalarExpr:
		return printExpr(node.(Expression), "", "")
	default:
		return "<unknown>"
	}
}

func printResult(s *ResultStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(printExpr(s.Expr, "", "  "))

	if s.Alias != nil {
		buf.WriteString(fmt.Sprintf(" AS %s", printExpr(s.Alias, "", "")))
	}

	return buf.String()
}

func printDropTable(s *DropTableStatement) string {
	return fmt.Sprintf("DROP TABLE %s", printExpr(s.Table, "", ""))
}

func printCreateTable(s *CreateTableStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("CREATE TABLE ")
	buf.WriteString(printExpr(s.Table, "", ""))
	buf.WriteString("(\n")

	for i := range s.Columns {
		c := s.Columns[i]

		buf.WriteString(printExpr(c.Name, "", " "))
		buf.WriteString(" ")
		buf.WriteString(c.Type.String())

		if c.Default != nil {
			buf.WriteString(" DEFAULT ")
			buf.WriteString(printExpr(c.Default, "", ""))
		}

		if c.NotNull {
			buf.WriteString(" NOT NULL")
		}

		if c.PrimaryKey {
			buf.WriteString(" PRIMARY KEY")
		}

		buf.WriteString("\n")
	}

	buf.WriteString(")\n")

	return buf.String()
}

func printDropDatabase(s *DropDatabaseStatement) string {
	return fmt.Sprintf("DROP DATABASE %s", printExpr(s.Name, "", ""))
}

func printCreateDatabase(s *CreateDatabaseStatement) string {
	return fmt.Sprintf("CREATE DATABASE %s", printExpr(s.Name, "", ""))
}

func printDelete(s *DeleteStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("DELETE FROM\n  %s", printExpr(s.Table, "", "")))

	if s.Where != nil {
		buf.WriteString(printWhere(s.Where))
	}

	return buf.String()
}

func printUpdate(s *UpdateStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("UPDATE\n")
	buf.WriteString("   " + printExpr(s.Table, "", ""))
	buf.WriteString("SET\n")

	for i := range s.Set {
		buf.WriteString(printSet(&s.Set[i]))
		buf.WriteString("\n")
	}

	if s.Where != nil {
		buf.WriteString(printWhere(s.Where))
	}

	return buf.String()
}

func printSet(s *SetStatement) string {
	return fmt.Sprintf("%s = %s", printExpr(s.Column, "", ""), printExpr(s.Value, "", ""))
}

func printInsert(s *InsertStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("INSERT INTO ")
	buf.WriteString(printExpr(s.Table, "", ""))
	buf.WriteString("(\n")

	for i := range s.Columns {
		buf.WriteString("  " + printExpr(s.Columns[i], "", "   "))
	}

	buf.WriteString(")\n")
	buf.WriteString("VALUES\n")
	buf.WriteString("(\n")

	for _, v := range s.Values {
		buf.WriteString("  " + printExpr(v, "", "   "))
	}

	buf.WriteString(")\n")

	return buf.String()
}

func printTree(n *Statements) string {
	buf := bytes.NewBuffer(nil)

	for _, s := range *n {
		buf.WriteString(Print(s))
	}

	return buf.String()
}

func printSelect(s *SelectStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("SELECT\n")

	for i := range s.Result {
		buf.WriteString("  " + printResult(&s.Result[i])) // printExpr(.Expr, "", "  ")
	}

	if s.From != nil {
		buf.WriteString(printFrom(s.From))
	}

	if s.Where != nil {
		buf.WriteString(printWhere(s.Where))
	}

	if s.OrderBy != nil {
		buf.WriteString(printOrderBy(s.OrderBy))
	}

	if s.Limit != nil {
		buf.WriteString(printLimit(s.Limit))
	}

	if s.Offset != nil {
		buf.WriteString(printOffset(s.Offset))
	}

	return buf.String()
}

func printFrom(s *FromStatement) string {
	return fmt.Sprintf("FROM\n  %s", printExpr(s.Table, "", "  "))
}

func printWhere(s *WhereStatement) string {
	return fmt.Sprintf("WHERE\n  %s", printExpr(s.Expr, "", "  "))
}

func printOrderBy(s *OrderByStatement) string {
	return fmt.Sprintf("ORDER BY\n  %s %s", printExpr(s.Column, "", ""), printExpr(s.Order, "", ""))
}

func printLimit(s *LimitStatement) string {
	return fmt.Sprintf("LIMIT\n  %s", printExpr(s.Value, "", ""))
}

func printOffset(s *OffsetStatement) string {
	return fmt.Sprintf("OFFSET\n  %s", printExpr(s.Value, "", ""))
}

func printExpr(expression Expression, prefix, childPrefix string) string {
	switch n := expression.(type) {
	case *BinaryExpr:
		buf := bytes.NewBuffer(nil)
		buf.WriteString(prefix)
		buf.WriteString(n.Operator.String())
		buf.WriteString("\n")

		left := printExpr(n.Left, childPrefix+"├── ", childPrefix+"│   ")
		buf.WriteString(left)

		right := printExpr(n.Right, childPrefix+"└── ", childPrefix+"    ")
		buf.WriteString(right)

		return buf.String()
	case *IdentExpr:
		return fmt.Sprintf("%s%s::ident\n", prefix, n.Name)
	case *UnaryExpr:
		return fmt.Sprintf("%s%s%s", prefix, n.Operator, printExpr(n.Right, "", ""))
	case *ScalarExpr:
		return fmt.Sprintf("%s%s\n", prefix, n.Literal)
	default:
		return fmt.Sprintf("%s<unexpected>\n", prefix)
	}
}
