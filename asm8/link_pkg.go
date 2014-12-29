package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func linkPkg(main *pkgObj) ([]byte, *lex8.Error) {
	lnk := newLinker()
	lnk.addPkgs(main)

	funcMain, index := main.Query("main")
	if funcMain.Type != SymFunc {
		return nil, lex8.Errorf("main function missing in %s", main.path)
	}

	used := traceUsed(lnk, main, index)

	funcs, vars, e := layout(used)
	if e != nil {
		return nil, e
	}

	w := newWriter()
	for _, f := range funcs {
		writeFunc(w, f)
	}

	for _, v := range vars {
		panic(v) // TODO: write data
	}

	return nil, nil
}
