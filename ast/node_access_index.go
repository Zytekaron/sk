package ast

import (
	"sk-go/token"
)

type AccessNode struct {
	Token *token.Token
	Value Node
}

func (a *AccessNode) String() string {
	panic("unimplemented")
}

func (a *AccessNode) GetToken() *token.Token {
	return a.Token
}
