package ast

import (
	"fmt"

	"lonnie.io/e8vm/fmt8"
)

func printVarDecl(p *fmt8.Printer, d *VarDecl) {
	ss := make([]string, len(d.Idents.Idents))
	for i, id := range d.Idents.Idents {
		ss[i] = id.Lit
	}

	fmt.Fprint(p, fmt8.Join(ss, ","))
	if d.Type != nil {
		fmt.Fprint(p, " ")
		printExpr(p, d.Type)
	}

	if d.Eq != nil {
		fmt.Fprint(p, " = ")
		printExpr(p, d.Exprs)
	}
}

func printVarDecls(p *fmt8.Printer, d *VarDecls) {
	if d.Lparen == nil {
		// single declare
		fmt.Fprintf(p, "var ")
		for _, decl := range d.Decls {
			printVarDecl(p, decl)
		}
	} else {
		fmt.Fprintf(p, "var (\n")
		p.Tab()
		for _, decl := range d.Decls {
			printVarDecl(p, decl)
			fmt.Println(p)
		}
		p.ShiftTab()
		fmt.Fprintf(p, ")")
	}
}

func printConstDecls(p *fmt8.Printer, d *ConstDecls) {
	fmt.Fprintf(p, "<todo: const decls>")
}
