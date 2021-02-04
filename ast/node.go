package ast

import "sk-go/token"

type Node interface {
	// todo consider this as Tok() or GetToken()
	// Token returns this node's token
	Token() *token.Token

	// String returns a string representation of this value
	String() string

	// Format formats the value into a pretty string
	Format(depth, offset int, visited []Node) string
}
