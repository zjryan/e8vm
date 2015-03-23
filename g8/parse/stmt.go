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

func parseBlockClosed(p *parser) *ast.Block {
	ret := new(ast.Block)
	ret.Lbrace = p.ExpectOp("{")
	for !p.SeeOp("}") {
		if p.See(lex8.EOF) {
			break
		}
		if stmt := parseStmt(p); stmt != nil {
			ret.Stmts = append(ret.Stmts, stmt)
		}
		p.skipErrStmt()
	}
	ret.Rbrace = p.ExpectOp("}")
	ret.Semi = p.ExpectSemi()
	return ret
}

func parseIfStmt(p *parser) *ast.IfStmt {
	panic("todo")
}

func parseForStmt(p *parser) *ast.ForStmt {
	panic("todo")
}

func parseSwitchStmt(p *parser) *ast.SwitchStmt {
	panic("todo")
}

func parseReturnStmt(p *parser) *ast.ReturnStmt {
	panic("todo")
}

func parseBreakStmt(p *parser) *ast.BreakStmt {
	panic("todo")
}

func parseContinueStmt(p *parser) *ast.ContinueStmt {
	panic("todo")
}

func parseStmt(p *parser) ast.Stmt {
	first := p.Token()
	if first.Type == Keyword {
		switch first.Lit {
		case "const":
			return parseConstDecls(p)
		case "var":
			return parseVarDecls(p)
		case "if":
			return parseIfStmt(p)
		case "for":
			return parseForStmt(p)
		case "switch":
			return parseSwitchStmt(p)
		case "return":
			return parseReturnStmt(p)
		case "break":
			return parseBreakStmt(p)
		case "continue":
			return parseContinueStmt(p)
		}
	}

	if p.SeeOp("{") {
		return parseBlockClosed(p)
	}

	exprs := parseExprList(p)
	if p.SeeOp("=") {
		// assigns statement
		ret := new(ast.AssignStmt)
		ret.Left = exprs
		ret.Assign = p.Shift()
		ret.Right = parseExprList(p)
		ret.Semi = p.ExpectSemi()
		return ret
	} else if p.SeeOp(":=") {
		// define statement
		ret := new(ast.DefineStmt)
		ret.Left = exprs
		ret.Define = p.Shift()
		ret.Right = parseExprList(p)
		ret.Semi = p.ExpectSemi()
		return ret
	} else if p.See(Semi) {
		// empty statement
		ret := new(ast.EmptyStmt)
		ret.Semi = p.Shift()
		return ret
	}

	p.ErrorfHere("invalid statement")
	p.skipErrStmt()
	return nil
}
