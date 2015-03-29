package parse

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ast"
)

func printExprs(p *fmt8.Printer, args ...interface{}) {
	for _, arg := range args {
		printExpr(p, arg)
	}
}

func printExpr(p *fmt8.Printer, expr ast.Expr) {
	switch expr := expr.(type) {
	case string:
		fmt.Fprintf(p, expr)
	case *ast.Operand:
		fmt.Fprintf(p, expr.Token.Lit)
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
	buf := new(bytes.Buffer)
	p := fmt8.NewPrinter(buf)
	printExpr(p, expr)
	return buf.String()
}
