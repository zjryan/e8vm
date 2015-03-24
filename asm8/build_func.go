package asm8

import (
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/sym8"
)

// buildFunc builds a function object from a function AST node.
func buildFunc(b *builder, f *funcDecl) *link8.Func {
	b.scope.Push()
	defer b.scope.Pop()

	b.BailOut()
	declareLabels(b, f)
	if !b.InJail() {
		setOffsets(b, f)
		fillLabels(b, f)
	}

	return makeFuncObj(b, f)
}

// declareLabels adds the labels into the scope,
// so that later they can be queried for filling.
func declareLabels(b *builder, f *funcDecl) {
	for _, stmt := range f.stmts {
		if !stmt.isLabel() {
			continue
		}

		lab := stmt.label
		op := stmt.Ops[0]
		sym := sym8.Make(lab, SymLabel, stmt, op.Pos)
		decl := b.scope.Declare(sym)
		if decl != nil {
			b.Errorf(op.Pos, "%q already declared", lab)
			b.Errorf(decl.Pos, "  here as a %s", symStr(decl.Type))
			continue
		}
	}
}

// setOffsets calculates the offset in function for each instruction.
func setOffsets(b *builder, f *funcDecl) {
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
func fillDelta(b *builder, t *lex8.Token, inst *uint32, d uint32) {
	if isJump(*inst) {
		*inst |= d & 0x3fffffff
	} else {
		// it is a branch
		if !inBrRange(d) {
			b.Errorf(t.Pos, "%q is out of branch range", t.Lit)
		}
		*inst |= d & 0x3ffff
	}
}

// fillLabels fills all the labels in the function. After the filling, all the
// symbols left will be fillLink, fillHigh and fillLow and all branches (which
// must use labels) will be filled.
func fillLabels(b *builder, f *funcDecl) {
	for _, s := range f.stmts {
		if s.isLabel() {
			continue
		}
		if s.fill != fillLabel {
			continue
		}
		if s.pkg != "" {
			panic("fill label with pack symbol")
		}

		t := s.symTok

		sym := b.scope.Query(s.sym)
		if sym == nil {
			b.Errorf(t.Pos, "label %q not declared", t.Lit)
			continue
		}

		if sym.Type != SymLabel {
			panic("not a label")
		}

		lab := sym.Item.(*funcStmt)
		delta := uint32(int32(lab.offset-s.offset-4) >> 2)
		fillDelta(b, t, &s.inst.inst, delta)
	}
}

func queryPkg(b *builder, t *lex8.Token, pack string) *importStmt {
	sym := b.scope.Query(pack)
	if sym == nil {
		b.Errorf(t.Pos, "package %q not found", pack)
		return nil
	} else if sym.Type != SymImport {
		b.Errorf(t.Pos, "%q is a %s, not a package", t.Lit, symStr(sym.Type))
		return nil
	}
	return sym.Item.(*importStmt)
}

func init() {
	as := func(b bool) {
		if !b {
			panic("bug")
		}
	}
	as(fillNone == 0 && fillLabel == 4)
	as(fillLink == link8.FillLink)
	as(fillHigh == link8.FillHigh)
	as(fillLow == link8.FillLow)
}

// resolveSymbol resolves the symbol in the statement,
// returns the symbol linking object and its <sym, pkg> index pair
// in the current package context.
//
// this function only resolves symbol that requires linking
// which are variables and functions
func resolveSymbol(b *builder, s *funcStmt) (typ int, pkg, index uint32) {
	t := s.symTok

	// TODO: this code part is too messy, need to clean this.
	if s.pkg == "" {
		sym := b.scope.Query(s.sym) // find the symbol in scope
		if sym != nil {
			typ = sym.Type
			if typ == SymVar || typ == SymFunc {
				index = b.curPkg.SymIndex(sym.Name)
			}
		}
	} else {
		p := queryPkg(b, t, s.pkg) // find the package importStmt
		if p != nil {
			var sym *link8.Symbol // for saving the linking symbol in the lib

			pkg = b.getIndex(p.as)       // package index in lib, based on alias
			b.pkgUsed[p.as] = struct{}{} // mark pkg used

			// TODO: we should find this in back in linkable when possible
			// this is required for handling consts
			sym, index = p.lib.Query(s.sym) // find the symbol in the package
			if sym != nil {
				if sym.Type == link8.SymFunc {
					typ = SymFunc
				} else if sym.Type == link8.SymVar {
					typ = SymVar
				} else {
					b.Errorf(t.Pos, "%q is an invalid linking symbol", t.Lit)
					return
				}
			}
		}
	}

	if typ == SymNone {
		b.Errorf(t.Pos, "%q not found", t.Lit)
	} else if typ == SymConst {
		b.Errorf(t.Pos, "const symbol filling not implemented yet")
		typ = SymNone // report as error
	} else if typ == SymImport || typ == SymLabel {
		b.Errorf(t.Pos, "cannot link %s %q", symStr(typ), t.Lit)
		typ = SymNone // report as error
	}

	return
}

func linkSymbol(b *builder, s *funcStmt, f *link8.Func) {
	t := s.symTok
	if b.curPkg == nil {
		b.Errorf(t.Pos, "no context for resolving %q", t.Lit)
		return // this may happen for bare function
	}

	typ, pkg, index := resolveSymbol(b, s)
	if typ == SymNone {
		return
	}

	if s.fill == fillLink && typ != SymFunc {
		b.Errorf(t.Pos, "%s %q is not a function", symStr(typ), t.Lit)
		return
	} else if pkg > 0 && !isPublic(s.sym) {
		// for imported package, check if it is public
		b.Errorf(t.Pos, "%q is not public", t.Lit)
		return
	}

	// save the link
	f.AddLink(s.fill, pkg, index)
}

// makeFuncObj converts a function AST node f into a function object. It
// resolves the symbols of fillLink, fillHigh and fillLow into <pack, sym>
// index pairs, using the symbol scope and the curPkg context in the Builder b.
func makeFuncObj(b *builder, f *funcDecl) *link8.Func {
	ret := link8.NewFunc()
	for _, s := range f.stmts {
		if s.isLabel() {
			continue // skip labels
		}
		ret.AddInst(s.inst.inst)

		if !(s.fill > fillNone && s.fill < fillLabel) {
			continue // only care about fillHigh, fillLow and fillLink
		}

		linkSymbol(b, s, ret)
	}

	if ret.TooLarge() {
		b.Errorf(f.Name.Pos, "too many instructions in func", f.Name.Lit)
	}

	return ret
}
