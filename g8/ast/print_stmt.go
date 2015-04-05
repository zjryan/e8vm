package ast

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/fmt8"
)

func printStmt(p *fmt8.Printer, stmt Stmt) {
	switch stmt := stmt.(type) {
	case *EmptyStmt:
		fmt.Fprint(p, "; // emtpy")
	case *Block:
		if len(stmt.Stmts) > 0 {
			fmt.Fprintln(p, "{")
			p.Tab()
			printStmt(p, stmt.Stmts)
			p.ShiftTab()
			fmt.Fprint(p, "}")
		} else {
			fmt.Fprint(p, "{ }")
		}
	case *BlockStmt:
		printStmt(p, stmt.Block)
	case []Stmt:
		for _, s := range stmt {
			printStmt(p, s)
			fmt.Fprintln(p)
		}
	case *IfStmt:
		printExprs(p, "if ", stmt.Expr, " ")
		printStmt(p, stmt.Body)
		if stmt.Else != nil {
			printStmt(p, stmt.Else)
		}
	case *ElseStmt:
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
	case *ForStmt:
		printExprs(p, "for ", stmt.Cond, " ")
		printStmt(p, stmt.Body)
	case *AssignStmt:
		printExprs(p, stmt.Left, " = ", stmt.Right)
	case *DefineStmt:
		printExprs(p, stmt.Left, " := ", stmt.Right)
	case *ExprStmt:
		printExprs(p, stmt.Expr)
	case *ReturnStmt:
		printExprs(p, "return ", stmt.Exprs)
	case *ContinueStmt:
		if stmt.Label == nil {
			printExprs(p, "continue")
		} else {
			printExprs(p, "continue ", stmt.Label.Lit)
		}
	case *BreakStmt:
		if stmt.Label == nil {
			printExprs(p, "break")
		} else {
			printExprs(p, "break ", stmt.Label.Lit)
		}
	case *FallthroughStmt:
		fmt.Fprint(p, "fallthrough")
	case *VarDecls:
		printVarDecls(p, stmt)
	case *ConstDecls:
		printConstDecls(p, stmt)
	default:
		fmt.Fprintf(p, "<!!%T>", stmt)
	}
}

// PrintStmts prints a list of statements
func PrintStmts(stmts []Stmt) string {
	buf := new(bytes.Buffer)
	p := fmt8.NewPrinter(buf)
	printStmt(p, stmts)
	return buf.String()
}
