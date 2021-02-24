package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

type ArrayNode struct {
	Token    *token.Token
	Elements []Node
}

func (a *ArrayNode) String() string {
	elements := make([]string, len(a.Elements))
	for _, el := range a.Elements {
		elements = append(elements, el.String())
	}
	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

func (a *ArrayNode) GetToken() *token.Token {
	return a.Token
}
