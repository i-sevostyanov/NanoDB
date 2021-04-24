package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

type Projection struct {
	Alias string
	Expr  expr.Node
}

type Project struct {
	projections []Projection
	child       Node
}

func NewProject(projections []Projection, child Node) *Project {
	return &Project{
		projections: projections,
		child:       child,
	}
}

func (p *Project) Columns() []string {
	columns := make([]string, 0, len(p.projections))

	for i := range p.projections {
		column := p.projections[i].Alias

		if column == "" {
			column = p.projections[i].Expr.String()
		}

		columns = append(columns, column)
	}

	return columns
}

func (p *Project) RowIter() (sql.RowIter, error) {
	iter, err := p.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("failed to get row iter: %w", err)
	}

	projections := make([]expr.Node, 0, len(p.projections))

	for i := range p.projections {
		projections = append(projections, p.projections[i].Expr)
	}

	iter = &projectIter{
		projections: projections,
		iter:        iter,
	}

	return iter, nil
}

type projectIter struct {
	projections []expr.Node
	iter        sql.RowIter
}

func (i *projectIter) Next() (sql.Row, error) {
	row, err := i.iter.Next()
	switch {
	case errors.Is(err, io.EOF):
		return nil, err
	case err != nil:
		return nil, fmt.Errorf("failed to get next row: %w", err)
	default:
		return i.projection(row)
	}
}

func (i *projectIter) Close() error {
	return i.iter.Close()
}

func (i *projectIter) projection(row sql.Row) (sql.Row, error) {
	project := make(sql.Row, 0, len(i.projections))

	for _, p := range i.projections {
		value, err := p.Eval(row)
		if err != nil {
			return nil, fmt.Errorf("failed to eval projection: %w", err)
		}

		project = append(project, value)
	}

	return project, nil
}
