package asm8

import (
	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

func checkTypeAll(p lex8.Logger, toks []*lex8.Token, typ int) bool {
	for _, t := range toks {
		if t.Type != typ {
			p.Errorf(t.Pos, "expect operand, got %s", parse.TypeStr(t.Type))
			return false
		}
	}
	return true
}
