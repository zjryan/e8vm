package ir

import (
	"lonnie.io/e8vm/link8"
)

// BuildPkg builds a package and returns the built lib
func BuildPkg(p *Pkg) *link8.Pkg {
	for _, f := range p.funcs {
		f.index = p.lib.DeclareFunc(f.name)
	}

	for _, f := range p.funcs {
		genFunc(p, f)
		writeFunc(p, f)
	}

	return p.lib
}
