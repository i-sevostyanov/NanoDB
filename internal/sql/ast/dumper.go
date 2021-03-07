package ast

import (
	"bytes"
	"fmt"
)

func Dump(node Node) string {
	if node == nil {
		return "<nil>"
	}

	switch n := node.(type) {
	case *Tree:
		return dumpTree(n)
	case *BadStatement:
		return fmt.Sprintf("BadStatement: %s::%s\n", n.Literal, n.Type.String())
	case *SelectStatement:
		return dumpSelect(n)
	case *FromStatement:
		return fmt.Sprintf("FROM\n  %s", Dump(n.Table))
	case *WhereStatement:
		return fmt.Sprintf("WHERE\n  %s", dumpExpr(n.Expr, "", "  "))
	case *OrderByStatement:
		return fmt.Sprintf("ORDER BY\n  %s  %s", Dump(n.Column), Dump(n.Order))
	case *LimitStatement:
		return fmt.Sprintf("LIMIT\n  %s", Dump(n.Value))
	case *OffsetStatement:
		return fmt.Sprintf("OFFSET\n  %s", Dump(n.Value))
	case *InsertStatement:
		return "InsertStatement"
	case *SetStatement:
		return "SetStatement"
	case *DeleteStatement:
		return "DeleteStatement"
	case *CreateTableStatement:
		return "CreateTableStatement"
	case *DropTableStatement:
		return "DropTableStatement"
	case *BinaryExpr, *IdentExpr, *UnaryExpr, *BadExpr, *ScalarExpr:
		return dumpExpr(node.(Expression), "", "")
	default:
		return "<unknown>"
	}
}

func dumpTree(n *Tree) string {
	buf := bytes.NewBuffer(nil)

	for _, s := range n.Statements {
		buf.WriteString(Dump(s))
	}

	return buf.String()
}

func dumpSelect(s *SelectStatement) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("SELECT\n")

	for i := range s.Result {
		buf.WriteString("  " + dumpExpr(s.Result[i].Expr, "", "  "))
	}

	if s.From != nil {
		buf.WriteString(Dump(s.From))
	}

	if s.Where != nil {
		buf.WriteString(Dump(s.Where))
	}

	if s.OrderBy != nil {
		buf.WriteString(Dump(s.OrderBy))
	}

	if s.Limit != nil {
		buf.WriteString(Dump(s.Limit))
	}

	if s.Offset != nil {
		buf.WriteString(Dump(s.Offset))
	}

	return buf.String()
}

func dumpExpr(expression Expression, prefix, childPrefix string) string {
	switch n := expression.(type) {
	case *BinaryExpr:
		buf := bytes.NewBuffer(nil)
		buf.WriteString(prefix)
		buf.WriteString(n.Operator.String())
		buf.WriteString("\n")

		left := dumpExpr(n.Left, childPrefix+"├── ", childPrefix+"│   ")
		buf.WriteString(left)

		right := dumpExpr(n.Right, childPrefix+"└── ", childPrefix+"    ")
		buf.WriteString(right)

		return buf.String()
	case *IdentExpr:
		return fmt.Sprintf("%s%s::ident\n", prefix, n.Name)
	case *UnaryExpr:
		return fmt.Sprintf("%s%s%s", prefix, n.Operator, dumpExpr(n.Right, "", ""))
	case *BadExpr:
		return fmt.Sprintf("%s%s:%s::badexpr\n", prefix, n.Type.String(), n.Literal)
	case *ScalarExpr:
		return fmt.Sprintf("%s%s\n", prefix, n.Literal)
	default:
		return fmt.Sprintf("%s<unexpected>\n", prefix)
	}
}
