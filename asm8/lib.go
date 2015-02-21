package asm8

import (
	"lonnie.io/e8vm/link8"
)

// Lib is the compiler output of a package
// it contains the package for linking,
// and also the symbols for importing
type lib struct {
	*link8.Pkg

	symbols map[string]*symbol
}

// NewPkgObj creates a new package compile object
func newLib(p string) *lib {
	ret := new(lib)
	ret.Pkg = link8.NewPackage(p)

	ret.symbols = make(map[string]*symbol)

	id := ret.Require(ret.Pkg) // require itself
	if id != 0 {
		panic("bug")
	}

	return ret
}

// Link returns the link8.Package for linking.
func (p *lib) Link() *link8.Pkg { return p.Pkg }

// Declare declares a symbol inside the package.  If the symbol is a function
// or variable, it is also declared as an object file symbol in the underlying
// link8.Package, and it returns the index.  If the symbol is a constant, it
// returns 0 after the declaration. Other types will panic. Redeclaration will
// panic.
func (p *lib) Declare(s *symbol) uint32 {
	_, found := p.symbols[s.Name]
	if found {
		panic("redeclare")
	}
	p.symbols[s.Name] = s

	switch s.Type {
	case SymConst:
		return 0
	case SymFunc:
		return p.Pkg.Declare(&link8.Symbol{
			Name: s.Name,
			Type: link8.SymFunc,
		})
	case SymVar:
		return p.Pkg.Declare(&link8.Symbol{
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
func (p *lib) query(name string) (*symbol, uint32) {
	ret, found := p.symbols[name]
	if !found {
		return nil, 0
	}

	switch ret.Type {
	case SymConst:
		return ret, 0
	case SymFunc, SymVar:
		s, index := p.Pkg.Query(name)
		if s == nil {
			panic("symbol missing")
		}
		return ret, index
	default:
		panic("bug")
	}
}
