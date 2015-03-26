package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Expr is a general expression in the language
type Expr interface{}

// Operand is an operand expression
type Operand struct {
	*lex8.Token
}

// OpExpr is a binary or unary operation that uses an operator
type OpExpr struct {
	A  Expr
	Op *lex8.Token
	B  Expr
}

// ParenExpr is an expression in a pair of parenthesis
type ParenExpr struct {
	Lparen *lex8.Token
	Rparen *lex8.Token
	Expr
}

// ExprList is a list of expressions
type ExprList struct {
	Exprs  []Expr
	Commas []*lex8.Token
}

// Len returns the length of the expression list
func (list *ExprList) Len() int { return len(list.Exprs) }

// CallExpr is a function call expression
type CallExpr struct {
	Func   Expr
	Lparen *lex8.Token
	Rparen *lex8.Token
	Args   *ExprList
}
