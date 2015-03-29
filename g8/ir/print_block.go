package ir

import (
	"fmt"
	"io"

	"lonnie.io/e8vm/fmt8"
)

func printBlock(p *fmt8.Printer, b *Block) {
	fmt.Fprintln(p, b)
	p.Tab()
	for _, op := range b.ops {
		printOp(p, op)
	}
	p.ShiftTab()
}

func printFunc(p *fmt8.Printer, f *Func) {
	fmt.Fprintf(p, "func %s {\n", f.name)
	p.Tab()

	for _, b := range f.body {
		printBlock(p, b)
	}

	p.ShiftTab()
	fmt.Fprintln(p, "}")
}

func printPkg(p *fmt8.Printer, pkg *Pkg) {
	fmt.Fprintf(p, "package %s\n", pkg.path)

	for _, f := range pkg.funcs {
		fmt.Fprintln(p)
		printFunc(p, f)
	}
}

// PrintFunc prints a the content of a function
func PrintPkg(out io.Writer, pkg *Pkg) error {
	p := fmt8.NewPrinter(out)
	printPkg(p, pkg)
	return p.Err()
}
