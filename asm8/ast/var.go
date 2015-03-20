package ast

import (
	"lonnie.io/e8vm/lex8"
)

// VarDecl is a variable declaration
type VarDecl struct {
	Stmts []*VarStmt

	Kw, Name             *lex8.Token
	Lbrace, Rbrace, Semi *lex8.Token
}

// VarStmt is a variable statement.
type VarStmt struct {
	Type *lex8.Token
	Args []*lex8.Token
}
