package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func printExprs(p *printer, exprs ...interface{}) {
	for _, expr := range exprs {
		printExpr(p, expr)
	}
}

func printExpr(p *printer, expr ast.Expr) {
	switch expr := expr.(type) {
	case string:
		p.printStr(expr)
	case *ast.Operand:
		printExprs(p, expr.Token.Lit)
	case *ast.OpExpr:
		if expr.A == nil {
			printExprs(p, "(", expr.Op.Lit, expr.B, ")")
		} else {
			printExprs(p, "(", expr.A, expr.Op.Lit, expr.B, ")")
		}
	case *ast.ParenExpr:
		printExprs(p, "(", expr.Expr, ")")
	case *ast.ExprList:
		printExprs(p, "[")
		for i, e := range expr.Exprs {
			if i != 0 {
				printExprs(p, ",")
			}
			printExprs(p, e)
		}
		printExprs(p, "]")
	case *ast.CallExpr:
		if expr.Args != nil {
			printExprs(p, expr.Func, "(", expr.Args, ")")
		} else {
			printExprs(p, expr.Func, "()")
		}
	default:
		panic("unknown expression type")
	}
}

// PrintExpr prints an expression
func PrintExpr(expr ast.Expr) string {
	p := newPrinter()
	printExpr(p, expr)
	return p.String()
}
