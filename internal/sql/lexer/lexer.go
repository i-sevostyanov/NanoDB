package lexer

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql/token"
)

const EOF = rune(0)

type Lexer struct {
	input      string
	ch         rune // current character
	offset     int  // character offset
	peekOffset int  // position after current character
}

func New(input string) *Lexer {
	lx := &Lexer{input: input}
	lx.next()

	return lx
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if isLetter(l.ch) {
		return l.readIdentifier()
	}

	if isDigit(l.ch) {
		return l.readNumber()
	}

	defer l.next()

	switch l.ch {
	case EOF:
		return token.New(token.EOF, l.offset)
	case '\'':
		return l.readString()
	case ',':
		return token.New(token.Comma, l.offset)
	case ';':
		return token.New(token.Semicolon, l.offset)
	case '(':
		return token.New(token.OpenParen, l.offset)
	case ')':
		return token.New(token.CloseParen, l.offset)
	case '=':
		return token.New(token.Equal, l.offset)
	case '+':
		return token.New(token.Add, l.offset)
	case '-':
		return token.New(token.Sub, l.offset)
	case '*':
		return token.New(token.Mul, l.offset)
	case '/':
		return token.New(token.Quo, l.offset)
	case '%':
		return token.New(token.Rem, l.offset)
	case '^':
		return token.New(token.Pow, l.offset)
	case '<':
		if next := l.peek(); next == '=' {
			l.next()
			return token.New(token.LessThanOrEqual, l.offset)
		}

		return token.New(token.LessThan, l.offset)
	case '>':
		if next := l.peek(); next == '=' {
			l.next()
			return token.New(token.GreaterThanOrEqual, l.offset)
		}

		return token.New(token.GreaterThan, l.offset)
	case '!':
		if next := l.peek(); next == '=' {
			l.next()
			return token.New(token.NotEqual, l.offset)
		}

		return token.New(token.Not, l.offset)
	default:
		return l.readIllegal()
	}
}

func (l *Lexer) next() {
	if l.peekOffset < len(l.input) {
		l.offset = l.peekOffset
		l.peekOffset++
		l.ch = rune(l.input[l.offset])
	} else {
		l.offset = len(l.input)
		l.ch = EOF
	}
}

func (l *Lexer) peek() rune {
	if l.peekOffset < len(l.input) {
		return rune(l.input[l.peekOffset])
	}

	return EOF
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.offset

	for isLetter(l.ch) || isDigit(l.ch) {
		l.next()
	}

	literal := l.input[start:l.offset]

	return token.Token{
		Type:    token.Lookup(literal),
		Literal: literal,
		Offset:  start,
		Length:  l.offset - start,
	}
}

func (l *Lexer) readNumber() token.Token {
	start := l.offset
	tokenType := token.Integer

	for isDigit(l.ch) {
		l.next()
	}

	if l.ch == '.' {
		tokenType = token.Float

		l.next()

		for isDigit(l.ch) {
			l.next()
		}
	}

	return token.Token{
		Type:    tokenType,
		Literal: l.input[start:l.offset],
		Offset:  start,
		Length:  l.offset - start,
	}
}

func (l *Lexer) readString() token.Token {
	l.next()
	start := l.offset

	for l.ch != '\'' && !isEOF(l.ch) {
		l.next()
	}

	return token.Token{
		Type:    token.String,
		Literal: l.input[start:l.offset],
		Offset:  start,
		Length:  l.offset - start,
	}
}

func (l *Lexer) readIllegal() token.Token {
	return token.Token{
		Type:    token.Illegal,
		Literal: string(l.ch),
		Offset:  l.offset - 1,
		Length:  1,
	}
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.next()
	}
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r'
}

func isEOF(r rune) bool {
	return r == EOF
}
