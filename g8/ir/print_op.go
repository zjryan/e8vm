package ir

import (
	"fmt"

	"lonnie.io/e8vm/fmt8"
)

func printOp(p *fmt8.Printer, op op) {
	switch op := op.(type) {
	case *arithOp:
		if op.a == nil {
			if op.op == "" {
				fmt.Fprintf(p, "%s = %s\n", op.dest, op.b)
			} else if op.op == "0" {
				fmt.Fprintf(p, "%s = 0\n", op.dest)
			} else {
				fmt.Fprintf(p, "%s = %s %s\n", op.dest, op.op, op.b)
			}
		} else {
			fmt.Fprintf(p, "%s = %s %s %s\n",
				op.dest, op.a, op.op, op.b,
			)
		}
	case *callOp:
		args := fmt8.Join(op.args, ",")
		fmt.Fprintf(p, "%s = %s(%s)\n", op.dest, op.f, args)
	default:
		panic(fmt.Errorf("invalid or unknown IR op: %T", op))
	}
}

func printJump(p *fmt8.Printer, j *blockJump) {
	if j == nil {
		return
	}

	switch j.typ {
	case jmpAlways:
		fmt.Fprintf(p, "goto %s\n", j.to)
	case jmpIf:
		fmt.Fprintf(p, "if %s goto %s\n", j.cond, j.to)
	case jmpIfNot:
		fmt.Fprintf(p, "if !%s goto %s\n", j.cond, j.to)
	default:
		panic("invalid jump type")
	}
}
