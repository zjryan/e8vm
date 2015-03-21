package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	// "lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/lex8"
)

type pkg struct {
	*ast.Pkg

	files []*file
}

func resolvePkg(log lex8.Logger, p *ast.Pkg) *pkg {
	ret := new(pkg)
	ret.Pkg = p

	ret.files = make([]*file, len(p.Files))
	for i, f := range p.Files {
		ret.files[i] = resolveFile(log, f)
	}

	return ret
}
