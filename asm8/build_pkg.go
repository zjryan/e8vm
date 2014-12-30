package asm8

func buildPkgScope(b *builder, pkg *pkg) {
	for _, file := range pkg.Files {
		for _, fn := range file.Funcs {
			t := fn.name
			sym := &symbol{
				fn.name.Lit,
				SymFunc,
				fn,
				t.Pos,
				pkg.Path,
			}
			exists := b.scope.Declare(sym)
			if exists != nil {
				b.err(t.Pos, "%q already declared", t.Lit)
				b.err(exists.Pos, "  previously declared here")
				continue
			}

			// also declare this symbol in package object
			index := b.curPkg.Declare(sym)
			fn.index = index
		}
	}
}

func buildPkg(b *builder, pkg *pkg) *lib {
	ret := newLib(pkg.Path)
	b.curPkg = ret

	b.scope.Push()
	defer b.scope.Pop()

	buildPkgScope(b, pkg)
	if b.Errs() != nil {
		return nil // error on declaring, so just return
	}

	for _, file := range pkg.Files {
		buildFile(b, file)
	}
	if b.Errs() != nil {
		return nil // error on building
	}

	return ret
}
