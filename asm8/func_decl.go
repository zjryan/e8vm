package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Func is an assembly function.
type funcDecl struct {
	stmts []*funcStmt

	kw, name             *lex8.Token
	lbrace, rbrace, semi *lex8.Token

	index uint32 // symbol decl index
}
