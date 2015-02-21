package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Inst is an assembly instruction.
type Inst struct {
	Inst uint32
	Pkg  string
	Sym  string

	Fill int

	SymTok *lex8.Token
}
