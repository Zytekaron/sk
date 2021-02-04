package token

import "fmt"

type Type string

type Token struct {
	Type  Type
	Value string
}

func New(t Type, value string) *Token {
	return &Token{
		Type:  t,
		Value: value,
	}
}

func NewType(t Type) *Token {
	return New(t, string(t))
}

func (t *Token) String() string {
	if string(t.Type) == t.Value {
		return t.Value
	}
	return fmt.Sprintf("%s(%s)", t.Type, t.Value)
}

const (
	// Inbuilt Types
	BYTE   = "Byte"
	SHORT  = "Short"
	INT    = "Int"
	LONG   = "Long"
	FLOAT  = "Float"
	DOUBLE = "Double"
	BOOL   = "Bool"
	STRING = "String"

	IDENTIFIER = "Identifier"
	KEYWORD    = "Keyword"

	// Generic Syntax
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	QUOTE      = "\""
	APOSTROPHE = "'"
	BACKTICK   = "`"

	QUESTION  = "?"
	COLON     = ":"
	SEMICOLON = ";"

	DOT       = "."
	TWO_DOT   = ".."
	THREE_DOT = "..."
	COMMA     = ","

	EOF = "EOF"

	// Operators
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"
	DIVIDE = "/"
	MODULO = "%"
	POW    = "**"

	BITWISE_AND    = "&"
	BITWISE_OR     = "|"
	BITWISE_NOT    = "~"
	BITWISE_XOR    = "^"
	BITWISE_LSHIFT = "<<"
	BITWISE_RSHIFT = ">>"
	// todo lshift2?

	// Comparison
	NOT            = "!"
	AND            = "&&"
	OR             = "||"
	LESS           = "<"
	GREATER        = ">"
	LESS_EQUALS    = "<="
	GREATER_EQUALS = ">="
	EQUALS         = "=="
	NOT_EQUALS     = "!="

	// Assignment
	EQ        = "="
	PLUS_EQ   = "+="
	MINUS_EQ  = "-="
	TIMES_EQ  = "*="
	DIVIDE_EQ = "/="
	MODULO_EQ = "%="
	AND_EQ    = "&="
	OR_EQ     = "|="
	NOT_EQ    = "~="
	XOR_EQ    = "^="
	// todo bitwise shift eq?
	PLUS_PLUS   = "++"
	MINUS_MINUS = "--"
)

var keywords = []string{
	"case",
	"const",
	"default",
	"else",
	"fn",
	"for",
	"if",
	"in",
	"let",
	"return",
	"switch",
}

var primitives = []string{
	"null",
	"true",
	"false",
}

var equals = []string{
	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"&=",
	"|=",
	"~=",
	"^=",
}

func InEquals(token *Token) bool {
	for _, e := range equals {
		if e == token.Value {
			return true
		}
	}
	return false
}
