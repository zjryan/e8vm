package parse

import (
	"io"

	"lonnie.io/e8vm/c8/ast"
	"lonnie.io/e8vm/lex8"
)

// Expr parses a bare expression and returns the ast node.
func Expr(f string, rc io.Reader) (ast.Expr, []*lex8.Error) {
	panic("todo")
}
