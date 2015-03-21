package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type file struct {
	*ast.File

	funcs []*funcDecl
	vars  []*varDecl
}

func resolveFile(log lex8.Logger, f *ast.File) *file {
	ret := new(file)
	ret.File = f

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

func resolveFiles(log lex8.Logger, files []*ast.File) []*file {
	ret := make([]*file, 0, len(files))

	for _, f := range files {
		ret = append(ret, resolveFile(log, f))
	}
	return ret
}
