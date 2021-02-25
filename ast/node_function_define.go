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
	for i, p := range f.Parameters {
		params[i] = p.String()
	}
	out.WriteString("fn ")
	out.WriteString(f.Name.Value)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") { ")
	out.WriteString(f.Body.String())
	out.WriteString(" }")
	return out.String()
}

func (f *FunctionNode) GetToken() *token.Token {
	return f.Name
}

func (f *FunctionNode) Format(depth, offset int, visited []Node) string {
	var out bytes.Buffer
	params := make([]string, len(f.Parameters))
	for i, p := range f.Parameters {
		params[i] = p.String()
	}
	out.WriteString("fn ")
	out.WriteString(f.Name.Value)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") { ")
	out.WriteString(f.Body.Format(depth + 4, offset, visited))
	out.WriteString(" }")
	return out.String()
}
