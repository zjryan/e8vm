package ast

import (
	"lonnie.io/e8vm/lex8"
)

// VarStmt is a variable statement.
type VarStmt struct {
	Type *lex8.Token
	Args []*lex8.Token

	// resolved
	Align uint32
	Data  []byte
}
