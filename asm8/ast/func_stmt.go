package ast

import (
	"lonnie.io/e8vm/lex8"
)

type FuncStmt struct {
	*Inst
	Label string

	Ops []*lex8.Token

	Offset uint32
}

func (s *FuncStmt) IsLabel() bool {
	return s.Inst == nil && s.Label != ""
}
