package ast

import (
	"bytes"
	"sk-go/token"
)

type SuffixOperationNode struct {
	Token    *token.Token
	Operator string
}

func (s *SuffixOperationNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(s.Token.Value)
	out.WriteString(s.Operator)
	out.WriteString(")")
	return out.String()
}
