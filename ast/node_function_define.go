package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

type FunctionNode struct {
	Name       *token.Token
	Parameters []*FunctionParamNode
	Body       *BlockNode
}

func (f *FunctionNode) String() string {
	var out bytes.Buffer
	params := make([]string, len(f.Parameters))
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.Name.Value)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f *FunctionNode) Token() *token.Token {
	return f.Name
}

func (f *FunctionNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
