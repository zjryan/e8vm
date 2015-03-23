package parse

import (
	"lonnie.io/e8vm/lex8"
)

// asm8 token types.
const (
	Keyword = iota
	Operand
	String
	Lbrace
	Rbrace
	Endl
	Semi
)

// Types provides a type name querier
var Types = func() *lex8.Types {
	ret := lex8.NewTypes()
	o := func(t int, name string) {
		ret.Register(t, name)
	}

	o(Keyword, "keyword")
	o(Operand, "operand")
	o(Lbrace, "'{'")
	o(Rbrace, "'}'")
	o(String, "string")
	o(Semi, "';'")
	o(Endl, "end-line")

	return ret
}()

// TypeStr returns the name of a token type.
func TypeStr(t int) string { return Types.Name(t) }
