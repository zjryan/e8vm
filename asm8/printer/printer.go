package printer

import (
	"fmt"
	"io"
	"strings"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

const indentStr = "    "

type Printer struct {
	out  io.Writer
	last *lex8.Token
}

func (p *Printer) nextComment(pos *lex8.Pos) *lex8.Token {
	return nil
}

func isLabel(lit string) bool {
	return len(lit) > 0 && lit[0] == '.'
}

func (p *Printer) printComment(cm *lex8.Token, indent string) {
	lit := cm.Lit
	if strings.HasPrefix(lit, "//") {
		fmt.Fprintln(p.out, lit)
	} else {
		if !strings.HasPrefix(lit, "/*") {
			panic("invalid comment")
		}
	}
}

func (p *Printer) printFuncStmt(s *ast.FuncStmt) {
	ops := s.Ops
	if len(ops) == 0 {
		return
	}

	op0 := ops[0]

	lineStart := &lex8.Pos{
		File: op0.Pos.File,
		Line: op0.Pos.Line,
		Col:  0,
	}

	for {
		cm := p.nextComment(lineStart)
		if cm == nil {
			break
		}
		p.printComment(cm, indentStr)
		p.last = cm
	}

	if !isLabel(op0.Lit) {
		fmt.Fprint(p.out, indentStr)
	}

	first := true
	out := func(t *lex8.Token) {
		if !first {
			fmt.Fprint(p.out, " ")
		}
		first = false

		fmt.Fprint(p.out, t.Lit)
		p.last = t
	}

	for _, op := range s.Ops {
		for {
			cm := p.nextComment(op.Pos)
			if cm == nil {
				break
			}

			out(cm)
		}

		out(op)
	}
}
