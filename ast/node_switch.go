package ast

import (
	"bytes"
	"sk-go/token"
)

type SwitchExpression struct {
	Token   *token.Token
	Value   Node
	Choices []*CaseExpression
}

func (s *SwitchExpression) String() string {
	var out bytes.Buffer
	out.WriteString("\nswitch (")
	out.WriteString(s.Value.String())
	out.WriteString(")\n{\n")
	for _, tmp := range s.Choices {
		if tmp != nil {
			out.WriteString(tmp.String())
		}
	}
	out.WriteString("}\n")
	return out.String()
}
