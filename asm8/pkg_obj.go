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

// Require assigns a relative index for the required package.
func (p *pkgObj) Require(req *pkgObj) uint32 {
	if index, found := p.pkgIndex[req.path]; found {
		return index
	}

	index := uint32(len(p.requires))
	p.pkgIndex[req.path] = index
	p.requires = append(p.requires, req)
	return index
}

// PkgIndex returns the relative package index of a required package.
// It panics when the package is not Require()'d yet.
func (p *pkgObj) PkgIndex(name string) uint32 {
	ret, found := p.pkgIndex[name]
	if !found {
		panic("not found")
	}
	return ret
}

// SymIndex returns the index of a symbol in the package.
// It panics when the symbol is not Declare()'d yet.
func (p *pkgObj) SymIndex(name string) uint32 {
	ret, found := p.symIndex[name]
	if !found {
		panic("not found")
	}
	return ret
}

// Declare declares a symbol and assigns a symbol index.
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

// Query returns the symbol with the particular name.
func (p *pkgObj) Query(name string) (*Symbol, uint32) {
	index, found := p.symIndex[name]
	if !found {
		return nil, 0
	}

	return p.symbols[index], index
}

type varObj struct{}
type constObj struct{}
