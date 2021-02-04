package ast

import (
	"bytes"
	"sk-go/token"
)

type ArrayAccessNode struct {
	Token *token.Token
	Array Node
	Index Node
}

func (a *ArrayAccessNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(a.Array.String())
	out.WriteString("[")
	out.WriteString(a.Index.String())
	out.WriteString("])")
	return out.String()
}
