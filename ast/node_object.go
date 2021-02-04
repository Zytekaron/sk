package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

type ObjectNode struct {
	Token *token.Token
	Pairs map[Node]Node
}

func (o *ObjectNode) String() string {
	var out bytes.Buffer
	pairs := make([]string, len(o.Pairs))
	for key, value := range o.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
