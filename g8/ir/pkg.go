package ir

import (
	"lonnie.io/e8vm/link8"
)

// Pkg is a package in its intermediate representation.
type Pkg struct {
	lib *link8.Pkg

	path string

	funcs []*Func
	vars  []*heapVar
}

// NewPkg creates a package with a particular path name.
func NewPkg(path string) *Pkg {
	ret := new(Pkg)
	ret.path = path
	return ret
}

// NewFunc creates a new function from the package.
func (p *Pkg) NewFunc(name string) *Func {
	ret := new(Func)
	ret.id = len(p.funcs)
	ret.name = name

	p.funcs = append(p.funcs, ret)
	return ret
}
