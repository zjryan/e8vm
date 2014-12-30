package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type funcStmt struct {
	*inst
	label string

	ops []*lex8.Token

	offset uint32
}

func (s *funcStmt) isLabel() bool {
	return s.inst == nil && s.label != ""
}
