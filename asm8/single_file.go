package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildSingleFile builds a package named "main" from a single file.
func BuildSingleFile(f string, rc io.ReadCloser) ([]byte, []*lex8.Error) {
	files := map[string]io.ReadCloser{f: rc}
	lib, es := Lang.Compile(files, nil)
	if es != nil {
		return nil, es
	}

	buf := new(link8.Buf)
	e := link8.LinkMain(lib.Lib(), buf)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return buf.Bytes(), nil
}
