package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func printStmt(p *printer, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.EmptyStmt:
		p.printStr("; // empty")
		p.printEndl()
	case *ast.Block:
		p.printStr("{")
		p.printEndl()
		p.Tab()
		printStmt(p, stmt.Stmts)
		p.ShiftTab()
		p.printStr("}")
		p.printEndl()
	case []ast.Stmt:
		for _, s := range stmt {
			printStmt(p, s)
		}
	case *ast.AssignStmt:
		printExprs(p, stmt.Left, " = ", stmt.Right)
		p.printEndl()
	case *ast.DefineStmt:
		printExprs(p, stmt.Left, " := ", stmt.Right)
		p.printEndl()
	case *ast.ExprStmt:
		printExprs(p, stmt.Expr)
		p.printEndl()
	default:
		panic("unknown statement type")
	}
}

// PrintStmt prints a statement
func PrintStmts(stmts []ast.Stmt) string {
	p := newPrinter()
	printStmt(p, stmts)
	return p.String()
}
