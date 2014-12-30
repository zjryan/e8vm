package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type varStmt struct {
	typ  *lex8.Token
	args []*lex8.Token

	align uint32
	data  []byte
}
