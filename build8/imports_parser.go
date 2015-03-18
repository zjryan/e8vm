package build8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

type importsParser struct {
	*lex8.Parser
}

var types = func() *lex8.Types {
	ret := lex8.NewTypes()
	o := func(t int, name string) {
		ret.Register(t, name)
	}

	o(operand, "operand")
	o(semi, ";")
	o(endl, "end-line")
	return ret
}()

func newImportsParser(file string, r io.Reader) *importsParser {
	ret := new(importsParser)
	x := newImportsLexer(file, r)
	ret.Parser = lex8.NewParser(x, types)

	return ret
}
