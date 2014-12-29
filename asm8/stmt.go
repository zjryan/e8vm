package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type stmt struct {
	*inst
	label string

	ops []*lex8.Token

	offset uint32
}
