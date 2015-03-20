package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseArgs(p *parser) (typ *lex8.Token, args []*lex8.Token) {
	typ = p.Expect(Operand)
	if typ == nil {
		p.skipErrStmt()
		return nil, nil
	}

	for !p.Accept(Semi) {
		if !p.InError() {
			t := p.Token()
			if t.Type == Operand || t.Type == String {
				args = append(args, t)
			} else {
				p.Errorf(t.Pos, "expect operand or string, got %s", p.TypeStr(t.Type))
			}
		}
		if p.See(lex8.EOF) {
			break
		}
		p.Next()
	}

	p.BailOut()

	return typ, args
}

func parseVarStmt(p *parser) *ast.VarStmt {
	typ, args := parseArgs(p)
	if typ == nil {
		return nil
	}

	ret := new(ast.VarStmt)
	ret.Type = typ
	ret.Args = args

	return ret
}
