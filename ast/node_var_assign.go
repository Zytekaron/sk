package ast

import (
	"bytes"
	"sk-go/token"
)

type VarAssignNode struct {
	Name     *token.Token
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

func (v *VarAssignNode) Token() *token.Token {
	return v.Name
}

func (v *VarAssignNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
