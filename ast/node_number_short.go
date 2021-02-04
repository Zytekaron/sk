package ast

import "sk-go/token"

type ShortNode struct {
	Token *token.Token
	Value int16
}

func (s *ShortNode) String() string {
	return s.Token.Value
}
