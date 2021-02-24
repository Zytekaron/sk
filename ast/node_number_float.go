package ast

import "sk-go/token"

type FloatNode struct {
	Token *token.Token
	Value float32
}

func (f *FloatNode) String() string {
	return f.Token.Value
}

func (f *FloatNode) GetToken() *token.Token {
	return f.Token
}
