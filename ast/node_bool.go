package ast

import (
	"sk-go/token"
)

type BoolNode struct {
	Token *token.Token
	Value bool
}

func (b *BoolNode) String() string {
	return b.Token.Value
}

func (b *BoolNode) GetToken() *token.Token {
	return b.Token
}
