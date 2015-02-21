package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseOps(p *parser) (ops []*lex8.Token) {
	for !p.acceptType(Semi) {
		t := p.expect(Operand)
		if t != nil {
			ops = append(ops, t)
		} else {
			ops = nil // error now
			if p.see(lex8.EOF) {
				break
			}
			p.next() // proceed anyway for other stuff
		}
	}

	p.clearErr()
	return ops
}

func parseFuncStmt(p *parser) *ast.FuncStmt {
	ops := parseOps(p)
	if len(ops) == 0 {
		return nil
	}

	op0 := ops[0]
	lead := op0.Lit
	if lead == "" {
		panic("empty operand")
	}

	if parseLabel(p, op0) {
		if len(ops) > 1 {
			p.err(op0.Pos, "label should take the entire line")
			return nil
		}
		return &ast.FuncStmt{Label: lead, Ops: ops}
	}

	return &ast.FuncStmt{Inst: parseInst(p, ops), Ops: ops}
}
