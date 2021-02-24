package ast

import (
	"math/big"
	"sk-go/token"
)

type BigFloatNode struct {
	Token *token.Token
	Value big.Float
}

func (b *BigFloatNode) String() string {
	return b.Token.Value
}

func (b *BigFloatNode) GetToken() *token.Token {
	return b.Token
}
