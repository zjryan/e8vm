package asm8

import (
	"lonnie.io/e8vm/link8"
)

// Lib is the compiler output of a package
// it contains the package for linking,
// and also the symbols for importing
type Lib struct {
	*link8.Package

	requires map[uint32]*Lib
	symbols  map[string]*symbol
}

// NewPkgObj creates a new package compile object
func newLib(p string) *Lib {
	ret := new(Lib)
	ret.Package = link8.NewPackage(p)

	ret.requires = make(map[uint32]*Lib)
	ret.symbols = make(map[string]*symbol)

	id := ret.Require(ret)
	if id != 0 {
		panic("bug")
	}

	return ret
}

// Require imports a package in and grants the package
// a import index.
func (p *Lib) Require(req *Lib) uint32 {
	ret := p.Package.Require(req.Package)
	_, found := p.requires[ret]
	if !found {
		p.requires[ret] = req
	}

	return ret
}

// PkgIndex returns the package import index, consistent with
// the underlying link8.Package.
func (p *Lib) LibIndex(path string) (*Lib, uint32) {
	pkg, index := p.Package.PkgIndex(path)
	if pkg == nil {
		return nil, 0
	}

	ret, found := p.requires[index]
	if !found {
		panic("bug")
	}

	return ret, index
}

// Declare declares a symbol inside the package.  If the symbol is a function
// or variable, it is also declared as an object file symbol in the underlying
// link8.Package, and it returns the index.  If the symbol is a constant, it
// returns 0 after the declaration. Other types will panic. Redeclaration will
// panic.
func (p *Lib) Declare(s *symbol) uint32 {
	_, found := p.symbols[s.Name]
	if found {
		panic("redeclare")
	}
	p.symbols[s.Name] = s

	switch s.Type {
	case SymConst:
		return 0
	case SymFunc:
		return p.Package.Declare(&link8.Symbol{
			Name: s.Name,
			Type: link8.SymFunc,
		})
	case SymVar:
		return p.Package.Declare(&link8.Symbol{
			Name: s.Name,
			Type: link8.SymVar,
		})
	default:
		panic("declare with invalid sym type")
	}
}

// Query returns the symbol declared by name and its symbol index
// if the symbol is a function or variable. It returns nil, 0 when
// the symbol of name is not found.
func (p *Lib) Query(name string) (*symbol, uint32) {
	ret, found := p.symbols[name]
	if !found {
		return nil, 0
	}

	switch ret.Type {
	case SymConst:
		return ret, 0
	case SymFunc, SymVar:
		s, index := p.Package.Query(name)
		if s == nil {
			panic("symbol missing")
		}
		return ret, index
	default:
		panic("bug")
	}
}
