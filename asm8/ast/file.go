package ast

import (
	"lonnie.io/e8vm/lex8"
)

// File represents a file.
type File struct {
	Funcs []*FuncDecl
	Vars  []*VarDecl

	Comments []*lex8.Token
}
