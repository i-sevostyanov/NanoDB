package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/i-sevostyanov/NanoDB/internal/sql/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parser"
)

func main() {
	go func() {
		for {
			fmt.Print("#> ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()

			fmt.Print(ast.Print(tree))

			for _, err := range errors {
				fmt.Printf("%v\n", err)
			}
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stop)

	<-stop
}
