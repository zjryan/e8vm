package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func buildFunc(b *Builder, f *Func) *funcObj {
	b.scope.Push()

	b.clearErr()

	declareLabels(b, f)

	if !b.hasError {
		setOffsets(b, f)
		fillLabels(b, f)
	}

	b.scope.Pop()

	return makeFuncObj(b, f)
}

// declareLabels adds the labels into the scope,
// so that later they can be queried for filling.
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

// setOffsets calculates the offset in function for each instruction.
func setOffsets(b *Builder, f *Func) {
	offset := uint32(0)

	for _, s := range f.stmts {
		s.offset = offset
		if s.isLabel() {
			continue
		}

		offset += 4
	}
}

// fillDelta will fill the particular instruction
// with a jumping/branching offset d.
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

// fillLabels fills all the labels in the function. After the filling, all the
// symbols left will be fillLink, fillHigh and fillLow and all branches (which
// must use labels) will be filled.
func fillLabels(b *Builder, f *Func) {
	for _, s := range f.stmts {
		if s.isLabel() {
			continue
		}
		if s.fill != fillLabel {
			continue
		}
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
	}
}

// makeFuncObj converts a function AST node f into a function object. It will
// resolve the symbols with fillLink, fillHigh and fillLow into <pack, sym>
// index pairs for the current package, using the symbol scope and the curPkg
// context in the builder.
func makeFuncObj(b *Builder, f *Func) *funcObj {
	ret := new(funcObj)

	if fillLabel != 4 {
		panic("bug")
	}

	for _, s := range f.stmts {
		if s.isLabel() {
			continue
		}

		ret.addInst(s.inst.inst)

		if !(s.fill > fillNone && s.fill < fillLabel) {
			continue
		}

		t := s.symTok
		if b.curPkg == nil {
			b.err(t.Pos, "no context for resolving %q", t.Lit)
			continue
		}

		var sym *Symbol
		if s.pack == "" {
			sym = b.scope.Query(s.symbol)
		} else {
			sym = b.scope.Query(s.pack)
			if sym == nil {
				b.err(t.Pos, "package %q not found", s.pack)
				continue
			} else if sym.Type != SymImport {
				t := s.symTok
				b.err(t.Pos, "%q is a %s, not a package",
					t.Lit, symStr(sym.Type))
				continue
			}

			pkg := sym.Item.(*Package)
			sym = pkg.symTable.Query(s.symbol)
		}

		if sym == nil {
			b.err(t.Pos, "%q not found", t.Lit)
			continue
		} else if sym.Type == SymImport || sym.Type == SymLabel {
			b.err(t.Pos, "%q is a %s, not a symbol",
				t.Lit, symStr(sym.Type))
			continue
		} else if s.fill == fillLink && sym.Type != SymFunc {
			b.err(t.Pos, "%q is a %s, not a function",
				t.Lit, symStr(sym.Type))
			continue
		}

		pkgIndex, pkgFound := b.curPkg.pkgIndex[sym.Package]
		if !pkgFound {
			panic("package not found in import")
		}
		pkg := b.curPkg.imports[pkgIndex]

		symIndex, symFound := pkg.symIndex[sym.Name]
		if !symFound {
			panic("symbol not found in package")
		}

		ret.addLink(s.fill, pkgIndex, symIndex)
	}

	return ret
}
