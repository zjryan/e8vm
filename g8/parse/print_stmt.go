package parse

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ast"
)

func printStmt(p *fmt8.Printer, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.EmptyStmt:
		fmt.Fprint(p, "; // emtpy")
	case *ast.Block:
		fmt.Fprintln(p, "{")
		p.Tab()
		printStmt(p, stmt.Stmts)
		p.ShiftTab()
		fmt.Fprint(p, "}")
	case []ast.Stmt:
		for _, s := range stmt {
			printStmt(p, s)
			fmt.Fprintln(p)
		}
	case *ast.IfStmt:
		printExprs(p, "if ", stmt.Expr, " ")
		printStmt(p, stmt.Body)
		if stmt.Else != nil {
			printStmt(p, stmt.Else)
		}
	case *ast.ElseStmt:
		if stmt.If == nil {
			printExprs(p, " else ")
			printStmt(p, stmt.Body)
		} else {
			printExprs(p, " else if ", stmt.Expr, " ")
			printStmt(p, stmt.Body)
			if stmt.Next != nil {
				printStmt(p, stmt.Next)
			}
		}
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
		fmt.Fprint(p, "fallthrough")
	default:
		fmt.Fprintf(p, "<!!%T>", stmt)
	}
}

// PrintStmts prints a list of statements
func PrintStmts(stmts []ast.Stmt) string {
	buf := new(bytes.Buffer)
	p := fmt8.NewPrinter(buf)
	printStmt(p, stmts)
	return buf.String()
}
