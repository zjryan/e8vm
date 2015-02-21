package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Func is an assembly function.
type FuncDecl struct {
	Stmts []*FuncStmt

	Kw, Name             *lex8.Token
	Lbrace, Rbrace, Semi *lex8.Token

	Index uint32 // symbol decl index
}
