package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Expr is a general expression in the language
type Expr interface{}

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

// CallExpr is a function call expression
type CallExpr struct {
	Func   Expr
	Lparen *lex8.Token
	Rparen *lex8.Token
	Args   *ExprList
}