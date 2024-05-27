package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/i-sevostyanov/NanoDB/internal/shell"
	"github.com/i-sevostyanov/NanoDB/internal/sql/engine"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/parser"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/planner"
	"github.com/i-sevostyanov/NanoDB/internal/storage/memory"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	sqlParser := engine.ParseFn(func(sql string) (ast.Node, error) {
		lx := lexer.New(sql)
		pr := parser.New(lx)
		return pr.Parse()
	})

	sqlCatalog := memory.NewCatalog()
	sqlPlanner := planner.New(sqlCatalog)
	sqlEngine := engine.New(sqlParser, sqlPlanner)
	tableWriter := shell.NewTableWriterFactory()

	sh := shell.New(os.Stdin, os.Stdout, sqlCatalog, sqlEngine, tableWriter)
	sh.Run(ctx)
}
