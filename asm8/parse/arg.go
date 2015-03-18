package parse

import (
	"lonnie.io/e8vm/lex8"
)

func argCount(p *parser, ops []*lex8.Token, n int) bool {
	if len(ops) == n+1 {
		return true
	}

	p.Errorf(ops[0].Pos, "%q needs %d arguments", ops[0].Lit, n)
	return false
}
