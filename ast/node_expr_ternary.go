package ast

import (
	"bytes"
	"sk-go/token"
)

type TernaryNode struct {
	Token     *token.Token
	Condition Node
	IfTrue    Node
	IfFalse   Node
}

func (t *TernaryNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(t.Condition.String())
	out.WriteString(" ? ")
	out.WriteString(t.IfTrue.String())
	out.WriteString(" : ")
	out.WriteString(t.IfFalse.String())
	out.WriteString(")")
	return out.String()
}
