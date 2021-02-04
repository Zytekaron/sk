package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

type CallNode struct {
	Token     *token.Token
	Function  Node
	Arguments []Node
}

func (c *CallNode) String() string {
	var out bytes.Buffer
	args := make([]string, len(c.Arguments))
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
