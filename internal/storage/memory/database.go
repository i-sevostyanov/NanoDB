package memory

import (
	"errors"
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Database struct {
	name   string
	tables map[string]*Table
}

func NewDatabase(name string) *Database {
	return &Database{
		name:   name,
		tables: make(map[string]*Table),
	}
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) ListTables() []sql.Table {
	tables := make([]sql.Table, 0, len(d.tables))

	for _, t := range d.tables {
		tables = append(tables, t)
	}

	return tables
}

func (d *Database) GetTable(name string) (sql.Table, error) {
	if table, ok := d.tables[name]; ok {
		return table, nil
	}

	return nil, fmt.Errorf("table %q not found", name)
}

func (d *Database) CreateTable(name string, scheme sql.Scheme) (sql.Table, error) {
	if _, ok := d.tables[name]; ok {
		return nil, errors.New("table already exist")
	}

	if len(scheme) == 0 {
		return nil, errors.New("scheme should not be empty")
	}

	table := NewTable(name, scheme)
	d.tables[name] = table

	return table, nil
}

func (d *Database) DropTable(name string) error {
	if _, ok := d.tables[name]; !ok {
		return fmt.Errorf("table %s not found", name)
	}

	delete(d.tables, name)

	return nil
}
