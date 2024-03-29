package repl

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/i-sevostyanov/NanoDB/internal/sql"

	"github.com/olekukonko/tablewriter"
)

const prompt = "#> "

var ErrQuit = errors.New("quit")

type Engine interface {
	Exec(database, sql string) (columns []string, iter sql.RowIter, err error)
}

// Repl is terminal-based front-end to NanoDB.
type Repl struct {
	input    io.Reader
	output   io.Writer
	catalog  sql.Catalog
	database sql.Database
	engine   Engine
	prompt   string
	closeCh  chan struct{}
}

func New(in io.Reader, out io.Writer, catalog sql.Catalog, engine Engine) *Repl {
	return &Repl{
		input:   in,
		output:  out,
		catalog: catalog,
		engine:  engine,
		prompt:  prompt,
		closeCh: make(chan struct{}),
	}
}

func (r *Repl) Run(ctx context.Context) error {
	r.write("repl is the NanoDB interactive terminal.\n")

	go func() {
		for {
			r.write(r.prompt)

			scanner := bufio.NewScanner(r.input)
			scanner.Scan()
			input := scanner.Text()

			if input == "" {
				continue
			}

			if repl, err := r.exec(input); err != nil {
				r.write(err.Error() + "\n")
			} else if repl != "" {
				r.write(repl)
			}
		}
	}()

	select {
	case <-ctx.Done():
		return nil
	case <-r.closeCh:
		return ErrQuit
	}
}

func (r *Repl) write(s string) {
	_, _ = r.output.Write([]byte(s))
}

func (r *Repl) exec(input string) (string, error) {
	switch input[0] {
	case '\\':
		return r.execCommand(input)
	default:
		return r.execQuery(input)
	}
}

func (r *Repl) execCommand(input string) (string, error) {
	cmd := strings.TrimSpace(input)
	params := strings.Fields(cmd)

	switch params[0] {
	case `\use`:
		return r.useDatabase(params)
	case `\databases`:
		return r.listDatabases()
	case `\tables`:
		return r.listTables()
	case `\describe`:
		return r.describeTable(params)
	case `\import`:
		return r.importFile(params)
	case `\help`:
		return r.showHelp(), nil
	case `\quit`:
		return r.quit(), nil
	default:
		return "", fmt.Errorf("unknown command: %v", params[0])
	}
}

func (r *Repl) useDatabase(params []string) (string, error) {
	if len(params) < 2 {
		return "", fmt.Errorf("database name not specified")
	}

	db, err := r.catalog.GetDatabase(params[1])
	if err != nil {
		return "", err
	}

	r.database = db
	r.prompt = fmt.Sprintf("%s %s", db.Name(), prompt)

	return "database changed\n", nil
}

func (r *Repl) listDatabases() (string, error) {
	databases, err := r.catalog.ListDatabases()
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	data := make([][]string, 0, len(databases))

	for i := range databases {
		data = append(data, []string{databases[i].Name()})
	}

	r.drawTable(buf, []string{"Database"}, data)
	buf.WriteString(fmt.Sprintf("(%d rows)\n\n", len(data)))

	return buf.String(), nil
}

func (r *Repl) listTables() (string, error) {
	if r.database == nil {
		return "", fmt.Errorf("connect to database first")
	}

	buf := bytes.NewBuffer(nil)
	tables := r.database.ListTables()
	data := make([][]string, 0, len(tables))

	for i := range tables {
		data = append(data, []string{tables[i].Name()})
	}

	r.drawTable(buf, []string{"Table"}, data)
	buf.WriteString(fmt.Sprintf("(%d rows)\n\n", len(data)))

	return buf.String(), nil
}

func (r *Repl) describeTable(params []string) (string, error) {
	if r.database == nil {
		return "", fmt.Errorf("connect to database first")
	}

	if len(params) < 2 {
		return "", fmt.Errorf("table name not specified")
	}

	table, err := r.database.GetTable(params[1])
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
			fmt.Sprintf("%t", columns[i].Nullable),
			defaultValue,
		}

		data = append(data, row)
	}

	r.drawTable(buf, []string{"Column", "Type", "Nullable", "Default"}, data)
	buf.WriteString("Indexes:\n")
	buf.WriteString(fmt.Sprintf("   PRIMARY KEY (%s) autoincrement\n\n", primaryKey.Name))

	return buf.String(), nil
}

func (r *Repl) importFile(params []string) (string, error) {
	if len(params) < 2 {
		return "", fmt.Errorf("filename not specified")
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

		if _, err = r.exec(stmt); err != nil {
			return "", err
		}
	}

	return "OK\n", nil
}

func (r *Repl) showHelp() string {
	help := `repl is the NanoDB interactive terminal.

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

func (r *Repl) quit() string {
	close(r.closeCh)
	return "Bye!\n"
}

func (r *Repl) execQuery(input string) (string, error) {
	var database string

	if r.database != nil {
		database = r.database.Name()
	}

	columns, rowIter, err := r.engine.Exec(database, input)
	if err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
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

	buf := bytes.NewBuffer(nil)

	if len(data) > 0 {
		r.drawTable(buf, columns, data)
		buf.WriteString(fmt.Sprintf("(%d rows)\n\n", len(data)))
	}

	return buf.String(), nil
}

func (r *Repl) drawTable(w io.Writer, headers []string, data [][]string) {
	tw := tablewriter.NewWriter(w)
	tw.SetColWidth(75)
	tw.AppendBulk(data)
	tw.SetAutoFormatHeaders(false)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)
	tw.SetHeader(headers)
	tw.Render()
}
