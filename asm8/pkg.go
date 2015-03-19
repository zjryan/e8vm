package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	*ast.Pkg

	imports []*importPkg
	files   []*file
}

type importPkg struct {
	as  string
	pkg *link8.Pkg
}
