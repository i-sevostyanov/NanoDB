package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/parser"
)

func main() {
	go func() {
		for {
			fmt.Print("#> ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			p := parser.New(lexer.New(input))

			tree, err := p.Parse()
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Print(ast.Print(tree))
			}
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stop)

	<-stop
}
