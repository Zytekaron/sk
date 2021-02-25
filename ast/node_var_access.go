package ast

import "sk-go/token"

type VarAccessNode struct {
	Token *token.Token
	Value string
}

func (v *VarAccessNode) String() string {
	return v.Value
}

func (v *VarAccessNode) GetToken() *token.Token {
	return v.Token
}
