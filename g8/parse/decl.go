package parse

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseIdentList(p *parser) *ast.IdentList {
	ret := new(ast.IdentList)
	for p.See(Ident) {
		ret.Idents = append(ret.Idents, p.Shift())
		if !p.SeeOp(",") {
			break
		}
		ret.Commas = append(ret.Commas, p.Shift())
	}
	return ret
}

func parseConstDecls(p *parser) *ast.ConstDecls {
	if !p.SeeKeyword("const") {
		panic("expect keyword")
	}

	p.ErrorfHere("const declare not implemented")
	p.Next()
	return nil
}

func parseVarDecl(p *parser) *ast.VarDecl {
	ret := new(ast.VarDecl)
	ret.Idents = parseIdentList(p)
	if !p.See(Semi) && !p.SeeOp("=") && !p.SeeOp(")") {
		ret.Type = parseType(p) // it has a type
	}

	if p.SeeOp("=") {
		ret.Eq = p.Shift()
		ret.Exprs = parseExprList(p)
	} else if ret.Type == nil {
		p.ErrorfHere("expect type")
	}

	if p.InError() {
		return nil
	}
	ret.Semi = p.ExpectSemi()

	return ret
}

func parseVarDecls(p *parser) *ast.VarDecls {
	if !p.SeeKeyword("var") {
		panic("expect keyword")
	}

	ret := new(ast.VarDecls)
	ret.Kw = p.Shift()

	if p.SeeOp("(") {
		ret.Lparen = p.Shift()
		for !p.See(lex8.EOF) {
			d := parseVarDecl(p)
			if d != nil {
				ret.Decls = append(ret.Decls, d)
			} else {
				p.skipErrStmt()
			}

			if p.SeeOp(")") {
				break
			}
		}
		ret.Rparen = p.ExpectOp("}")
		ret.Semi = p.ExpectSemi()

		return ret
	}

	d := parseVarDecl(p)
	ret.Decls = []*ast.VarDecl{d}
	return ret // no semi means it is not a group
}

func parseStruct(p *parser) *ast.Struct {
	if !p.SeeKeyword("struct") {
		panic("expect keyword")
	}

	p.ErrorfHere("struct declare not implemented")
	p.Next()
	return nil
}
