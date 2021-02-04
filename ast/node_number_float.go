package ast

import "sk-go/token"

type FloatNode struct {
	Token *token.Token
	Value float32
}

func (f *FloatNode) String() string {
	return f.Token.Value
}
