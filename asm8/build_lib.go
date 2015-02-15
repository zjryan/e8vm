package asm8

func buildPkgScope(b *builder, pkg *pkg) {
	decl := func(sym *symbol) bool {
		exists := b.scope.Declare(sym)
		if exists != nil {
			b.err(sym.Pos, "%q already declared", sym.Name)
			b.err(exists.Pos, "  previously declared here")
			return false
		}
		return true
	}

	for _, file := range pkg.Files {
		// declare functions
		for _, fn := range file.Funcs {
			t := fn.name
			sym := &symbol{
				t.Lit,
				SymFunc,
				fn,
				t.Pos,
				pkg.Path,
			}

			if !decl(sym) {
				continue
			}

			fn.index = b.curPkg.Declare(sym) // assign link index
		}

		// declare variables
		for _, v := range file.Vars {
			t := v.name
			sym := &symbol{
				t.Lit,
				SymVar,
				v,
				t.Pos,
				pkg.Path,
			}

			if !decl(sym) {
				continue
			}

			v.index = b.curPkg.Declare(sym) // assign link index
		}
	}
}

func buildLib(b *builder, pkg *pkg) *Lib {
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
