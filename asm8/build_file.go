package asm8

func buildFile(b *builder, f *file) {
	b.scope.Push() // file scope
	defer b.scope.Pop()

	// TODO: import required packages, and add them into the symbol table

	pkg := b.curPkg
	for _, fn := range f.Funcs {
		if obj := buildFunc(b, fn); obj != nil {
			pkg.DefineFunc(fn.index, obj)
		}
	}

	for _, v := range f.Vars {
		if obj := buildVar(b, v); obj != nil {
			pkg.DefineVar(v.index, obj)
		}
	}
}
