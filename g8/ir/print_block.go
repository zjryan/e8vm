package ir

import (
	"fmt"
	"io"

	"lonnie.io/e8vm/fmt8"
)

func printBlock(p *fmt8.Printer, b *Block) {
	fmt.Fprintf(p, "%s:\n", b)
	p.Tab()
	for _, op := range b.ops {
		printOp(p, op)
	}
	printJump(p, b.jump)
	p.ShiftTab()
}

func printFunc(p *fmt8.Printer, f *Func) {
	fmt.Fprintf(p, "func %s {\n", f.name)
	p.Tab()

	for b := f.prologue.next; b != f.epilogue; b = b.next {
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

// PrintPkg prints a the content of a IR package
func PrintPkg(out io.Writer, pkg *Pkg) error {
	p := fmt8.NewPrinter(out)
	printPkg(p, pkg)
	return p.Err()
}
