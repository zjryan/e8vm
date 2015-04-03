package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func parseConstDecls(p *parser) *ast.ConstDecls {
	p.ErrorfHere("const declare not implemented")
	p.Next()
	return nil
}

func parseVarDecls(p *parser) *ast.VarDecls {
	p.ErrorfHere("var declare not implemented")
	p.Next()
	return nil
}
