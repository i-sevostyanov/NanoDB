package sql

//go:generate mockgen -source=type.go -destination ./type_mock.go -package sql

// Catalog holds meta-information about databases.
type Catalog interface {
	GetDatabase(name string) (Database, error)
	ListDatabases() ([]Database, error)
	CreateDatabase(name string) (Database, error)
	DropDatabase(name string) error
}

// Database represents the backend of an SQL database.
type Database interface {
	Name() string
	GetTable(name string) (Table, error)
	ListTables() []Table
	CreateTable(name string, scheme Scheme) (Table, error)
	DropTable(name string) error
}

// Table represents the backend of an SQL table.
type Table interface {
	Name() string
	Scheme() Scheme
	PrimaryKey() Column
	Sequence() Sequence
	Scan() (RowIter, error)
	Insert(key int64, row Row) error
	Delete(key int64) error
	Update(key int64, row Row) error
}

// Sequence returns a sequentially increasing value every time you call Next.
type Sequence interface {
	Next() int64
}

// Scheme is the definition of a table (column-name => definition).
type Scheme map[string]Column

// Column is the definition of a table column.
type Column struct {
	Position   uint8
	Name       string
	DataType   DataType
	PrimaryKey bool
	Nullable   bool
	Default    Value
}

type CompareType int

const (
	Less    CompareType = -1
	Equal   CompareType = 0
	Greater CompareType = 1
)
