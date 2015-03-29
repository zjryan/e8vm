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
				fmt.Fprintf(p, "%s = %s\n",
					op.dest, op.b,
				)
			} else {
				fmt.Fprintf(p, "%s = %s %s\n",
					op.dest, op.op, op.b,
				)
			}
		} else {
			fmt.Fprintf(p, "%s = %s %s %s\n",
				op.dest, op.a, op.op, op.b,
			)
		}
	case *callOp:
		fmt.Fprintf(p, "%s = %s(", op.dest, op.f)
		for i, arg := range op.args {
			if i > 0 {
				fmt.Fprint(p, ",")
			}
			fmt.Fprint(p, arg)
		}
		fmt.Fprint(p, ")\n")
	default:
		panic(fmt.Errorf("invalid or unknown IR op: %T", op))
	}
}
