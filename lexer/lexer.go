package lexer

import (
	"sk-go/token"
	"strings"
)

const EOF = rune(0)

const (
	NUMBERS = "0123456789"
	IDENTIFIERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_$"
)

var (
	KEYWORDS = []string{
		"int", "long", "float", "double", "string", "bool",
		"fn", "var", "const",
		"if", "else", "for", "of", "in", "while", "switch", "case", "return", "break", "continue",
		"mew", "delete",
	}
	BOOLEANS = []string{"true", "false"}
)

// Lexer holds our object-state.
type Lexer struct {
	// index is where the lexer is at in the input
	index int

	// ch is the character being read
	ch rune

	// characters is the input string as []rune
	characters []rune
}

// Create a new Lexer to lex a given string input
func New(src string) *Lexer {
	return &Lexer{
		index: -1,
		characters: []rune(src),
	}
}

func (l *Lexer) advance() {
	l.index++
	if l.index >= len(l.characters) {
		l.index--
		l.ch = EOF
	} else {
		l.ch = l.characters[l.index]
	}
}

// todo: check if I need to advance before all these
//  returns (if so, do that in the if
//  statement with some fancy function stuff)
func (l *Lexer) Next() *token.Token {
	l.advance()

	//fmt.Printf("reading '%c' (i=%d)\n", l.ch, l.index)

	switch l.ch {
	case ' ', '\r', '\n', '\t':
		return l.Next()
	case '/':
		switch l.ch {
		case '/':
			l.readComment()
			return l.Next()
		case '*':
			l.readBlockComment()
		case '=':
			return token.NewType(token.DIVIDE_EQ)
		default:
			return token.NewType(token.DIVIDE)
		}
	case '*':
		switch l.ch {
		case '*':
			return token.NewType(token.POW)
		case '=':
			return token.NewType(token.TIMES_EQ)
		default:
			return token.NewType(token.TIMES)
		}
	case '+':
		switch l.ch {
		case '+':
			l.advance()
			return token.NewType(token.PLUS_PLUS)
		case '=':
			l.advance()
			return token.NewType(token.PLUS_EQ)
		default:
			return token.NewType(token.PLUS)
		}
	case '-':
		switch l.ch {
		case '-':
			l.advance()
			return token.NewType(token.MINUS_MINUS)
		case '=':
			l.advance()
			return token.NewType(token.MINUS_EQ)
		default:
			return token.NewType(token.MINUS)
		}
	case '=':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.EQUALS)
		default:
			return token.NewType(token.EQ)
		}
	case '&':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.AND_EQ)
		case '&':
			l.advance()
			return token.NewType(token.AND)
		default:
			return token.NewType(token.BITWISE_AND)
		}
	case '|':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.OR_EQ)
		case '|':
			l.advance()
			return token.NewType(token.OR)
		default:
			return token.NewType(token.BITWISE_OR)
		}
	case '^':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.XOR_EQ)
		default:
			return token.NewType(token.BITWISE_XOR)
		}
	case '~':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.NOT_EQ)
		default:
			return token.NewType(token.BITWISE_NOT)
		}
	case '.':
		// if the second char is not a dot
		if l.ch != '.' {
			return token.NewType(token.DOT)
		}
		l.advance()
		// if the third char is not a dot
		if l.ch != '.' {
			return token.NewType(token.TWO_DOT)
		}
		l.advance()
		return token.NewType(token.THREE_DOT)
	case '!':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.NOT_EQUALS)
		default:
			return token.NewType(token.NOT)
		}
	case '%':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.MODULO_EQ)
		default:
			return token.NewType(token.MODULO)
		}
	case '<':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.LESS_EQUALS)
		default:
			return token.NewType(token.LESS)
		}
	case '>':
		switch l.ch {
		case '=':
			l.advance()
			return token.NewType(token.GREATER_EQUALS)
		default:
			return token.NewType(token.GREATER)
		}
	case '(':
		return token.NewType(token.LPAREN)
	case ')':
		return token.NewType(token.RPAREN)
	case '{':
		return token.NewType(token.LBRACE)
	case '}':
		return token.NewType(token.RBRACE)
	case '[':
		return token.NewType(token.LBRACKET)
	case ']':
		return token.NewType(token.RBRACKET)
	case ',':
		return token.NewType(token.COMMA)
	case '?':
		return token.NewType(token.QUESTION)
	case ':':
		return token.NewType(token.COLON)
	case ';':
		return token.NewType(token.SEMICOLON)
	case '"', '\'', '`':
		return l.readString()
	default:
		if strings.ContainsRune(NUMBERS, l.ch) {
			return l.readNumber()
		}
		if strings.ContainsRune(IDENTIFIERS, l.ch) {
			return l.readIdentifier()
		}
	}

	return nil
}

