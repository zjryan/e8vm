package ast

import (
	"lonnie.io/e8vm/lex8"
)

// File represents a file.
type File struct {
	Imports *ImportDecl

	Decls    []interface{}
	Comments []*lex8.Token
}
