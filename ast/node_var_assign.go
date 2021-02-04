package ast

import (
	"bytes"
	"sk-go/token"
)

type VarAssignNode struct {
	Token    *token.Token
	Name     *VarAccessNode
	Operator string
	Value    Node
}

func (as *VarAssignNode) String() string {
	var out bytes.Buffer
	out.WriteString(as.Name.String())
	out.WriteString(as.Operator)
	out.WriteString(as.Value.String())
	return out.String()
}
