package link8

import (
	"path"
)

// Pkg is the compiling object of a package. It is the linking
// unit for programs.
type Pkg struct {
	path string

	requires []*Pkg    // all the packages that requires for building
	symbols  []*Symbol // all the symbol objects

	symIndex map[string]uint32 // map from symbol names to index in symNames
	pkgIndex map[string]uint32 // map from package path to index in imports

	funcs map[uint32]*Func
	vars  map[uint32]*Var
}

// NewPkg creates a new package for path p.
func NewPkg(p string) *Pkg {
	ret := new(Pkg)
	ret.path = p

	ret.symIndex = make(map[string]uint32)
	ret.pkgIndex = make(map[string]uint32)

	ret.funcs = make(map[uint32]*Func)
	ret.vars = make(map[uint32]*Var)

	index := ret.Require(ret)
	if index != 0 {
		panic("bug")
	}

	return ret
}

// Name returns the package's default name.
func (p *Pkg) Name() string { return path.Base(p.path) }

// Path returns the package's path string.
func (p *Pkg) Path() string { return p.path }

// Require assigns a relative index for the required package.
func (p *Pkg) Require(req *Pkg) uint32 {
	if index, found := p.pkgIndex[req.path]; found {
		return index
	}

	index := uint32(len(p.requires))
	p.pkgIndex[req.path] = index
	p.requires = append(p.requires, req)
	return index
}

// PkgIndex returns the package imported and also its import
// index.
func (p *Pkg) PkgIndex(name string) (*Pkg, uint32) {
	index, found := p.pkgIndex[name]
	if !found {
		return nil, 0
	}

	return p.requires[index], index
}

// SymIndex returns the index of a symbol in the package.
// It panics when the symbol is not Declare()'d yet.
func (p *Pkg) SymIndex(name string) uint32 {
	ret, found := p.symIndex[name]
	if !found {
		panic("not found")
	}
	return ret
}

// Declare declares a symbol and assigns a symbol index.
func (p *Pkg) Declare(s *Symbol) uint32 {
	_, found := p.symIndex[s.Name]
	if found {
		panic("redeclaring")
	}

	index := uint32(len(p.symbols))
	p.symIndex[s.Name] = index
	p.symbols = append(p.symbols, s)
	return index
}

// Query returns the symbol with the particular name.
func (p *Pkg) Query(name string) (*Symbol, uint32) {
	index, found := p.symIndex[name]
	if !found {
		return nil, 0
	}

	return p.symbols[index], index
}

// HasFunc checks if the package has a function of a particular name.
func (p *Pkg) HasFunc(name string) bool {
	sym, _ := p.Query(name)
	if sym == nil || sym.Type != SymFunc {
		return false
	}
	return true
}

// DefineFunc instantiates a function object for a particular index.
func (p *Pkg) DefineFunc(index uint32, f *Func) {
	sym := p.symbols[index]
	if sym.Type != SymFunc {
		panic("not a function")
	}

	p.funcs[index] = f
}

// DefineVar instantiates a variable object for a particular index.
func (p *Pkg) DefineVar(index uint32, v *Var) {
	sym := p.symbols[index]
	if sym.Type != SymVar {
		panic("not a var")
	}

	p.vars[index] = v
}

// Func returns the function of index.
func (p *Pkg) Func(index uint32) *Func {
	ret, found := p.funcs[index]
	if !found {
		panic("not found")
	}
	return ret
}

// Var returns the variable of index.
func (p *Pkg) Var(index uint32) *Var {
	ret, found := p.vars[index]
	if !found {
		panic("not found")
	}
	return ret
}
