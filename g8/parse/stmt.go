package parse

import (
	"io"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseBlock(p *parser) *ast.Block {
	ret := new(ast.Block)
	ret.Lbrace = p.ExpectOp("{")
	if ret.Lbrace == nil {
		return ret
	}

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
	return ret
}

func parseElse(p *parser) *ast.ElseStmt {
	if !p.SeeKeyword("else") {
		panic("must start with keyword")
	}

	ret := new(ast.ElseStmt)
	ret.Else = p.Shift()

	if p.SeeKeyword("if") {
		ret.If = p.Shift()
		ret.Expr = parseExpr(p)
	}

	if p.InError() {
		return ret
	}

	if !p.SeeOp("{") {
		p.ErrorfHere("missing else body")
		return ret
	}

	ret.Body = parseBlock(p)
	// might have another else
	if ret.If != nil && p.SeeKeyword("else") {
		ret.Next = parseElse(p)
	}

	return ret
}

func parseIfBody(p *parser) (ret ast.Stmt, isBlock bool) {
	if p.SeeOp("{") {
		return parseBlock(p), true
	}
	if p.SeeKeyword("return") {
		return parseReturnStmt(p, false), false
	} else if p.SeeKeyword("break") {
		return parseBreakStmt(p, false), false
	} else if p.SeeKeyword("continue") {
		return parseContinueStmt(p, false), false
	}

	p.ErrorfHere("expect if body")
	return nil, false
}

// if <cond> { <stmts> }
// if <cond> return <expr>
// if <cond> break
// if <cond> continue
// if <cond> { <stmts> } else { <stmts> }
// if <cond> { <stmts> } else if { <stmts> }
// if <cond> { <stmts> } else if { <stmts> } else { <stmts> }
func parseIfStmt(p *parser) *ast.IfStmt {
	if !p.SeeKeyword("if") {
		panic("must start with keyword")
	}

	ret := new(ast.IfStmt)
	ret.If = p.Shift()
	ret.Expr = parseExpr(p)
	if p.InError() {
		return ret
	}

	var isBlock bool
	ret.Body, isBlock = parseIfBody(p)
	if p.InError() {
		return ret
	}

	if isBlock && p.SeeKeyword("else") {
		// else clause only happens when the body is block
		ret.Else = parseElse(p)
		if p.InError() {
			return ret
		}
	}
	ret.Semi = p.ExpectSemi()
	return ret
}

// for <cond> { <stmts> }
func parseForStmt(p *parser) *ast.ForStmt {
	if !p.SeeKeyword("for") {
		panic("must start with keyword")
	}

	ret := new(ast.ForStmt)
	ret.Kw = p.Shift()
	ret.Cond = parseExpr(p)
	if p.InError() {
		return ret
	}

	ret.Body = parseBlock(p)
	if p.InError() {
		return ret
	}

	ret.Semi = p.ExpectSemi()
	return ret
}

func parseReturnStmt(p *parser, withSemi bool) *ast.ReturnStmt {
	ret := new(ast.ReturnStmt)
	ret.Kw = p.ExpectKeyword("return")
	if !p.SeeSemi() {
		ret.Exprs = parseExprList(p)
	}
	if withSemi {
		ret.Semi = p.ExpectSemi()
	}
	return ret
}

func parseBreakStmt(p *parser, withSemi bool) *ast.BreakStmt {
	ret := new(ast.BreakStmt)
	ret.Kw = p.ExpectKeyword("break")
	if p.See(Ident) {
		ret.Label = p.Expect(Ident)
	}
	if withSemi {
		ret.Semi = p.ExpectSemi()
	}
	return ret
}

func parseContinueStmt(p *parser, withSemi bool) *ast.ContinueStmt {
	ret := new(ast.ContinueStmt)
	ret.Kw = p.ExpectKeyword("continue")
	if p.See(Ident) {
		ret.Label = p.Expect(Ident)
	}
	if withSemi {
		ret.Semi = p.ExpectSemi()
	}
	return ret
}

func parseFallthroughStmt(p *parser) *ast.FallthroughStmt {
	ret := new(ast.FallthroughStmt)
	ret.Kw = p.ExpectKeyword("fallthrough")
	ret.Semi = p.ExpectSemi()
	return ret
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
		//case "switch":
		//	return parseSwitchStmt(p)
		case "return":
			return parseReturnStmt(p, true)
		case "break":
			return parseBreakStmt(p, true)
		case "continue":
			return parseContinueStmt(p, true)
		case "fallthrough":
			return parseFallthroughStmt(p)
		}
	}

	if p.SeeOp("{") {
		ret := new(ast.BlockStmt)
		ret.Block = parseBlock(p)
		ret.Semi = p.ExpectSemi()
		return ret
	} else if p.See(Semi) {
		ret := new(ast.EmptyStmt)
		ret.Semi = p.Shift()
		return ret
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
	} else if semi := p.AcceptSemi(); semi != nil {
		if exprs.Len() != 1 {
			p.ErrorfHere("expression list is not a valid statement")
			p.BailOut() // reset this error
			return nil
		}

		ret := new(ast.ExprStmt)
		ret.Expr = exprs.Exprs[0]
		ret.Semi = semi
		return ret
	}

	p.ErrorfHere("invalid statement")
	p.Next() // always make some progress
	return nil
}

// Stmts parses a file input stream as a list of statements,
// like a bare function body.
func Stmts(f string, r io.Reader) ([]ast.Stmt, []*lex8.Error) {
	var ret []ast.Stmt

	p, _ := newParser(f, r)
	for !p.See(lex8.EOF) {
		stmt := parseStmt(p)
		if stmt != nil {
			ret = append(ret, stmt)
		}

		if p.InError() {
			p.skipErrStmt()
		}
	}

	if es := p.Errs(); es != nil {
		return nil, es
	}

	return ret, nil
}
