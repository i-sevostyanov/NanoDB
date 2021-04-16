package engine

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

//go:generate mockgen -source=engine.go -destination ./engine_mock_test.go -package engine_test

type Parser interface {
	Parse(sql string) (ast.Node, error)
}

type Planner interface {
	Plan(database string, node ast.Node) (plan.Node, error)
}

type ParseFn func(sql string) (ast.Node, error)

func (fn ParseFn) Parse(sql string) (ast.Node, error) {
	return fn(sql)
}

type Engine struct {
	parser  Parser
	planner Planner
}

func New(parser Parser, planner Planner) *Engine {
	return &Engine{
		parser:  parser,
		planner: planner,
	}
}

func (e *Engine) Exec(database, input string) (sql.RowIter, error) {
	astNode, err := e.parser.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sql query: %w", err)
	}

	planNode, err := e.planner.Plan(database, astNode)
	if err != nil {
		return nil, fmt.Errorf("failed to build query plan: %w", err)
	}

	return planNode.RowIter()
}
