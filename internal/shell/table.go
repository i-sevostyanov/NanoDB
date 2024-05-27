package shell

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Table struct {
	w *tabwriter.Writer
}

func NewTableWriter(w io.Writer) *Table {
	return &Table{
		w: tabwriter.NewWriter(w, 0, 0, 0, ' ', tabwriter.TabIndent),
	}
}

func (t *Table) WriteTable(headers []string, data [][]string, showRowsCount bool) {
	divLine := t.formatDividerLine(headers, data)

	t.write(divLine)
	t.write(t.formatRow(headers))
	t.write(divLine)

	for _, row := range data {
		t.write(t.formatRow(row))
	}

	t.write(divLine)

	if showRowsCount {
		t.write(t.formatRowsCount(data))
	}

	t.flush()
}

func (t *Table) formatDividerLine(headers []string, data [][]string) string {
	columnsWidth := t.columnsWidth(headers, data)
	columns := make([]string, len(headers))

	for i, size := range columnsWidth {
		columns[i] = strings.Repeat("-", size+2)
	}

	return fmt.Sprintf("+%s\t+\n", strings.Join(columns, "\t+"))
}

func (t *Table) formatRow(columns []string) string {
	return fmt.Sprintf("| %s\t|\n", strings.Join(columns, "\t| "))
}

func (t *Table) formatRowsCount(data [][]string) string {
	return fmt.Sprintf("(%d rows)\n\n", len(data))
}

func (t *Table) write(line string) {
	_, _ = t.w.Write([]byte(line))
}

func (t *Table) flush() {
	_ = t.w.Flush()
}

func (t *Table) columnsWidth(headers []string, data [][]string) []int {
	columns := make([]int, len(headers))

	for i := range data {
		for j := range data[i] {
			if len(headers[j])+2 > columns[j] {
				columns[j] = len(headers[j])
			}

			if len(data[i][j]) > columns[j] {
				columns[j] = len(data[i][j])
			}
		}
	}

	return columns
}

type TableWriterFactory struct{}

func NewTableWriterFactory() TableWriterFactory {
	return TableWriterFactory{}
}

func (f TableWriterFactory) WriteTable(w io.Writer, headers []string, data [][]string, showRowsCount bool) {
	NewTableWriter(w).WriteTable(headers, data, showRowsCount)
}
