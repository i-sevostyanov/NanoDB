package plan

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate mockgen -source=drop.go -destination ./drop_mock_test.go -package plan_test

type DatabaseDropper interface {
	DropDatabase(name string) error
}

type DropDatabase struct {
	dropper DatabaseDropper
	name    string
}

func NewDropDatabase(dropper DatabaseDropper, name string) *DropDatabase {
	return &DropDatabase{
		dropper: dropper,
		name:    name,
	}
}

func (d *DropDatabase) Columns() []string {
	return nil
}

func (d *DropDatabase) RowIter() (sql.RowIter, error) {
	if err := d.dropper.DropDatabase(d.name); err != nil {
		return nil, fmt.Errorf("failed to drop database: %w", err)
	}

	return sql.RowsIter(), nil
}

type TableDropper interface {
	DropTable(name string) error
}

type DropTable struct {
	dropper TableDropper
	name    string
}

func (d *DropTable) Columns() []string {
	return nil
}

func NewDropTable(dropper TableDropper, name string) *DropTable {
	return &DropTable{
		dropper: dropper,
		name:    name,
	}
}

func (d *DropTable) RowIter() (sql.RowIter, error) {
	if err := d.dropper.DropTable(d.name); err != nil {
		return nil, fmt.Errorf("failed to drop table: %w", err)
	}

	return sql.RowsIter(), nil
}
