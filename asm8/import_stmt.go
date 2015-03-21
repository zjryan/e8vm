package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/link8"
)

type importStmt struct {
	*ast.ImportStmt

	as   string
	path string
	lib  *link8.Pkg
	used bool
}
