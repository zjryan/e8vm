package ir

import (
	"lonnie.io/e8vm/link8"
)

// Pkg is a package in its intermediate representation.
type Pkg struct {
	lib *link8.Pkg

	path string

	funcs []*Func
	vars  []*heapSym
}

// NewPkg creates a package with a particular path name.
func NewPkg(path string) *Pkg {
	ret := new(Pkg)
	ret.path = path
	ret.lib = link8.NewPkg(path)

	return ret
}

// NewFunc creates a new function from the package.
func (p *Pkg) NewFunc(name string, sig *FuncSig) *Func {
	ret := newFunc(name, len(p.funcs), sig)
	p.funcs = append(p.funcs, ret)
	return ret
}

// Require imporpts a linkable package.
func (p *Pkg) Require(pkg *link8.Pkg) uint32 { return p.lib.Require(pkg) }
