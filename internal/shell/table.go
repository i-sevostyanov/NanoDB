package shell

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Table struct{}

func NewTableWriter() Table {
	return Table{}
}

func (t Table) WriteTable(headers []string, data [][]string, showRowsCount bool) string {
	divLine := t.formatDividerLine(headers, data)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 0, ' ', tabwriter.TabIndent)

	t.write(w, divLine)
	t.write(w, t.formatRow(headers))
	t.write(w, divLine)

	for _, row := range data {
		t.write(w, t.formatRow(row))
	}

	t.write(w, divLine)

	if showRowsCount {
		t.write(w, t.formatRowsCount(data))
	}

	_ = w.Flush()

	return buf.String()
}

func (t Table) formatDividerLine(headers []string, data [][]string) string {
	columnsWidth := t.columnsWidth(headers, data)
	columns := make([]string, len(headers))

	for i, size := range columnsWidth {
		columns[i] = strings.Repeat("-", size+2)
	}

	return fmt.Sprintf("+%s\t+\n", strings.Join(columns, "\t+"))
}

func (t Table) formatRow(columns []string) string {
	return fmt.Sprintf("| %s\t|\n", strings.Join(columns, "\t| "))
}

func (t Table) formatRowsCount(data [][]string) string {
	return fmt.Sprintf("(%d rows)\n\n", len(data))
}

func (t Table) write(w io.Writer, line string) {
	_, _ = w.Write([]byte(line))
}

func (t Table) columnsWidth(headers []string, data [][]string) []int {
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
