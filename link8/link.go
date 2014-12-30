package link8

import (
	"fmt"
)

// LinkMain produces a image for a main() in a package.
func LinkMain(main *Package) ([]byte, error) {
	lnk := NewLinker()
	lnk.AddPkgs(main)

	funcMain, index := main.Query("main")
	if funcMain == nil || funcMain.Type != SymFunc {
		return nil, fmt.Errorf("main function missing")
	}

	used := traceUsed(lnk, main, index)

	funcs, vars, e := layout(used)
	if e != nil {
		return nil, e
	}

	w := newWriter()
	for _, f := range funcs {
		writeFunc(w, f.pkg, f.Func())
	}

	for _, v := range vars {
		writeVar(w, v.Var())
	}

	return w.bytes(), nil
}

// LinkBareFunc produces a image of a single function that has no links.
func LinkBareFunc(f *Func) ([]byte, error) {
	if f.TooLarge() {
		return nil, fmt.Errorf("code section too large")
	}

	w := newWriter()
	w.writeBareFunc(f)
	return w.bytes(), nil
}
