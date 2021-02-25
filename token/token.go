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
	BYTE     = "Byte"
	SHORT    = "Short"
	INT      = "Int"
	LONG     = "Long"
	BIGINT   = "BigInt"
	FLOAT    = "Float"
	DOUBLE   = "Double"
	BIGFLOAT = "BigFloat"
	BOOL     = "Bool"
	STRING   = "String"

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
	ASSIGN        = "="
	PLUS_ASSIGN   = "+="
	MINUS_ASSIGN  = "-="
	TIMES_ASSIGN  = "*="
	DIVIDE_ASSIGN = "/="
	MODULO_ASSIGN = "%="
	AND_ASSIGN    = "&="
	OR_ASSIGN     = "|="
	NOT_ASSIGN    = "~="
	XOR_ASSIGN    = "^="
	// todo bitwise shift eq?
	INCREMENT = "++"
	DECREMENT = "--"
)

var Keywords = []string{
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

var Primitives = []string{
	"null",
	"true",
	"false",
}

var Prefix = []Type{
	PLUS,
	MINUS,
	INCREMENT,
	NOT,
}

var Postfix = []Type{
	DECREMENT,
}

var SimpleMath = []Type{
	PLUS,
	MINUS,
}

var ComplexMath = []Type{
	TIMES,
	DIVIDE,
	MODULO,
}

var SimpleComparison = []Type{
	AND,
	OR,
}

var ComplexComparison = []Type{
	EQUALS,
	NOT_EQUALS,
	GREATER_EQUALS,
	LESS_EQUALS,
	GREATER,
	LESS,
}

var Assignment = []Type{
	ASSIGN,
	PLUS_ASSIGN,
	MINUS_ASSIGN,
	TIMES_ASSIGN,
	DIVIDE_ASSIGN,
	MODULO_ASSIGN,
	AND_ASSIGN,
	OR_ASSIGN,
	NOT_ASSIGN,
	XOR_ASSIGN,
}
