package parse

import (
	"lonnie.io/e8vm/g8/ast"
)

func parseConstDecls(p *parser) *ast.ConstDecls {
	if !p.SeeKeyword("const") {
		panic("expect keyword")
	}

	p.ErrorfHere("const declare not implemented")
	p.Next()
	return nil
}

func parseVarDecls(p *parser) *ast.VarDecls {
	if !p.SeeKeyword("var") {
		panic("expect keyword")
	}

	p.ErrorfHere("var declare not implemented")
	p.Next()
	return nil
}

func parseStruct(p *parser) *ast.Struct {
	if !p.SeeKeyword("struct") {
		panic("expect keyword")
	}

	p.ErrorfHere("struct declare not implemented")
	p.Next()
	return nil
}
