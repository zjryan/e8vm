package asm8

import (
	"lonnie.io/e8vm/link8"
)

// PkgObj is a package object.
type PkgObj struct {
	*link8.Package

	requires map[uint32]*PkgObj
	symbols  map[string]*Symbol
}

// NewPkgObj creates a new package compile object
func NewPkgObj(p string) *PkgObj {
	ret := new(PkgObj)
	ret.Package = link8.NewPackage(p)
	ret.symbols = make(map[string]*Symbol)

	return ret
}

// Require imports a package in and grants the package
// a import index.
func (p *PkgObj) Require(req *PkgObj) uint32 {
	ret := p.Package.Require(req.Package)
	_, found := p.requires[ret]
	if !found {
		p.requires[ret] = req
	}

	return ret
}

// PkgIndex returns the package import index, consistent with
// the underlying link8.Package.
func (p *PkgObj) PkgIndex(path string) (*PkgObj, uint32) {
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

func (p *PkgObj) Declare(s *Symbol) uint32 {
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

func (p *PkgObj) Query(name string) (*Symbol, uint32) {
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
		panic("invalid query")
	}
}
