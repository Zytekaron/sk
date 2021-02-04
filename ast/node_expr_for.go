package ast

import (
	"bytes"
	"sk-go/token"
)

type ForNode struct {
	Token       *token.Token
	Declaration Node
	Condition   Node
	Increment   Node
	Then        *BlockNode
}

func (f *ForNode) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	out.WriteString(f.Declaration.String())
	out.WriteString("; ")
	out.WriteString(f.Condition.String())
	out.WriteString("; ")
	out.WriteString(f.Increment.String())
	out.WriteString(" {")
	out.WriteString(f.Then.String())
	out.WriteString("}")
	return out.String()
}