func (l *Lexer) readComment() {
	for l.ch != '\n' && l.ch != EOF {
		l.advance()
	}
}

func (l *Lexer) readNumber() *token.Token {
	var buf strings.Builder
	decimal := false

	for l.ch != EOF && strings.ContainsRune(NUMBERS + ".", l.ch) {
		if l.ch == '.' {
			if decimal {
				break
			}
			decimal = true
			buf.WriteRune('.')
		} else {
			buf.WriteRune(l.ch)
		}
		l.advance()
	}

	if l.ch == 'L' {
		l.advance()
		return token.New(token.LONG, buf.String())
	}

	if l.ch == 'f' {
		l.advance()
		return token.New(token.FLOAT, buf.String())
	}

	if l.ch == 'd' {
		l.advance() // default is double
	}

	if decimal {
		return token.New(token.DOUBLE, buf.String())
	}

	return token.New(token.INT, buf.String())
}

func (l *Lexer) readString() *token.Token {
	var buf strings.Builder
	quote := l.ch
	l.advance() // left quote

	for l.ch != EOF {
		if l.ch == '\\' {
			l.advance()
			if l.ch == EOF {
				// unexpected eof
				return nil
			} else if l.ch == 'n' {
				buf.WriteRune('\n')
			} else if l.ch == 't' {
				buf.WriteRune('\t')
			} else if l.ch == '"' {
				buf.WriteRune('"')
			} else if l.ch == '`' {
				buf.WriteRune('`')
			} else if l.ch == '\'' {
				buf.WriteRune('\'')
			} else if l.ch == 'u' {
				// todo implement characters (new String(Character.toChars()))
				for i := 0; i < 4; i++ {
					l.advance()
				}
				buf.WriteString("<?>")
			} else {
				// invalid escape sequence
				return nil
			}
		} else if l.ch == quote {
			break
		} else {
			buf.WriteRune(l.ch)
		}
		l.advance()
	}

	l.advance() // right quote
	return token.New(token.STRING, buf.String())
}

// fixme todo check this thoroughly!

func (l *Lexer) readBlockComment() {
	found := false

	var last rune
	for !found {
		last = l.ch
		l.advance()

		if last == EOF {
			found = true
			break
		}

		if last == '*' && l.ch == '/' {
			found = true
		}
	}

	l.readSpaces()
}

func (l *Lexer) readIdentifier() *token.Token {
    var buf strings.Builder

	for l.ch != EOF && strings.ContainsRune(IDENTIFIERS+NUMBERS, l.ch) {
		buf.WriteRune(l.ch)
		l.advance()
	}

	id := buf.String()
	if contains(KEYWORDS, id) {
		return token.New(token.KEYWORD, id)
	} else if contains(BOOLEANS, id) {
		return token.New(token.BOOL, id)
	} else {
		return token.New(token.IDENTIFIER, id)
	}
}

func (l *Lexer) readSpaces() {
	for l.ch == ' ' && l.ch != EOF {
		l.advance()
	}
}

func contains(str []string, r string) bool {
	for _, c := range str {
		if c == r {
			return true
		}
	}
	return false
}