package ast

import "sk-go/token"

type LongNode struct {
	Token *token.Token
	Value int64
}

func (l *LongNode) String() string {
	return l.Token.Value
}

func (l *LongNode) GetToken() *token.Token {
	return l.Token
}
