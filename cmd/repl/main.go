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

	"golang.org/x/sync/errgroup"
)

func main() {
	errCanceled := errors.New("canceled")

	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		parseFn := engine.ParseFn(func(sql string) (ast.Node, error) {
			lx := lexer.New(sql)
			pr := parser.New(lx)
			return pr.Parse()
		})

		catalog := memory.NewCatalog()
		aPlanner := planner.New(catalog)
		anEngine := engine.New(parseFn, aPlanner)
		aRepl := repl.New(os.Stdin, os.Stdout, catalog, anEngine)

		return aRepl.Run(ctx)
	})

	gr.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(signals)

		select {
		case <-ctx.Done():
			return nil
		case <-signals:
			return errCanceled
		}
	})

	if err := gr.Wait(); err != nil {
		switch {
		case errors.Is(err, errCanceled), errors.Is(err, repl.ErrQuit):
		default:
			log.Fatalf("Failed to start: %v", err)
		}
	}
}
