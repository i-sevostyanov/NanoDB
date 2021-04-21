package memory

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type Catalog struct {
	databases map[string]sql.Database
}

func NewCatalog() *Catalog {
	return &Catalog{
		databases: make(map[string]sql.Database),
	}
}

func (c *Catalog) GetDatabase(name string) (sql.Database, error) {
	if database, ok := c.databases[name]; ok {
		return database, nil
	}

	return nil, fmt.Errorf("database %q not found", name)
}

func (c *Catalog) ListDatabases() ([]sql.Database, error) {
	databases := make([]sql.Database, 0, len(c.databases))

	for name := range c.databases {
		databases = append(databases, c.databases[name])
	}

	return databases, nil
}

func (c *Catalog) CreateDatabase(name string) (sql.Database, error) {
	if _, ok := c.databases[name]; ok {
		return nil, fmt.Errorf("database %q already exist", name)
	}

	database := NewDatabase(name)
	c.databases[name] = database

	return database, nil
}

func (c *Catalog) DropDatabase(name string) error {
	if _, ok := c.databases[name]; ok {
		delete(c.databases, name)
		return nil
	}

	return fmt.Errorf("database %q not found", name)
}
