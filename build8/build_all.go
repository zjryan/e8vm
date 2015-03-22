package build8

import (
	"lonnie.io/e8vm/lex8"
)

// BuildAll packages under the home path
func BuildAll(homePath string, verbose bool, lang Lang) []*lex8.Error {
	b := NewBuilder(homePath)
	b.AddLang(lang)
	b.Verbose = verbose

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
