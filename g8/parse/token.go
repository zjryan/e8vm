package parse

import (
	"lonnie.io/e8vm/lex8"
)

// The token types used by the parser.
const (
	Keyword = iota
	Ident
	Int
	Float
	Char
	String
	Operator
	Semi
	Endl
)

// Types provides a type name querier
var Types = func() *lex8.Types {
	ret := lex8.NewTypes()
	o := func(t int, name string) {
		ret.Register(t, name)
	}

	o(Keyword, "keyword")
	o(Ident, "identifier")
	o(Int, "integer")
	o(Float, "float")
	o(Char, "char")
	o(String, "string")
	o(Operator, "operator")
	o(Semi, "semicolon")
	o(Endl, "endl")

	return ret
}()

// TypeStr returns the name of a token type.
func TypeStr(t int) string { return Types.Name(t) }
