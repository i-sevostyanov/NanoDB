package plan

import (
	"errors"
	"fmt"
	"io"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

type Project struct {
	projections []expr.Node
	child       Node
}

func NewProject(projections []expr.Node, child Node) *Project {
	return &Project{
		projections: projections,
		child:       child,
	}
}

func (p *Project) RowIter() (sql.RowIter, error) {
	iter, err := p.child.RowIter()
	if err != nil {
		return nil, fmt.Errorf("failed to get row iter: %w", err)
	}

	iter = &projectIter{
		projections: p.projections,
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
			return nil, fmt.Errorf("failted to eval projection: %w", err)
		}

		project = append(project, value)
	}

	return project, nil
}
