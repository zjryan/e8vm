package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type file struct {
	*ast.File

	funcs   []*funcDecl
	vars    []*varDecl
	imports *importDecl
}

func resolveFile(log lex8.Logger, f *ast.File) *file {
	ret := new(file)
	ret.File = f

	if f.Imports != nil {
		ret.imports = resolveImportDecl(log, f.Imports)
	}

	for _, d := range f.Decls {
		if d, ok := d.(*ast.FuncDecl); ok {
			ret.funcs = append(ret.funcs, resolveFunc(log, d))
		}

		if d, ok := d.(*ast.VarDecl); ok {
			ret.vars = append(ret.vars, resolveVar(log, d))
		}
	}

	return ret
}
