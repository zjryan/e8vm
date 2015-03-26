package link8

import (
	"fmt"
	"io"
)

// LinkMain produces a image for a main() in a package.
func LinkMain(main *Pkg, out io.Writer) error {
	lnk := NewLinker()
	lnk.AddPkgs(main)

	funcMain, index := main.SymbolByName("main")
	if funcMain == nil || funcMain.Type != SymFunc {
		return fmt.Errorf("main function missing")
	}

	used := traceUsed(lnk, main, index)

	funcs, vars, e := layout(used)
	if e != nil {
		return e
	}

	w := newWriter(out)
	for _, f := range funcs {
		writeFunc(w, f.pkg, f.Func())
	}
	for _, v := range vars {
		writeVar(w, v.Var())
	}
	return w.Err()
}

// LinkBareFunc produces a image of a single function that has no links.
func LinkBareFunc(f *Func) ([]byte, error) {

	if f.TooLarge() {
		return nil, fmt.Errorf("code section too large")
	}

	buf := new(Buf)
	w := newWriter(buf)
	w.writeBareFunc(f)
	e := w.Err()
	if e != nil {
		return nil, e
	}

	return buf.Bytes(), nil
}
