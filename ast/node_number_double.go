package ast

import "sk-go/token"

type DoubleNode struct {
	Token *token.Token
	Value float64
}

func (d *DoubleNode) String() string {
	return d.Token.Value
}

func (d *DoubleNode) GetToken() *token.Token {
	return d.Token
}
