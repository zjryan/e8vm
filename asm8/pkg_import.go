package asm8

import (
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// PkgImport states a package import clause
type PkgImport struct {
	As    string
	Pkg   *link8.Package
	Tok   *lex8.Token
	AsTok *lex8.Token

	use   int
	index int
}
