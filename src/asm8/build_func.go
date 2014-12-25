package asm8

func buildFunc(b *Builder, f *Func) {
	b.scope.Push()

	declareLabels(b, f)
	setStmtOffset(b, f)
	fillStmtLabels(b, f)

	b.scope.Pop()
}

func declareLabels(b *Builder, f *Func) {
	for _, stmt := range f.stmts {
		if !stmt.isLabel() {
			continue
		}

		lab := stmt.label
		op := stmt.ops[0]
		sym := b.scope.Declare(&Symbol{
			Name: lab,
			Type: SymLabel,
			Item: stmt,
			Pos:  op.Pos,
		})

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

	for _, stmt := range f.stmts {
		stmt.offset = offset
		if stmt.isLabel() {
			continue
		}

		offset += 4
		offset += uint32(len(stmt.extras)) * 4
	}
}

func fillStmtLabels(b *Builder, f *Func) {
	for _, s := range f.stmts {
		if s.isLabel() {
			continue
		}
		if s.fill != fillLabel {
			continue
		}

		s.fill = fillNone

		if s.pack != "" {
			panic("fill label with pack symbol")
		}

		name := s.symbol
		op := s.ops[len(s.ops)-1] // last op must be the label
		if op.Lit != name {
			panic("picking the wrong op")
		}

		sym := b.scope.Query(s.symbol)
		if sym.Type != SymLabel {
			panic("not a label")
		}

		if sym == nil {
			b.err(op.Pos, "label %q not declared", name)
			continue
		}

		lab := sym.Item.(*stmt)
		delta := (lab.offset + 4 - s.offset) >> 2

		inst := &s.inst.inst
		if isJump(*inst) {
			*inst |= delta & 0x3fffffff
		} else {
			// it is a branch
			if !inBrRange(delta) {
				b.err(op.Pos, "label %q out of range of branching", name)
			}
			*inst |= delta & 0x3ffff
		}
	}
}
