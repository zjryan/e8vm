package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Expr is a general expression in the language
type Expr interface{}

// BinaryOp is an binary operation
type BinaryOp struct {
	E1 Expr
	Op *lex8.Token
	E2 Expr
}
