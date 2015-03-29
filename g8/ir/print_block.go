package ir

import (
	"fmt"
	"io"
)

func printBlock(p *printer, b *Block) {
	p.printStr(fmt.Sprintf("[block %d]", b.id))
	p.Endline()

	for _, op := range b.ops {
		printOp(p, op)
	}
}

func printFunc(p *printer, f *Func) {
	p.printStr(fmt.Sprintf("func %s", f.name))
	p.Endline()

	// printBlock(p, f.prologue)
	for _, b := range f.body {
		printBlock(p, b)
	}
	// printBlock(p, f.epilogue)
}

// PrintFunc prints a the content of a function
func PrintFunc(out io.Writer, f *Func) error {
	p := new(printer)
	p.out = out

	printFunc(p, f)
	return p.e
}
