package build8

import (
	"lonnie.io/e8vm/lex8"
)

// BuildAll packages covered by a builder
func BuildAll(b *Builder) []*lex8.Error {
	pkgs, e := b.ListPkgs()
	if e != nil {
		return lex8.SingleErr(e)
	}

	for _, p := range pkgs {
		_, es := b.prepare(p)
		if es != nil {
			return es
		}
	}

	for _, p := range pkgs {
		_, es := b.build(p)
		if es != nil {
			return es
		}
	}

	return nil
}
