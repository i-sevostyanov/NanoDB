package sql

//go:generate go run go.uber.org/mock/mockgen -typed -source=value.go -destination ./value_mock.go -package sql

type Value interface {
	Raw() any
	String() string
	DataType() DataType
}
