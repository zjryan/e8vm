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

// a listing of possible declarations
var decls = []interface{}{
	new(Func),
	new(Var),
}
