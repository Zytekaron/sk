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

func (v *VarAssignNode) String() string {
	var out bytes.Buffer
	out.WriteString(v.Name.String())
	out.WriteString(v.Operator)
	out.WriteString(v.Value.String())
	return out.String()
}

func (v *VarAssignNode) GetToken() *token.Token {
	return v.Name
}

func (v *VarAssignNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
