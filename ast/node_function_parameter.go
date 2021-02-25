package ast

import (
	"fmt"
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
	return fmt.Sprintf("%s = %v", f.Name, f.Default)
}

func (f *FunctionParamNode) GetToken() *token.Token {
	return f.Name
}

func (f *FunctionParamNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
