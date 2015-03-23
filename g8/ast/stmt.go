package ast

import (
	"lonnie.io/e8vm/lex8"
)

type Stmt interface{}

type ExprStmt struct {
	Expr
	Semi *lex8.Token
}

// Assign: exprList = exprList
type AssignStmt struct {
	Left   *ExprList
	Assign *lex8.Token
	Right  *ExprList
	Semi   *lex8.Token
}

// Define: exprList := exprList
type DefineStmt struct {
	Left   *ExprList
	Define *lex8.Token
	Right  *ExprList
	Semi   *lex8.Token
}

// Block is a statement block
type Block struct {
	Lbrace *lex8.Token
	Stmts  []Stmt
	Rbrace *lex8.Token
	Semi   *lex8.Token
}

// IfStmt is an if statement, possibly with an else of else if
// following
type IfStmt struct {
	If   *lex8.Token
	Expr Expr
	Body Stmt
	Else *ElseStmt // optional for else or else if
	Semi *lex8.Token
}

// ElseStmt
type ElseStmt struct {
	If   *lex8.Token // optional
	Expr Expr        // optional for else if
	Body Stmt
	Else *ElseStmt // optional for else if
}

// ForStmt
type ForStmt struct {
	Kw   *lex8.Token
	Init Stmt
	Cond Expr
	Iter Stmt
	Body Stmt
	Semi *lex8.Token
}

// ReturnStmt
type ReturnStmt struct {
	Kw    *lex8.Token
	Exprs *ExprList
	Semi  *lex8.Token
}

// ContinueStmt
type ContinueStmt struct {
	Kw, Label, Semi *lex8.Token
}

// BreakStmt
type BreakStmt struct {
	Kw, Label, Semi *lex8.Token
}

// EmptyStmt
type EmptyStmt struct {
	Semi *lex8.Token
}

// SwitchStmt
// switch expr {
//    case ..:
//    case ..:
// }
type SwitchStmt struct {
	Kw   *lex8.Token
	Semi *lex8.Token
}
