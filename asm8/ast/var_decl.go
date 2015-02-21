package ast

import (
	"lonnie.io/e8vm/lex8"
)

// VarDecl is a variable declaration
type VarDecl struct {
	Stmts []*VarStmt

	Kw, Name             *lex8.Token
	Lbrace, Rbrace, Semi *lex8.Token

	Index uint32 // symbol decl index
}
