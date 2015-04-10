package asm8

func buildFile(b *builder, f *file) {
	// TODO: import required packages, and add them into the symbol table

	pkg := b.curPkg
	for _, fn := range f.funcs {
		if obj := buildFunc(b, fn); obj != nil {
			ind := b.getIndex(fn.Name.Lit)
			pkg.DefineFunc(ind, obj)
		}
	}

	for _, v := range f.vars {
		if obj := buildVar(b, v); obj != nil {
			ind := b.getIndex(v.Name.Lit)
			pkg.DefineVar(ind, obj)
		}
	}
}
