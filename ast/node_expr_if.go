package ast

import (
	"bytes"
	"sk-go/token"
)

type IfExpression struct {
	Token     *token.Token
	Condition Node
	Then      *BlockNode
	Else      *BlockNode
}

func (i *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Then.String())
	if i.Else != nil {
		out.WriteString("else")
		out.WriteString(i.Else.String())
	}
	return out.String()
}
