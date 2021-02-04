package ast

import (
	"sk-go/token"
	"sk-go/types"
)

type FunctionParamNode struct {
	Name    *token.Token
	Type    *types.Type
	Default Node
	Spread  bool
}

func (f *FunctionParamNode) String() string {
	return f.Name.Value // fixme ??
}

func (f *FunctionParamNode) Token() *token.Token {
	return f.Name
}

func (f *FunctionParamNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
