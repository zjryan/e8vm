package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildBareFunc builds a function body into an image.
func BuildBareFunc(f string, r io.Reader) ([]byte, []*lex8.Error) {
	p := newParser(f, r)
	fn := parseBareFunc(p)
	if es := p.Errs(); es != nil {
		return nil, es
	}

	b := newBuilder()
	fobj := buildFunc(b, fn)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	ret, e := link8.LinkBareFunc(fobj)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return ret, nil
}
