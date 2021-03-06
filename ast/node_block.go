package ast

import (
	"bytes"
	"sk-go/token"
)

type BlockNode struct {
	Token      *token.Token
	Statements []Node
}

func (b *BlockNode) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (b *BlockNode) GetToken() *token.Token {
	return b.Token
}

func (b *BlockNode) Format(depth, offset int, visited []Node) string {
	panic("implement me") // fixme
}
