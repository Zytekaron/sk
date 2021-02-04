package ast

import "sk-go/token"

type IntNode struct {
	Token *token.Token
	Value int32
}

func (i *IntNode) String() string {
	return i.Token.Value
}
