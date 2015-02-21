package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
)

func buildFile(b *builder, f *ast.File) {
	b.scope.Push() // file scope
	defer b.scope.Pop()

	// TODO: import required packages, and add them into the symbol table

	pkg := b.curPkg
	for _, fn := range f.Funcs {
		if obj := buildFunc(b, fn); obj != nil {
			pkg.DefineFunc(fn.Index, obj)
		}
	}

	for _, v := range f.Vars {
		if obj := buildVar(b, v); obj != nil {
			pkg.DefineVar(v.Index, obj)
		}
	}
}
