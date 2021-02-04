package ast

import (
	"bytes"
	"sk-go/token"
)

type BinaryOperationNode struct {
	Token    *token.Token
	Left     Node
	Operator string
	Right    Node
}

func (b *BinaryOperationNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(b.Left.String())
	out.WriteString(" ")
	out.WriteString(b.Operator)
	out.WriteString(" ")
	out.WriteString(b.Right.String())
	out.WriteString(")")
	return out.String()
}
