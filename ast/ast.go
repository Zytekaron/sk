package ast

import (
	"bytes"
	"sk-go/token"
)

// Program represents a complete program.
type Program BlockNode

func (p *Program) TokenValue() *token.Token {
	if len(p.Statements) > 0 {
		return p.Statements[0].GetToken()
	}
	return nil
}

// String returns this object as a string.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
