package parse

import (
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

// Expr parses a bare expression and returns the ast node.
func Expr(f string, rc io.ReadCloser) (ast.Expr, []*lex8.Error) {
	// p, _ := newParser(f, rc)
	defer rc.Close()
	panic("todo")
}
