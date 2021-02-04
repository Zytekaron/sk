package ast

import (
	"bytes"
	"sk-go/token"
)

type ConstStatement struct {
	Token *token.Token
	Name  *VarAccessNode
	Value Node
}

func (c *ConstStatement) String() string {
	var out bytes.Buffer
	out.WriteString(c.Token.Value)
	out.WriteString(" ")
	out.WriteString(c.Name.Token.Value)
	out.WriteString(" = ")
	if c.Value != nil {
		out.WriteString(c.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
