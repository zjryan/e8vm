package ir

import (
	"lonnie.io/e8vm/link8"
)

func genPkg(p *Pkg) *link8.Pkg {
	p.lib = link8.NewPkg(p.path)

	for _, f := range p.funcs {
		f.index = p.lib.DeclareFunc(f.name)
	}

	for _, f := range p.funcs {
		genFunc(p, f)
		writeFunc(p, f)
	}

	return p.lib
}
