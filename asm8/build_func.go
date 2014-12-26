package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func buildFunc(b *Builder, f *Func) {
	b.scope.Push()

	b.clearErr()

	declareLabels(b, f)

	if !b.hasError {
		setStmtOffset(b, f)
		fillStmtLabels(b, f)
	}

	b.scope.Pop()
}

func declareLabels(b *Builder, f *Func) {
	for _, stmt := range f.stmts {
		if !stmt.isLabel() {
			continue
		}

		lab := stmt.label
		op := stmt.ops[0]
		sym := &Symbol{
			Name: lab,
			Type: SymLabel,
			Item: stmt,
			Pos:  op.Pos,
		}

		decl := b.scope.Declare(sym)
		if decl != nil {
			b.err(op.Pos, "%q already declared", lab)
			b.err(decl.Pos, "  here as a %s", symStr(decl.Type))
			continue
		}
	}
}

func setStmtOffset(b *Builder, f *Func) {
	offset := uint32(0)

	for _, s := range f.stmts {
		s.offset = offset
		if s.isLabel() {
			continue
		}

		offset += 4
		offset += uint32(len(s.extras)) * 4
	}
}

func fillDelta(b *Builder, t *lex8.Token, inst *uint32, d uint32) {
	if isJump(*inst) {
		*inst |= d & 0x3fffffff
	} else {
		// it is a branch
		if !inBrRange(d) {
			b.err(t.Pos, "%q is out of branch range", t.Lit)
		}
		*inst |= d & 0x3ffff
	}
}

func fillStmtLabels(b *Builder, f *Func) {
	for _, s := range f.stmts {
		if s.isLabel() {
			continue
		}

		switch s.fill {
		case fillNone:
			continue // need to do nothing
		case fillLabel:
			if s.pack != "" {
				panic("fill label with pack symbol")
			}

			t := s.symTok

			sym := b.scope.Query(s.symbol)
			if sym == nil {
				b.err(t.Pos, "label %q not declared", t.Lit)
				continue
			}

			if sym.Type != SymLabel {
				panic("not a label")
			}

			lab := sym.Item.(*stmt)
			delta := uint32(int32(lab.offset-s.offset-4) >> 2)
			fillDelta(b, t, &s.inst.inst, delta)
		default:
			b.err(s.ops[0].Pos, "not implemented filling: %d", s.fill)
		}
	}
}
