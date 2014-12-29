package asm8

import (
	"path"
)

type pkgObj struct {
	name string
	path string

	requires []*pkgObj // all the packages that requires for building
	symbols  []*Symbol // all the symbol objects

	symIndex map[string]uint32 // map from symbol names to index in symNames
	pkgIndex map[string]uint32 // map from package path to index in imports

	funcs  map[uint32]*funcObj
	vars   map[uint32]*varObj
	consts map[uint32]*constObj

	absIndex uint32 // absolute index for linking
}

func newPkgObj(p string) *pkgObj {
	ret := new(pkgObj)
	ret.name = path.Base(p)
	ret.path = p

	ret.symIndex = make(map[string]uint32)
	ret.pkgIndex = make(map[string]uint32)

	ret.funcs = make(map[uint32]*funcObj)
	ret.vars = make(map[uint32]*varObj)
	ret.consts = make(map[uint32]*constObj)

	index := ret.Require(ret)
	if index != 0 {
		panic("bug")
	}

	return ret
}

func (p *pkgObj) Require(req *pkgObj) uint32 {
	if index, found := p.pkgIndex[req.path]; found {
		return index
	}

	index := uint32(len(p.requires))
	p.pkgIndex[req.path] = index
	p.requires = append(p.requires, req)
	return index
}

func (p *pkgObj) Declare(s *Symbol) uint32 {
	_, found := p.symIndex[s.Name]
	if found {
		panic("redeclaring")
	}

	index := uint32(len(p.symbols))
	p.symIndex[s.Name] = index
	p.symbols = append(p.symbols, s)
	return index
}

func (p *pkgObj) Query(name string) *Symbol {
	index, found := p.symIndex[name]
	if !found {
		return nil
	}

	return p.symbols[index]
}

type varObj struct{}
type constObj struct{}
