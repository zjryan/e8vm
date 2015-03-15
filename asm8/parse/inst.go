package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type instParse func(*parser, []*lex8.Token) (*ast.Inst, bool)
type instParsers []instParse

func (ips instParsers) parse(p *parser, ops []*lex8.Token) *ast.Inst {
	for _, ip := range ips {
		if i, hit := ip(p, ops); hit {
			return i
		}
	}

	op0 := ops[0]
	p.err(op0.Pos, "invalid asm instruction %q", op0.Lit)
	return nil
}
