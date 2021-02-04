package ast

import (
	"bytes"
	"sk-go/token"
)

type LetStatement struct {
	Token *token.Token
	Name  *VarAccessNode
	Value Node
}

func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.Token.Value)
	out.WriteString(" ")
	out.WriteString(l.Name.Token.Value)
	out.WriteString(" = ")
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
