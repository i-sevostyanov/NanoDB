package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/i-sevostyanov/NanoDB/internal/repl"
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

	parseFn := engine.ParseFn(func(sql string) (ast.Node, error) {
		lx := lexer.New(sql)
		pr := parser.New(lx)
		return pr.Parse()
	})

	catalog := memory.NewCatalog()
	aPlanner := planner.New(catalog)
	anEngine := engine.New(parseFn, aPlanner)
	aRepl := repl.New(os.Stdin, os.Stdout, catalog, anEngine)

	if err := aRepl.Run(ctx); err != nil {
		switch {
		case errors.Is(err, repl.ErrQuit):
		default:
			log.Printf("repl: %v\n", err)
		}
	}
}
