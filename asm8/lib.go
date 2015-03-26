package asm8

import (
	"io"

	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/sym8"
)

// Lib is the compiler output of a package
// it contains the package for linking,
// and also the symbols for importing
type lib struct {
	*link8.Pkg

	symbols map[string]*sym8.Symbol
}

// NewPkgObj creates a new package compile object
func newLib(p string) *lib {
	ret := new(lib)
	ret.Pkg = link8.NewPkg(p)

	ret.symbols = make(map[string]*sym8.Symbol)

	id := ret.Require(ret.Pkg) // require itself
	if id != 0 {
		panic("bug")
	}

	return ret
}

// Link returns the link8.Package for linking.
func (p *lib) Link() *link8.Pkg { return p.Pkg }

func (p *lib) Declare(s *sym8.Symbol) uint32 {
	_, found := p.symbols[s.Name()]
	if found {
		panic("redeclare")
	}
	p.symbols[s.Name()] = s

	switch s.Type {
	case SymConst:
		panic("todo")
	case SymFunc:
		return p.Pkg.DeclareFunc(s.Name())
	case SymVar:
		return p.Pkg.DeclareVar(s.Name())
	default:
		panic("declare with invalid sym type")
	}
}

// Query returns the symbol declared by name and its symbol index
// if the symbol is a function or variable. It returns nil, 0 when
// the symbol of name is not found.
func (p *lib) query(name string) (*sym8.Symbol, uint32) {
	ret, found := p.symbols[name]
	if !found {
		return nil, 0
	}

	switch ret.Type {
	case SymConst:
		return ret, 0
	case SymFunc, SymVar:
		s, index := p.Pkg.SymbolByName(name)
		if s == nil {
			panic("symbol missing")
		}
		return ret, index
	default:
		panic("bug")
	}
}

// Lib retunrs the linkable lib.
func (p *lib) Lib() *link8.Pkg {
	return p.Pkg
}

// Save marshalls the library out.
func (p *lib) Save(w io.Writer) error {
	panic("todo")
}
