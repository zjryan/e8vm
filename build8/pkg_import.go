package build8

import (
	"lonnie.io/e8vm/lex8"
)

type pkgImport struct {
	path string
	as   string

	pathToken *lex8.Token
	asToken   *lex8.Token

	used bool
}
