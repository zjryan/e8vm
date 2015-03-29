package ir

import (
	"fmt"
	"io"
)

type printer struct {
	out io.Writer
	e   error
}

func (p *printer) printStr(s string) {
	if p.e != nil {
		return
	}
	_, e := fmt.Print(s)
	p.e = e
}

func (p *printer) Print(args ...interface{}) {
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			p.printStr(arg)
		case *stackVar:
			p.printStr(arg.name)
		case *number:
			p.printStr(fmt.Sprintf("%d", arg.v))
		default:
			panic(fmt.Errorf("invalid or unimplemented ref: %T", arg))
		}
	}
}

func (p *printer) Endline() { p.printStr("\n") }

func printOp(p *printer, op op) {
	switch op := op.(type) {
	case *arithOp:
		if op.a == nil {
			p.Print(op.dest, "=", op.op, op.b)
		} else {
			p.Print(op.dest, "=", op.a, op.op, op.b)
		}
		p.Endline()
	case *callOp:
		p.Print(op.dest, "=", op.f, "(")
		for _, arg := range op.args {
			p.Print(arg)
		}
		p.Print(")")
	default:
		panic(fmt.Errorf("invalid or unknown IR op: %T", op))
	}
}
