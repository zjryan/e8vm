package link8

// pkgSym is a reference to a symbol in a package
type pkgSym struct {
	pkg *Pkg
	sym uint32
}

func (ps pkgSym) Type() int {
	return ps.pkg.symbols[ps.sym].Type
}

func (ps pkgSym) Func() *Func {
	return ps.pkg.Func(ps.sym)
}

func (ps pkgSym) Var() *Var {
	return ps.pkg.Var(ps.sym)
}

func (ps pkgSym) Import(index uint32) *Pkg {
	return ps.pkg.requires[index]
}
