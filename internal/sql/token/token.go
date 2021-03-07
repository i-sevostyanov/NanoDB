package token

import (
	"strconv"
	"strings"
)

type Type uint

type Token struct {
	Type    Type
	Literal string
	Offset  int
	Length  int
}

const (
	Illegal Type = iota
	EOF
	Ident

	// Special chars
	Comma      // ,
	OpenParen  // (
	CloseParen // )

	// Comparison operators
	Equal              // =
	LessThan           // <
	GreaterThan        // >
	NotEqual           // !=
	LessThanOrEqual    // <=
	GreaterThanOrEqual // >=

	// Logical operators
	And
	Or
	Not

	// Mathematical operators
	Add // +
	Sub // -
	Mul // *
	Quo // /
	Rem // %
	Pow // ^

	// Types
	Integer
	Float
	String
	Boolean
	Null

	// Constants
	True
	False

	// Keywords
	Create
	Table
	Drop
	Select
	As
	From
	Where
	Order
	By
	Asc
	Desc
	Limit
	Offset
	Insert
	Into
	Values
	Update
	Set
	Delete
)

var tokens = [...]string{
	Illegal: "Illegal",
	EOF:     "EOF",
	Ident:   "Ident",

	Comma:      ",",
	OpenParen:  "(",
	CloseParen: ")",

	Equal:              "=",
	LessThan:           "<",
	GreaterThan:        ">",
	NotEqual:           "!=",
	LessThanOrEqual:    "<=",
	GreaterThanOrEqual: ">=",

	And: "AND",
	Or:  "OR",
	Not: "NOT",

	Add: "+",
	Sub: "-",
	Mul: "*",
	Quo: "/",
	Rem: "%",
	Pow: "^",

	Integer: "INTEGER",
	Float:   "FLOAT",
	String:  "STRING",
	Boolean: "BOOLEAN",
	Null:    "NULL",

	True:  "TRUE",
	False: "FALSE",

	Create: "CREATE",
	Table:  "TABLE",
	Drop:   "DROP",
	Select: "SELECT",
	As:     "AS",
	From:   "FROM",
	Where:  "WHERE",
	Order:  "ORDER",
	By:     "BY",
	Asc:    "ASC",
	Desc:   "DESC",
	Limit:  "LIMIT",
	Offset: "OFFSET",
	Insert: "INSERT",
	Into:   "INTO",
	Values: "VALUES",
	Update: "UPDATE",
	Set:    "SET",
	Delete: "DELETE",
}

func (t Type) String() string {
	s := ""

	if t < Type(len(tokens)) {
		s = tokens[t]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}

	return s
}

func Lookup(ident string) Type {
	keywords := map[string]Type{
		"INTEGER": Integer,
		"FLOAT":   Float,
		"STRING":  String,
		"BOOLEAN": Boolean,
		"CREATE":  Create,
		"TABLE":   Table,
		"DROP":    Drop,
		"SELECT":  Select,
		"AS":      As,
		"FROM":    From,
		"WHERE":   Where,
		"ORDER":   Order,
		"BY":      By,
		"ASC":     Asc,
		"DESC":    Desc,
		"LIMIT":   Limit,
		"OFFSET":  Offset,
		"INSERT":  Insert,
		"INTO":    Into,
		"VALUES":  Values,
		"UPDATE":  Update,
		"SET":     Set,
		"DELETE":  Delete,
		"AND":     And,
		"OR":      Or,
		"NOT":     Not,
		"TRUE":    True,
		"FALSE":   False,
		"NULL":    Null,
	}

	if t, ok := keywords[strings.ToUpper(ident)]; ok {
		return t
	}

	return Ident
}

func New(t Type, offset int) Token {
	literal := t.String()
	length := len(t.String())

	return Token{
		Type:    t,
		Literal: literal,
		Offset:  offset - length,
		Length:  length,
	}
}

const LowestPrecedence = 0

// Precedence returns the operator precedence of the binary operator
func (t Type) Precedence() int {
	switch t {
	case Or:
		return 1
	case And:
		return 2
	case Equal, NotEqual, LessThan, LessThanOrEqual, GreaterThan, GreaterThanOrEqual:
		return 3
	case Add, Sub:
		return 4
	case Mul, Quo, Rem:
		return 5
	case Pow:
		return 6
	default:
		return LowestPrecedence
	}
}
