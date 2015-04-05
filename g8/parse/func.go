package parse

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parsePara(p *parser) *ast.Para {
	ret := new(ast.Para)
	if p.See(Ident) {
		ret.Ident = p.Shift()
		if !(p.SeeOp(",") || p.SeeOp(")")) {
			ret.Type = parseType(p)
		}
	} else {
		ret.Type = parseType(p)
	}
	return ret
}

func parseParaList(p *parser) *ast.ParaList {
	if !p.SeeOp("(") {
		panic("expect left paren")
	}

	ret := new(ast.ParaList)
	ret.Lparen = p.Shift()

	if p.SeeOp(")") {
		// empty parameter list
		ret.Rparen = p.Shift()
		return ret
	}

	for !p.See(lex8.EOF) {
		para := parsePara(p)
		if p.InError() {
			return nil
		}

		ret.Paras = append(ret.Paras, para)
		if p.SeeOp(",") {
			ret.Commas = append(ret.Commas, p.Shift())
		}

		if p.SeeOp(")") {
			break
		}
	}

	ret.Rparen = p.ExpectOp(")")
	return ret
}

func parseFunc(p *parser) *ast.Func {
	if !p.SeeKeyword("func") {
		panic("expect keyword")
	}

	ret := new(ast.Func)
	ret.Kw = p.Shift()
	ret.Name = p.Expect(Ident)
	ret.Args = parseParaList(p)
	if p.InError() {
		return nil
	}

	if p.SeeOp("(") {
		ret.Rets = parseParaList(p)
		if p.InError() {
			return nil
		}
		if len(ret.Rets.Paras) == 0 {
			p.Errorf(ret.Rets.Lparen.Pos, "return list cannot be empty")
			return nil
		}
	} else if !p.SeeOp("{") {
		ret.RetType = parseType(p)
	}

	if p.InError() {
		return nil
	}

	ret.Body = parseBlock(p)
	ret.Semi = p.ExpectSemi()
	return ret
}
