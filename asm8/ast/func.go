package ast

import (
	"lonnie.io/e8vm/lex8"
)

// FuncDecl is an assembly function.
type FuncDecl struct {
	Stmts []*FuncStmt

	Kw, Name             *lex8.Token
	Lbrace, Rbrace, Semi *lex8.Token
}

// FuncStmt is a statement in a assembly function.
// It is either a instruction or a label.
type FuncStmt struct {
	Ops []*lex8.Token
}
