package ast

import "sk-go/token"

type ByteNode struct {
	Token *token.Token
	Value int8
}

func (b *ByteNode) String() string {
	return b.Token.Value
}
