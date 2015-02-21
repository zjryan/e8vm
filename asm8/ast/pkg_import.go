package ast

import (
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// PkgImport states a package import clause
type PkgImport struct {
	As    string
	Tok   *lex8.Token
	AsTok *lex8.Token

	Pkg *link8.Package
}
