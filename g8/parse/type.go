package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func parseType(p *parser) ast.Expr {
	if p.See(Ident) {
		return &ast.Operand{p.Shift()}
	}

	// TODO: pointer, slice, and array type support
	p.ErrorfHere("expect a type name")
	return nil
}
