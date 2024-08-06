package shell

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

const prompt = "#> "

type TableWriter interface {
	WriteTable(headers []string, data [][]string, showRowsCount bool) string
}

type Engine interface {
	Exec(database, sql string) (columns []string, iter sql.RowIter, err error)
}

// Shell is terminal-based front-end to NanoDB.
type Shell struct {
	input    io.Reader
	output   io.Writer
	catalog  sql.Catalog
	database sql.Database
	engine   Engine
	tw       TableWriter
	prompt   string
	closeCh  chan struct{}
}

func New(in io.Reader, out io.Writer, catalog sql.Catalog, engine Engine, tw TableWriter) *Shell {
	return &Shell{
		input:   in,
		output:  out,
		catalog: catalog,
		engine:  engine,
		tw:      tw,
		prompt:  prompt,
		closeCh: make(chan struct{}),
	}
}

func (s *Shell) Run(ctx context.Context) {
	s.write("shell is the NanoDB interactive terminal.\n")

	go func() {
		for {
			s.write(s.prompt)

			scanner := bufio.NewScanner(s.input)
			scanner.Scan()
			input := scanner.Text()

			if input == "" {
				continue
			}

			if reply, err := s.exec(input); err != nil {
				s.write(err.Error() + "\n")
			} else if reply != "" {
				s.write(reply)
			}
		}
	}()

	select {
	case <-ctx.Done():
	case <-s.closeCh:
	}
}

func (s *Shell) write(line string) {
	_, _ = s.output.Write([]byte(line))
}

func (s *Shell) exec(input string) (string, error) {
	switch input[0] {
	case '\\':
		return s.execCommand(input)
	default:
		return s.execQuery(input)
	}
}

func (s *Shell) execCommand(input string) (string, error) {
	cmd := strings.TrimSpace(input)
	params := strings.Fields(cmd)

	switch params[0] {
	case `\use`:
		return s.useDatabase(params)
	case `\databases`:
		return s.listDatabases()
	case `\tables`:
		return s.listTables()
	case `\describe`:
		return s.describeTable(params)
	case `\import`:
		return s.importFile(params)
	case `\help`:
		return s.showHelp(), nil
	case `\quit`:
		return s.quit(), nil
	default:
		return "", fmt.Errorf("unknown command: %v", params[0])
	}
}

func (s *Shell) useDatabase(params []string) (string, error) {
	if len(params) < 2 {
		return "", errors.New("database name not specified")
	}

	db, err := s.catalog.GetDatabase(params[1])
	if err != nil {
		return "", err
	}

	s.database = db
	s.prompt = fmt.Sprintf("%s %s", db.Name(), prompt)

	return "database changed\n", nil
}

func (s *Shell) listDatabases() (string, error) {
	databases, err := s.catalog.ListDatabases()
	if err != nil {
		return "", err
	}

	data := make([][]string, 0, len(databases))

	for i := range databases {
		data = append(data, []string{databases[i].Name()})
	}

	table := s.tw.WriteTable([]string{"Database"}, data, true)

	return table, nil
}

func (s *Shell) listTables() (string, error) {
	if s.database == nil {
		return "", errors.New("connect to database first")
	}

	tables := s.database.ListTables()
	data := make([][]string, 0, len(tables))

	for i := range tables {
		data = append(data, []string{tables[i].Name()})
	}

	table := s.tw.WriteTable([]string{"Table"}, data, true)

	return table, nil
}

func (s *Shell) describeTable(params []string) (string, error) {
	if s.database == nil {
		return "", errors.New("connect to database first")
	}

	if len(params) < 2 {
		return "", errors.New("table name not specified")
	}

	table, err := s.database.GetTable(params[1])
	if err != nil {
		return "", err
	}

	scheme := table.Scheme()
	primaryKey := table.PrimaryKey()
	columns := make([]sql.Column, len(scheme))

	for name := range scheme {
		columns[scheme[name].Position] = sql.Column{
			Position:   scheme[name].Position,
			Name:       name,
			DataType:   scheme[name].DataType,
			PrimaryKey: scheme[name].PrimaryKey,
			Nullable:   scheme[name].Nullable,
			Default:    scheme[name].Default,
		}
	}

	buf := bytes.NewBuffer(nil)
	data := make([][]string, 0, len(columns))

	for i := range columns {
		var defaultValue string

		if columns[i].Default != nil {
			defaultValue = columns[i].Default.String()
		}

		row := []string{
			columns[i].Name,
			columns[i].DataType.String(),
			strconv.FormatBool(columns[i].Nullable),
			defaultValue,
		}

		data = append(data, row)
	}

	tb := s.tw.WriteTable([]string{"Column", "Type", "Nullable", "Default"}, data, false)
	buf.WriteString(tb)
	buf.WriteString("Indexes:\n")
	buf.WriteString(fmt.Sprintf("   PRIMARY KEY (%s) autoincrement\n\n", primaryKey.Name))

	return buf.String(), nil
}

func (s *Shell) importFile(params []string) (string, error) {
	if len(params) < 2 {
		return "", errors.New("filename not specified")
	}

	data, err := os.ReadFile(params[1])
	if err != nil {
		return "", err
	}

	stmts := strings.Split(string(data), ";")

	for i := range stmts {
		stmt := strings.TrimSpace(stmts[i])
		if stmt == "" {
			continue
		}

		if _, err = s.exec(stmt); err != nil {
			return "", err
		}
	}

	return "OK\n", nil
}

func (s *Shell) showHelp() string {
	help := `shell is the NanoDB interactive terminal.

Commands:
  \use <database>                  Use specified database
  \databases                       List databases
  \tables                          List tables
  \describe <table>                Show table definition
  \import <absolute path to file>  Import from file
  \help                            Show help
  \quit                            Quit
`

	return help
}

func (s *Shell) quit() string {
	close(s.closeCh)

	return "Bye!\n"
}

func (s *Shell) execQuery(input string) (string, error) {
	var database string

	if s.database != nil {
		database = s.database.Name()
	}

	columns, rowIter, err := s.engine.Exec(database, input)
	if err != nil {
		return "", fmt.Errorf("execute query: %w", err)
	}

	var data [][]string

loop:
	for {
		var row sql.Row

		row, err = rowIter.Next()
		switch {
		case errors.Is(err, io.EOF):
			break loop
		case err != nil:
			return "", err
		}

		values := make([]string, 0, len(row))

		for i := range row {
			values = append(values, row[i].String())
		}

		data = append(data, values)
	}

	if err = rowIter.Close(); err != nil {
		return "", err
	}

	if len(data) > 0 {
		return s.tw.WriteTable(columns, data, true), nil
	}

	return "", nil
}
