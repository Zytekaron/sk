package ast

import "sk-go/token"

type Node interface {
	// Token returns this node's token
	GetToken() *token.Token

	// String returns a string representation of this value
	String() string

	// Format formats the object using a pretty printing format
	//Format(depth, offset int, visited []Node) string
}
