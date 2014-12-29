package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// buildFunc builds a function object from a function AST node.
func buildFunc(b *Builder, f *Func) *funcObj {
	b.scope.Push()
	defer b.scope.Pop()

	b.clearErr()
	declareLabels(b, f)
	if !b.hasError {
		setOffsets(b, f)
		fillLabels(b, f)
	}

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

func queryPkg(b *Builder, t *lex8.Token, pack string) *pkgObj {
	sym := b.scope.Query(pack)
	if sym == nil {
		b.err(t.Pos, "package %q not found", pack)
		return nil
	} else if sym.Type != SymImport {
		b.err(t.Pos, "%q is a %s, not a package", t.Lit, symStr(sym.Type))
		return nil
	}
	return sym.Item.(*pkgObj)
}

func init() {
	if fillLabel != 4 || fillNone != 0 {
		panic("bug")
	}
}

// resolveSymbol resolves a symbol, returns the symbol object
// and its <sym, pkg> index pair in the current package context.
func resolveSymbol(b *Builder, s *stmt) (ret *Symbol, pkg, index uint32) {
	t := s.symTok

	if s.pack == "" {
		ret = b.scope.Query(s.symbol) // find the symbol in scope
		if ret != nil {
			pkg = b.curPkg.PkgIndex(ret.Package)
			index = b.curPkg.requires[pkg].SymIndex(ret.Name)
		}
	} else {
		p := queryPkg(b, t, s.pack) // find the package
		if p == nil {
			return
		}

		pkg = b.curPkg.PkgIndex(p.path)
		ret, index = p.Query(s.symbol)
		if ret != nil && ret.Package != p.path {
			panic("bug")
		}
	}

	if ret == nil {
		b.err(t.Pos, "%q not found", t.Lit)
		return nil, 0, 0
	} else if ret.Type == SymImport || ret.Type == SymLabel {
		b.err(t.Pos, "%s %q not a symbol", symStr(ret.Type), t.Lit)
		return nil, 0, 0
	}

	return
}

func linkSymbol(b *Builder, s *stmt, f *funcObj) {
	t := s.symTok
	if b.curPkg == nil {
		b.err(t.Pos, "no context for resolving %q", t.Lit)
		return // this may happen for bare function
	}

	sym, pkg, index := resolveSymbol(b, s)
	if sym == nil {
		return
	}

	if s.fill == fillLink && sym.Type != SymFunc {
		b.err(t.Pos, "%s %q is not a function", symStr(sym.Type), t.Lit)
		return
	} else if pkg > 0 && !IsPublic(sym.Name) {
		// for imported package, check if it is public
		b.err(t.Pos, "%q is not public", t.Lit)
		return
	}

	// save the link
	f.addLink(s.fill, pkg, index)
}

// makeFuncObj converts a function AST node f into a function object. It
// resolves the symbols of fillLink, fillHigh and fillLow into <pack, sym>
// index pairs, using the symbol scope and the curPkg context in the Builder b.
func makeFuncObj(b *Builder, f *Func) *funcObj {
	ret := new(funcObj)
	for _, s := range f.stmts {
		if s.isLabel() {
			continue // skip labels
		}
		ret.addInst(s.inst.inst)

		if !(s.fill > fillNone && s.fill < fillLabel) {
			continue // only care about fillHigh, fillLow and fillLink
		}

		linkSymbol(b, s, ret)
	}

	if ret.TooLarge() {
		b.err(f.name.Pos, "too many instructions in func", f.name.Lit)
	}

	return ret
}
