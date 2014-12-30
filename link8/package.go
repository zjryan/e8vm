package link8

import (
	"path"
)

// Package is the compiling object of a package. It is the linking
// unit for programs.
type Package struct {
	path string

	requires []*Package // all the packages that requires for building
	symbols  []*Symbol  // all the symbol objects

	symIndex map[string]uint32 // map from symbol names to index in symNames
	pkgIndex map[string]uint32 // map from package path to index in imports

	funcs map[uint32]*Func
	vars  map[uint32]*Var
}

// NewPackage creates a new package of path p.
func NewPackage(p string) *Package {
	ret := new(Package)
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
func (p *Package) Name() string { return path.Base(p.path) }

// Path returns the package's path string.
func (p *Package) Path() string { return p.path }

// Require assigns a relative index for the required package.
func (p *Package) Require(req *Package) uint32 {
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
func (p *Package) PkgIndex(name string) (*Package, uint32) {
	index, found := p.pkgIndex[name]
	if !found {
		return nil, 0
	}

	return p.requires[index], index
}

// SymIndex returns the index of a symbol in the package.
// It panics when the symbol is not Declare()'d yet.
func (p *Package) SymIndex(name string) uint32 {
	ret, found := p.symIndex[name]
	if !found {
		panic("not found")
	}
	return ret
}

// Declare declares a symbol and assigns a symbol index.
func (p *Package) Declare(s *Symbol) uint32 {
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
func (p *Package) Query(name string) (*Symbol, uint32) {
	index, found := p.symIndex[name]
	if !found {
		return nil, 0
	}

	return p.symbols[index], index
}

// DefineFunc instantiates a function object for a particular index.
func (p *Package) DefineFunc(index uint32, f *Func) {
	sym := p.symbols[index]
	if sym.Type != SymFunc {
		panic("not a function")
	}

	p.funcs[index] = f
}

// DefineVar instantiates a variable object for a particular index.
func (p *Package) DefineVar(index uint32, v *Var) {
	sym := p.symbols[index]
	if sym.Type != SymVar {
		panic("not a var")
	}

	p.vars[index] = v
}

// Func returns the function of index.
func (p *Package) Func(index uint32) *Func {
	ret, found := p.funcs[index]
	if !found {
		panic("not found")
	}
	return ret
}

// Var returns the variable of index.
func (p *Package) Var(index uint32) *Var {
	ret, found := p.vars[index]
	if !found {
		panic("not found")
	}
	return ret
}
