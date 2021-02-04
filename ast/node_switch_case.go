package ast

import (
	"bytes"
	"sk-go/token"
	"strings"
)

// CaseExpression handles the case within a switch statement
type CaseExpression struct {
	Token *token.Token
	Default bool
	Expr []Node
	Block *BlockNode
}

func (c *CaseExpression) String() string {
	var out bytes.Buffer
	if c.Default {
		out.WriteString("default ")
	} else {
		out.WriteString("case ")
		var tmp []string
		for _, exp := range c.Expr {
			tmp = append(tmp, exp.String())
		}
		out.WriteString(strings.Join(tmp, ","))
	}
	out.WriteString(c.Block.String())
	return out.String()
}
