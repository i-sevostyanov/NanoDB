package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
)

func TestLexer_NextToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input     string
		tokenType token.Type
		literal   string
	}{
		{
			input:     "&",
			tokenType: token.Illegal,
			literal:   "&",
		},
		{
			input:     "",
			tokenType: token.EOF,
			literal:   "EOF",
		},
		{
			input:     "id",
			tokenType: token.Ident,
			literal:   "id",
		},
		{
			input:     "_id",
			tokenType: token.Ident,
			literal:   "_id",
		},
		{
			input:     "i9",
			tokenType: token.Ident,
			literal:   "i9",
		},
		{
			input:     "    i9  ",
			tokenType: token.Ident,
			literal:   "i9",
		},
		{
			input:     ",",
			tokenType: token.Comma,
			literal:   token.Comma.String(),
		},
		{
			input:     ";",
			tokenType: token.Semicolon,
			literal:   token.Semicolon.String(),
		},
		{
			input:     "(",
			tokenType: token.OpenParen,
			literal:   token.OpenParen.String(),
		},
		{
			input:     ")",
			tokenType: token.CloseParen,
			literal:   token.CloseParen.String(),
		},
		{
			input:     "=",
			tokenType: token.Equal,
			literal:   token.Equal.String(),
		},
		{
			input:     "<",
			tokenType: token.LessThan,
			literal:   token.LessThan.String(),
		},
		{
			input:     ">",
			tokenType: token.GreaterThan,
			literal:   token.GreaterThan.String(),
		},
		{
			input:     "!=",
			tokenType: token.NotEqual,
			literal:   token.NotEqual.String(),
		},
		{
			input:     "!",
			tokenType: token.Not,
			literal:   token.Not.String(),
		},
		{
			input:     "<=",
			tokenType: token.LessThanOrEqual,
			literal:   token.LessThanOrEqual.String(),
		},
		{
			input:     ">=",
			tokenType: token.GreaterThanOrEqual,
			literal:   token.GreaterThanOrEqual.String(),
		},
		{
			input:     "AND",
			tokenType: token.And,
			literal:   token.And.String(),
		},
		{
			input:     "OR",
			tokenType: token.Or,
			literal:   token.Or.String(),
		},
		{
			input:     "NOT",
			tokenType: token.Not,
			literal:   token.Not.String(),
		},
		{
			input:     "+",
			tokenType: token.Add,
			literal:   token.Add.String(),
		},
		{
			input:     "-",
			tokenType: token.Sub,
			literal:   token.Sub.String(),
		},
		{
			input:     "*",
			tokenType: token.Mul,
			literal:   token.Mul.String(),
		},
		{
			input:     "/",
			tokenType: token.Div,
			literal:   token.Div.String(),
		},
		{
			input:     "%",
			tokenType: token.Mod,
			literal:   token.Mod.String(),
		},
		{
			input:     "^",
			tokenType: token.Pow,
			literal:   token.Pow.String(),
		},
		{
			input:     "10",
			tokenType: token.Integer,
			literal:   "10",
		},
		{
			input:     "10.2",
			tokenType: token.Float,
			literal:   "10.2",
		},
		{
			input:     "'value'",
			tokenType: token.Text,
			literal:   "value",
		},
		{
			input:     "true",
			tokenType: token.Boolean,
			literal:   "true",
		},
		{
			input:     "false",
			tokenType: token.Boolean,
			literal:   "false",
		},
		{
			input:     "CREATE",
			tokenType: token.Create,
			literal:   token.Create.String(),
		},
		{
			input:     "TABLE",
			tokenType: token.Table,
			literal:   token.Table.String(),
		},
		{
			input:     "DATABASE",
			tokenType: token.Database,
			literal:   token.Database.String(),
		},
		{
			input:     "DROP",
			tokenType: token.Drop,
			literal:   token.Drop.String(),
		},
		{
			input:     "SELECT",
			tokenType: token.Select,
			literal:   token.Select.String(),
		},
		{
			input:     "AS",
			tokenType: token.As,
			literal:   token.As.String(),
		},
		{
			input:     "FROM",
			tokenType: token.From,
			literal:   token.From.String(),
		},
		{
			input:     "WHERE",
			tokenType: token.Where,
			literal:   token.Where.String(),
		},
		{
			input:     "ORDER",
			tokenType: token.Order,
			literal:   token.Order.String(),
		},
		{
			input:     "BY",
			tokenType: token.By,
			literal:   token.By.String(),
		},
		{
			input:     "ASC",
			tokenType: token.Asc,
			literal:   token.Asc.String(),
		},
		{
			input:     "DESC",
			tokenType: token.Desc,
			literal:   token.Desc.String(),
		},
		{
			input:     "LIMIT",
			tokenType: token.Limit,
			literal:   token.Limit.String(),
		},
		{
			input:     "OFFSET",
			tokenType: token.Offset,
			literal:   token.Offset.String(),
		},
		{
			input:     "INSERT",
			tokenType: token.Insert,
			literal:   token.Insert.String(),
		},
		{
			input:     "INTO",
			tokenType: token.Into,
			literal:   token.Into.String(),
		},
		{
			input:     "VALUES",
			tokenType: token.Values,
			literal:   token.Values.String(),
		},
		{
			input:     "UPDATE",
			tokenType: token.Update,
			literal:   token.Update.String(),
		},
		{
			input:     "SET",
			tokenType: token.Set,
			literal:   token.Set.String(),
		},
		{
			input:     "DELETE",
			tokenType: token.Delete,
			literal:   token.Delete.String(),
		},
		{
			input:     "DEFAULT",
			tokenType: token.Default,
			literal:   token.Default.String(),
		},
		{
			input:     "NULL",
			tokenType: token.Null,
			literal:   token.Null.String(),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			lx := lexer.New(test.input)
			nextToken := lx.NextToken()
			assert.Equal(t, test.tokenType, nextToken.Type)
			assert.Equal(t, test.literal, nextToken.Literal)
		})
	}
}
