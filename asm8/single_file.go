package asm8

import (
	"bytes"
	"io"

	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildSingleFile builds a package named "main" from a single file.
func BuildSingleFile(f string, rc io.ReadCloser) ([]byte, []*lex8.Error) {
	pkg := build8.NewSimplePkg(f, rc)

	es := Lang().Compile(pkg)
	if es != nil {
		return nil, es
	}

	buf := new(bytes.Buffer)
	e := link8.LinkMain(pkg.Compiled().Lib(), buf, "main")
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return buf.Bytes(), nil
}
