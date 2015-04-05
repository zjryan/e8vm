package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func parseParaList(p *parser) *ast.ParaList {
	panic("todo")
}

func parseFuncDecl(p *parser) *ast.FuncDecl {
	if !p.SeeKeyword("func") {
		panic("expect keyword")
	}

	ret := new(ast.FuncDecl)
	ret.Kw = p.Shift()
	ret.Name = p.Expect(Ident)
	ret.Lparen = p.ExpectOp("(")
	if p.InError() {
		return nil
	}

	if !p.SeeOp(")") {
		ret.Args = parseParaList(p)
		if p.InError() {
			return nil
		}
	}

	ret.Rparen = p.ExpectOp(")")
	if p.InError() {
		return nil
	}

	if p.SeeOp("(") {
		ret.RetLparen = p.Shift()
		ret.Rets = parseParaList(p)
		if p.InError() {
			return nil
		}
		ret.RetRparen = p.ExpectOp(")")
	} else if !p.SeeOp("{") {
		ret.RetType = parseType(p)
	}

	if p.InError() {
		return nil
	}

	ret.Body = parseBlock(p)
	ret.Semi = p.ExpectSemi()

	if p.InError() {
		return nil
	}

	return ret
}
