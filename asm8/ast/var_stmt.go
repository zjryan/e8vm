package ast

import (
	"lonnie.io/e8vm/lex8"
)

type VarStmt struct {
	Type *lex8.Token
	Args []*lex8.Token

	Align uint32
	Data  []byte
}
