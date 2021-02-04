package ast

import (
	"bytes"
	"sk-go/token"
)

type PrefixOperationNode struct {
	Token *token.Token
	Operator string
	Right Node
}

func (p *PrefixOperationNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}
