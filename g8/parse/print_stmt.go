package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func printStmt(p *printer, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.EmptyStmt:
		p.printStr("; // empty")
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
		return // skip the endl
	case *ast.AssignStmt:
		printExprs(p, stmt.Left, " = ", stmt.Right)
	case *ast.DefineStmt:
		printExprs(p, stmt.Left, " := ", stmt.Right)
	case *ast.ExprStmt:
		printExprs(p, stmt.Expr)
	case *ast.ReturnStmt:
		printExprs(p, "return ", stmt.Exprs)
	case *ast.ContinueStmt:
		if stmt.Label == nil {
			printExprs(p, "continue")
		} else {
			printExprs(p, "continue ", stmt.Label.Lit)
		}
	case *ast.BreakStmt:
		if stmt.Label == nil {
			printExprs(p, "break")
		} else {
			printExprs(p, "break ", stmt.Label.Lit)
		}
	case *ast.FallthroughStmt:
		p.printStr("fallthrough")
	default:
		panic("unknown statement type")
	}
	p.printEndl()
}

// PrintStmts prints a list of statements
func PrintStmts(stmts []ast.Stmt) string {
	p := newPrinter()
	printStmt(p, stmts)
	return p.String()
}
