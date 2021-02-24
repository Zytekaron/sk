package ast

import (
	"math/big"
	"sk-go/token"
)

type BigIntNode struct {
	Token *token.Token
	Value big.Int
}

func (b *BigIntNode) String() string {
	return b.Token.Value
}

func (b *BigIntNode) GetToken() *token.Token {
	return b.Token
}
