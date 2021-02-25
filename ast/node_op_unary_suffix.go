package ast

import (
	"bytes"
	"sk-go/token"
)

type SuffixOperationNode struct {
	Token    *token.Token
	Operator string
	Operand  Node
}

func (s *SuffixOperationNode) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(s.Operator)
	out.WriteString(s.Operator)
	out.WriteString(")")
	return out.String()
}

func (s *SuffixOperationNode) GetToken() *token.Token {
	return s.Token
}
