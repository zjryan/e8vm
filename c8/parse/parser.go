package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

type parser struct {
	*lex8.ErrorList
}

func newParser(f string, rc io.Reader) *parser {
	ret := new(parser)
	return ret
}
