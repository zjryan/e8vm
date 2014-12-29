package link8

// pkgSym is a reference to a symbol in a package
type pkgSym struct {
	pkg *Package
	sym uint32
}

func (ps pkgSym) Type() int {
	return ps.pkg.symbols[ps.sym].Type
}

func (ps pkgSym) Func() *Func {
	return ps.pkg.Func(ps.sym)
}

func (ps pkgSym) Import(index uint32) *Package {
	return ps.pkg.requires[index]
}
