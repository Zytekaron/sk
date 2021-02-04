package ast

import "sk-go/token"

type DoubleNode struct {
	Token *token.Token
	Value int64
}

func (d *DoubleNode) String() string {
	return d.Token.Value
}
