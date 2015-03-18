package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type instParse func(lex8.Logger, []*lex8.Token) (*ast.Inst, bool)
type instParsers []instParse

func (ips instParsers) parse(log lex8.Logger, ops []*lex8.Token) *ast.Inst {
	for _, ip := range ips {
		if i, hit := ip(log, ops); hit {
			return i
		}
	}

	op0 := ops[0]
	log.Errorf(op0.Pos, "invalid asm instruction %q", op0.Lit)
	return nil
}
