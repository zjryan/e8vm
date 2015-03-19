package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseOps(p *parser) (ops []*lex8.Token) {
	for !p.Accept(Semi) {
		t := p.Expect(Operand)
		if t != nil {
			ops = append(ops, t)
		} else {
			ops = nil // error now
			if p.See(lex8.EOF) {
				break
			}
			p.Next() // proceed anyway for other stuff
		}
	}

	p.BailOut()
	return ops
}

func parseFuncStmt(p *parser) *ast.FuncStmt {
	ops := parseOps(p)
	if len(ops) == 0 {
		return nil
	}

	return &ast.FuncStmt{Ops: ops}
}
