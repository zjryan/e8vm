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

func newPkg(path string) *Pkg {
	ret := new(Pkg)
	ret.path = path
	return ret
}

func (p *Pkg) newFunc() *Func {
	ret := new(Func)
	ret.id = len(p.funcs)
	return ret
}
