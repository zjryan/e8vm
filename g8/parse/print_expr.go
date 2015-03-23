package parse

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/g8/ast"
)

type printer struct {
	buf *bytes.Buffer
	e   error
}

func newPrinter() *printer {
	ret := new(printer)
	ret.buf = new(bytes.Buffer)
	return ret
}

func (p *printer) printStr(s string) {
	if p.e != nil {
		return
	}
	_, e := fmt.Fprint(p.buf, s)
	p.e = e
}

func (p *printer) Print(args ...interface{}) {
	for _, arg := range args {
		p.printExpr(arg)
	}
}

func (p *printer) Error() error { return p.e }

func (p *printer) String() string { return p.buf.String() }

func (p *printer) printExpr(expr ast.Expr) {
	switch expr := expr.(type) {
	case string:
		p.printStr(expr)
	case *ast.Operand:
		p.Print(expr.Token.Lit)
	case *ast.OpExpr:
		if expr.A == nil {
			p.Print("(", expr.Op.Lit, expr.B, ")")
		} else {
			p.Print("(", expr.A, expr.Op.Lit, expr.B, ")")
		}
	case *ast.ParenExpr:
		p.Print("(", expr.Expr, ")")
	case *ast.ExprList:
		for i, e := range expr.Exprs {
			if i != 0 {
				p.Print(",")
			}
			p.Print(e)
		}
	case *ast.CallExpr:
		p.Print(expr.Func, "(", expr.Args, ")")
	default:
		panic("unknown expression type")
	}
}

// PrintExpr prints an expression
func PrintExpr(expr ast.Expr) string {
	p := newPrinter()
	p.printExpr(expr)
	return p.String()
}
