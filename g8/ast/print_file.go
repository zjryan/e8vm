package ast

import (
	"fmt"

	"lonnie.io/e8vm/fmt8"
)

func printTopDecl(p *fmt8.Printer, d Decl) {
	switch d := d.(type) {
	case *Func:
		printFunc(p, d)
	case *Struct:
		printStruct(p, d)
	case *VarDecls:
		printVarDecls(p, d)
	case *ConstDecls:
		printConstDecls(p, d)
	default:
		fmt.Fprintf(p, "<!!%T>", d)
	}
}

func printFile(p *fmt8.Printer, f *File) {
	for _, decl := range f.Decls {
		printTopDecl(p, decl)
		fmt.Fprintf(p, "\n")
	}
}
