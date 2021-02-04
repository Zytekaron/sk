package ast

import (
	"bytes"
	"sk-go/token"
)

type WhileNode struct {
	Token     *token.Token
	Condition Node
	Body      *BlockNode
}

func (w *WhileNode) String() string {
	var out bytes.Buffer
	out.WriteString("while ")
	out.WriteString(w.Condition.String())
	out.WriteString(" {")
	out.WriteString(w.Body.String())
	out.WriteString("}")
	return out.String()
}
