package ast

import (
	"bytes"
	"sk-go/token"
)

type ReturnNode struct {
	Token       *token.Token
	ReturnValue Node
}

func (r *ReturnNode) String() string {
	var out bytes.Buffer
	out.WriteString(r.Token.Value)
	out.WriteString(" ")
	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.GetToken().Value)
	}
	out.WriteString(";")
	return out.String()
}
