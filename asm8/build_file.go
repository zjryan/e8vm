package asm8

func buildFile(b *builder, f *file) {
	b.scope.Push() // file scope
	defer b.scope.Pop()

	// TODO: import required packages, and add them into the symbol table

	pkg := b.curPkg
	for _, f := range f.Funcs {
		obj := buildFunc(b, f)
		if obj != nil {
			pkg.DefineFunc(f.index, obj)
		}
	}
}
