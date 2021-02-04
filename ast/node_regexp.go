package ast

import (
	"fmt"
	"sk-go/token"
)

type RegexpNode struct {
	Token *token.Token
	Value string
	Flags string
}

func (r *RegexpNode) String() string {
	return fmt.Sprintf("/%s/%s", r.Value, r.Flags)
}
