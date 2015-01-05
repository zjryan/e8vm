package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func checkAllType(p *parser, args []*lex8.Token, t int) bool {
	for _, arg := range args {
		if arg.Type != t {
			p.err(arg.Pos, "expect %s, got %s", typeStr(t), typeStr(arg.Type))
			return false
		}
	}
	return true
}
