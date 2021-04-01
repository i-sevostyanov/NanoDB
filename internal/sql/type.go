package sql

// Column is the definition of a table column.
type Column struct {
	Position   uint8
	Name       string
	DataType   DataType
	PrimaryKey bool
	Nullable   bool
	Default    interface{}
}

// Scheme is the definition of a table.
type Scheme map[string]Column

// Sequence returns a sequentially increasing value every time you call Next.
type Sequence interface {
	Next() int64
}

// Catalog holds meta-information about databases.
type Catalog interface {
	GetDatabase(name string) (Database, error)
	ListDatabases() ([]Database, error)
	CreateDatabase(name string) (Database, error)
	DropDatabase(name string) error
}

type Database interface {
	Name() string
	CreateTable(name string, scheme Scheme) error
	DropTable(name string) error
	Tables() []Table
	Table(name string) (Table, error)
}

// Table represents the backend of a SQL table.
type Table interface {
	Name() string
	Scheme() Scheme
	RowIter() (RowIter, error)
	Insert(row Row) (Row, error)

	// Insert(columns []string, values []Raw) (RowIter, error)
	// Update(columns map[string]Raw, filter Expr) (RowIter, error)
	// Delete(filter Expr) (RowIter, error)
}
