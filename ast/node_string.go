package ast

import "sk-go/token"

type StringNode struct {
	Token *token.Token
	Value string
}

func (s *StringNode) String() string {
	return s.Token.Value
}

func (s *StringNode) GetToken() *token.Token {
	return s.Token
}
