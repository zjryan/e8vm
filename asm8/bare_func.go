package asm8

import (
	"io"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildBareFunc builds a function body into an image.
func BuildBareFunc(f string, rc io.ReadCloser) ([]byte, []*lex8.Error) {
	fn, es := parse.BareFunc(f, rc)
	if es != nil {
		return nil, es
	}

	elist := lex8.NewErrorList()
	r := resolveFunc(elist, fn)
	if elist.Errs != nil {
		return nil, elist.Errs
	}

	b := newBuilder()
	fobj := buildFunc(b, r)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	ret, e := link8.LinkBareFunc(fobj)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return ret, nil
}
