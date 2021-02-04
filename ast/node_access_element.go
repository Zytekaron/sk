package ast

import (
	"bytes"
	"sk-go/token"
)

type ObjectAccessNode struct {
	Token  *token.Token
	Object Node
	Field  Node
}

func (o *ObjectAccessNode) String() string {
	var out bytes.Buffer
	out.WriteString(o.Object.String())
	out.WriteString("[")
	out.WriteString(o.Field.String())
	out.WriteString("]")
	return out.String()
}
