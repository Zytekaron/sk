package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

type FunctionCallNode struct {
	Token     *token.Token
	Function  *token.Token
	Arguments []Node
}

func (f *FunctionCallNode) String() string {
	var out bytes.Buffer
	args := make([]string, len(f.Arguments))
	for _, a := range f.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(f.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (f *FunctionCallNode) GetToken() *token.Token {
	return f.Token
}
