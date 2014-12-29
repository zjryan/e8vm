package asm8

import (
	"lonnie.io/e8vm/link8"
)

type PkgObj struct {
	*link8.Package

	symbols map[string]*Symbol
}

func newPkgObj(p string) *PkgObj {
	ret := new(PkgObj)
	ret.Package = link8.NewPackage(p)
	ret.symbols = make(map[string]*Symbol)

	return ret
}

func (p *PkgObj) Declare(s *Symbol) uint32 {
	_, found := symbols[s.Name]
	if found {
		panic("redeclare")
	}
	symbols[s.Name] = s

	switch s.Type {
	case SymConst:
		return 0
	case SymFunc:
		return p.Package.Declare(&link8.Symbol {
			Name: s.Name,
			Type: link8.SymFunc,
		})
	case SymVar:
		return p.Package.Declare(&link8.Symbol {
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
