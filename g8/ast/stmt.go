package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Stmt is a general statement
type Stmt interface{}

// ExprStmt is a statement with just an expression
type ExprStmt struct {
	Expr
	Semi *lex8.Token
}

// AssignStmt is an assignment statement:
// exprList = exprList
type AssignStmt struct {
	Left   *ExprList
	Assign *lex8.Token
	Right  *ExprList
	Semi   *lex8.Token
}

// DefineStmt is a statement that defines one or a list of variables.
// exprList := exprList
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
}

// BlockStmt is a block statement
type BlockStmt struct {
	*Block
	Semi *lex8.Token
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

// ElseStmt is the dangling statement block after if
type ElseStmt struct {
	Else *lex8.Token
	If   *lex8.Token // optional
	Expr Expr        // optional for else if
	Body *Block
	Next *ElseStmt // next else statment
}

// ForStmt is a loop statement
type ForStmt struct {
	Kw   *lex8.Token
	Init Stmt
	Cond Expr
	Iter Stmt
	Body *Block
	Semi *lex8.Token
}

// ReturnStmt is a statement of return.
// return <expr>
type ReturnStmt struct {
	Kw    *lex8.Token
	Exprs *ExprList
	Semi  *lex8.Token
}

// ContinueStmt is the continue statement
// continue [<label>]
type ContinueStmt struct{ Kw, Label, Semi *lex8.Token }

// BreakStmt is the break statement
// break [<label>]
type BreakStmt struct{ Kw, Label, Semi *lex8.Token }

// FallthroughStmt is the fallthrough statement
// fallthrough
type FallthroughStmt struct{ Kw, Semi *lex8.Token }

// EmptyStmt is an empty statement created by
// an orphan semicolon
type EmptyStmt struct {
	Semi *lex8.Token
}

// SwitchStmt is a case switching statement
// switch expr {
//    case ..:
//    case ..:
// }
// TODO:
type SwitchStmt struct {
	Kw   *lex8.Token
	Semi *lex8.Token
}
