package ast

import (
	"fmt"

	"lonnie.io/e8vm/lex8"
)

// ExprPos returns the starting position of an expression.
func ExprPos(e Expr) *lex8.Pos {
	switch e := e.(type) {
	case *Operand:
		return e.Token.Pos
	case *OpExpr:
		return ExprPos(e.A)
	case *ParenExpr:
		return e.Lparen.Pos
	case *ExprList:
		if len(e.Exprs) == 0 {
			return nil
		}
		return ExprPos(e.Exprs[0])
	case *CallExpr:
		return ExprPos(e.Func)
	default:
		panic(fmt.Errorf("invalid expression type: %T", e))
	}
}
