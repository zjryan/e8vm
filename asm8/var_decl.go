package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type varDecl struct {
	stmts []*varStmt

	kw, name             *lex8.Token
	lbrace, rbrace, semi *lex8.Token

	index uint32 // symbol decl index
}
