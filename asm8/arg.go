package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func argCount(log lex8.Logger, ops []*lex8.Token, n int) bool {
	if len(ops) == n+1 {
		return true
	}

	log.Errorf(ops[0].Pos, "%q needs %d arguments", ops[0].Lit, n)
	return false
}
