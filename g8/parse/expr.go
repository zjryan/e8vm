package parse

import (
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseOperand(p *parser) ast.Expr {
	if p.See(Ident) || p.See(Int) || p.See(Float) ||
		p.See(String) || p.See(Char) {
		return &ast.Operand{p.Shift()}
	} else if p.SeeOp("(") {
		lp := p.Shift()
		expr := parseExpr(p)
		rp := p.ExpectOp(")")
		if rp == nil {
			return nil
		}

		return &ast.ParenExpr{
			Lparen: lp,
			Rparen: rp,
			Expr:   expr,
		}
	}

	t := p.Token()
	p.Errorf(t.Pos, "expect an operand, got %s", p.typeStr(t))
	return nil
}

func parseExprList(p *parser) *ast.ExprList {
	ret := new(ast.ExprList)
	for {
		expr := parseExpr(p)
		if p.InError() {
			return nil
		}
		ret.Exprs = append(ret.Exprs, expr)
		if !p.SeeOp(",") {
			break
		}
		ret.Commas = append(ret.Commas, p.Shift())
	}
	return ret
}

func parseExprListClosed(p *parser, closeWith string) *ast.ExprList {
	ret := new(ast.ExprList)

	for {
		expr := parseExpr(p)
		if p.InError() {
			return nil
		}
		ret.Exprs = append(ret.Exprs, expr)

		if p.SeeOp(closeWith) {
			return ret
		}

		comma := p.ExpectOp(",")
		if comma == nil {
			return nil
		}
		ret.Commas = append(ret.Commas, comma)

		// could be a trailing comma
		if p.SeeOp(closeWith) {
			return ret
		}
	}

	return ret
}

func parsePrimaryExpr(p *parser) ast.Expr {
	ret := parseOperand(p)

	for {
		if !p.SeeOp("(") {
			break
		}

		lp := p.Shift()

		if p.SeeOp(")") {
			rp := p.Shift()
			ret = &ast.CallExpr{
				Func:   ret,
				Args:   nil,
				Lparen: lp,
				Rparen: rp,
			}
			continue
		}

		lst := parseExprListClosed(p, ")")
		if p.InError() {
			return nil
		}
		rp := p.ExpectOp(")")
		if rp == nil {
			return nil
		}

		ret = &ast.CallExpr{
			Func:   ret,
			Args:   lst,
			Lparen: lp,
			Rparen: rp,
		}
	}

	return ret
}

func parseUnaryExpr(p *parser) ast.Expr {
	if p.SeeOp("+", "-", "!", "^") {
		t := p.Shift()
		expr := parseUnaryExpr(p)
		return &ast.OpExpr{A: nil, Op: t, B: expr}
	}

	return parsePrimaryExpr(p)
}

func opPrec(op string) int {
	switch op {
	case "|":
		return 0
	case "&":
		return 1
	case "==", "!=", "<", "<=", ">=", ">":
		return 2
	case "+", "-", "||", "^":
		return 3
	case "*", "%", "/", "<<", ">>", "&&":
		return 4
	}
	return -1
}

func parseBinaryExpr(p *parser, prec int) ast.Expr {
	ret := parseUnaryExpr(p)
	if p.InError() {
		return nil
	}

	if p.See(Operator) {
		startPrec := opPrec(p.Token().Lit)
		for i := startPrec; i >= prec; i-- {
			for p.See(Operator) {
				if opPrec(p.Token().Lit) != i {
					break
				}

				op := p.Shift()
				bop := new(ast.OpExpr)
				bop.A = ret
				bop.Op = op
				bop.B = parseBinaryExpr(p, i+1)
				ret = bop
			}
		}
	}

	return ret
}

func parseExpr(p *parser) ast.Expr {
	return parseBinaryExpr(p, 0)
}

// Exprs parses a list of expressions and returns an array of ast node of
// these expressions.
func Exprs(f string, r io.Reader) ([]ast.Expr, []*lex8.Error) {
	var ret []ast.Expr

	p, _ := newParser(f, r)
	for !p.See(lex8.EOF) {
		expr := parseExpr(p)
		if expr != nil {
			ret = append(ret, expr)
		}

		p.ExpectSemi()
		if p.InError() {
			p.skipErrStmt()
		}
	}

	if es := p.Errs(); es != nil {
		return nil, es
	}

	return ret, nil
}
