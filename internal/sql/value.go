package sql

//go:generate mockgen -source=value.go -destination ./value_mock.go -package sql

type Value interface {
	Raw() any
	String() string
	DataType() DataType
}
