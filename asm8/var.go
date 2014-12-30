package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type varStmt struct {
	typ  *lex8.Token
	toks []*lex8.Token
}

type Var struct {
	stmts []*varStmt

	kw, name             *lex8.Token
	lbrace, rbrace, semi *lex8.Token

	index uint32 // symbol decl index
}
