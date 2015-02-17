package build8

import (
	"lonnie.io/e8vm/lex8"
)

// BuildAll packages under the home path
func BuildAll(homePath string) []*lex8.Error {
	b := NewBuilder(homePath)

	pkgs, e := b.ListPkgs()
	if e != nil {
		return lex8.SingleErr(e)
	}

	for _, p := range pkgs {
		es := b.Build(p)
		if es != nil {
			return es
		}
	}

	return nil
}
