package asm8

func buildFile(b *builder, f *file) {
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
