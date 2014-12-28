package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Func is an assembly function.
type Func struct {
	stmts []*stmt

	kw, name             *lex8.Token
	lbrace, rbrace, semi *lex8.Token

	addr uint32
	size uint32
}
