package ast

import (
	"lonnie.io/e8vm/lex8"
)

// FuncStmt is a statement in a assembly function.
// It is either a instruction or a label.
type FuncStmt struct {
	*Inst
	Label string

	Ops []*lex8.Token

	Offset uint32
}

// IsLabel checks if the statement is a label
func (s *FuncStmt) IsLabel() bool {
	return s.Inst == nil && s.Label != ""
}
