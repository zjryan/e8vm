package ast

import (
	"lonnie.io/e8vm/lex8"
)

// FuncStmt is a statement in a assembly function.
// It is either a instruction or a label.
type FuncStmt struct {
	Ops []*lex8.Token
}
