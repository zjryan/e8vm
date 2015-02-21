package build8

import (
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type pkgImport struct {
	path string
	as   string

	pathToken *lex8.Token
	asToken   *lex8.Token

	lib *link8.Pkg
}
