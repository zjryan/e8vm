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

		comma := p.Shift()
		ret.Commas = append(ret.Commas, comma)
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

		lst := parseExprList(p)
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

func parseBinaryExpr(p *parser, prec int) ast.Expr {
	ret := parseUnaryExpr(p)
	if p.InError() {
		return nil
	}

	if prec == 0 {
		for p.SeeOp("+", "-") {
			tok := p.Shift()
			bop := new(ast.OpExpr)
			bop.A = ret
			bop.Op = tok
			bop.B = parseBinaryExpr(p, prec+1)
			ret = bop
		}
	}

	return ret
}

func parseExpr(p *parser) ast.Expr {
	return parseBinaryExpr(p, 0)
}

// Expr parses a bare expression and returns the ast node.
func Expr(f string, rc io.ReadCloser) (ast.Expr, []*lex8.Error) {
	p, _ := newParser(f, rc)
	ret := parseExpr(p)
	if es := p.Errs(); es != nil {
		return nil, es
	}
	return ret, nil
}
