// Package token defines constants representing the lexical tokens of the NanoDB's SQL dialect.
package token

import (
	"strconv"
	"strings"
)

// Type is the set of lexical tokens.
type Type uint

// Token represents a token or text string returned from the lexer.
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
	Semicolon  // ;
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
	Div // /
	Mod // %
	Pow // ^

	// Types
	Integer
	Float
	String
	Boolean
	Null

	// Keywords
	Create
	Table
	Database
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
	Default
	Primary
	Key
)

var tokens = [...]string{
	Illegal: "Illegal",
	EOF:     "EOF",
	Ident:   "Ident",

	Comma:      ",",
	Semicolon:  ";",
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
	Div: "/",
	Mod: "%",
	Pow: "^",

	Integer: "INTEGER",
	Float:   "FLOAT",
	String:  "STRING",
	Boolean: "BOOLEAN",
	Null:    "NULL",

	Create:   "CREATE",
	Table:    "TABLE",
	Database: "DATABASE",
	Drop:     "DROP",
	Select:   "SELECT",
	As:       "AS",
	From:     "FROM",
	Where:    "WHERE",
	Order:    "ORDER",
	By:       "BY",
	Asc:      "ASC",
	Desc:     "DESC",
	Limit:    "LIMIT",
	Offset:   "OFFSET",
	Insert:   "INSERT",
	Into:     "INTO",
	Values:   "VALUES",
	Update:   "UPDATE",
	Set:      "SET",
	Delete:   "DELETE",
	Default:  "DEFAULT",
	Primary:  "PRIMARY",
	Key:      "KEY",
}

// String returns the string corresponding to the token t.
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

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Type {
	keywords := map[string]Type{
		"INTEGER":  Integer,
		"FLOAT":    Float,
		"STRING":   String,
		"BOOLEAN":  Boolean,
		"TRUE":     Boolean,
		"FALSE":    Boolean,
		"CREATE":   Create,
		"TABLE":    Table,
		"DATABASE": Database,
		"DROP":     Drop,
		"SELECT":   Select,
		"AS":       As,
		"FROM":     From,
		"WHERE":    Where,
		"ORDER":    Order,
		"BY":       By,
		"ASC":      Asc,
		"DESC":     Desc,
		"LIMIT":    Limit,
		"OFFSET":   Offset,
		"INSERT":   Insert,
		"INTO":     Into,
		"VALUES":   Values,
		"UPDATE":   Update,
		"SET":      Set,
		"DELETE":   Delete,
		"AND":      And,
		"OR":       Or,
		"NOT":      Not,
		"NULL":     Null,
		"DEFAULT":  Default,
		"PRIMARY":  Primary,
		"KEY":      Key,
	}

	if t, ok := keywords[strings.ToUpper(ident)]; ok {
		return t
	}

	return Ident
}

// New is a short-hand method for creating the new token.
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

// LowestPrecedence is non-operators precedence.
const LowestPrecedence = 0

// Precedence returns the operator precedence of the binary operator.
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
	case Mul, Div, Mod:
		return 5
	case Pow:
		return 6
	default:
		return LowestPrecedence
	}
}

// IsRightAssociative returns true if the operator is right associative.
func (t Type) IsRightAssociative() bool {
	return t == Pow
}
