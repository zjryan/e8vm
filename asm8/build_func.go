package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// buildFunc builds a function object from a function AST node.
func buildFunc(b *builder, f *ast.FuncDecl) *link8.Func {
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
func declareLabels(b *builder, f *ast.FuncDecl) {
	for _, stmt := range f.Stmts {
		if !stmt.IsLabel() {
			continue
		}

		lab := stmt.Label
		op := stmt.Ops[0]
		sym := &symbol{
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
func setOffsets(b *builder, f *ast.FuncDecl) {
	offset := uint32(0)

	for _, s := range f.Stmts {
		s.Offset = offset
		if s.IsLabel() {
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
			b.err(t.Pos, "%q is out of branch range", t.Lit)
		}
		*inst |= d & 0x3ffff
	}
}

// fillLabels fills all the labels in the function. After the filling, all the
// symbols left will be fillLink, fillHigh and fillLow and all branches (which
// must use labels) will be filled.
func fillLabels(b *builder, f *ast.FuncDecl) {
	for _, s := range f.Stmts {
		if s.IsLabel() {
			continue
		}
		if s.Fill != ast.FillLabel {
			continue
		}
		if s.Pkg != "" {
			panic("fill label with pack symbol")
		}

		t := s.SymTok

		sym := b.scope.Query(s.Sym)
		if sym == nil {
			b.err(t.Pos, "label %q not declared", t.Lit)
			continue
		}

		if sym.Type != SymLabel {
			panic("not a label")
		}

		lab := sym.Item.(*ast.FuncStmt)
		delta := uint32(int32(lab.Offset-s.Offset-4) >> 2)
		fillDelta(b, t, &s.Inst.Inst, delta)
	}
}

func queryPkg(b *builder, t *lex8.Token, pack string) *ast.PkgImport {
	sym := b.scope.Query(pack)
	if sym == nil {
		b.err(t.Pos, "package %q not found", pack)
		return nil
	} else if sym.Type != SymImport {
		b.err(t.Pos, "%q is a %s, not a package", t.Lit, symStr(sym.Type))
		return nil
	}
	return sym.Item.(*ast.PkgImport)
}

func init() {
	as := func(b bool) {
		if !b {
			panic("bug")
		}
	}
	as(ast.FillNone == 0 && ast.FillLabel == 4)
	as(ast.FillLink == link8.FillLink)
	as(ast.FillHigh == link8.FillHigh)
	as(ast.FillLow == link8.FillLow)
}

// resolveSymbol resolves the symbol in the statement,
// returns the symbol linking object and its <sym, pkg> index pair
// in the current package context.
func resolveSymbol(b *builder, s *ast.FuncStmt) (typ int, pkg, index uint32) {
	t := s.SymTok

	if s.Pkg == "" {
		sym := b.scope.Query(s.Sym) // find the symbol in scope
		if sym != nil {
			var p *link8.Package
			p, pkg = b.curPkg.PkgIndex(sym.Package)
			index = p.SymIndex(sym.Name)
			typ = sym.Type
		}
	} else {
		p := queryPkg(b, t, s.Pkg) // find the package
		if p != nil {
			var sym *link8.Symbol
			pkg = p.Index
			p.Use = true
			sym, index = p.Pkg.Query(s.Sym)
			if sym != nil {
				// should we use a consistant
				if sym.Type == link8.SymFunc {
					typ = SymFunc
				} else if sym.Type == link8.SymVar {
					typ = SymVar
				} else {
					b.err(t.Pos, "%q is an invalid linking symbol", t.Lit)
					return
				}
			}
		}
	}

	if typ == SymNone {
		b.err(t.Pos, "%q not found", t.Lit)
	} else if typ == SymConst {
		b.err(t.Pos, "const symbol filling not implemented yet")
	} else if typ == SymImport || typ == SymLabel {
		b.err(t.Pos, "cannot link %s %q", symStr(typ), t.Lit)
	}

	return
}

func linkSymbol(b *builder, s *ast.FuncStmt, f *link8.Func) {
	t := s.SymTok
	if b.curPkg == nil {
		b.err(t.Pos, "no context for resolving %q", t.Lit)
		return // this may happen for bare function
	}

	// TODO: this now looks for the symbol
	typ, pkg, index := resolveSymbol(b, s)
	if typ == SymNone {
		return
	}

	if s.Fill == ast.FillLink && typ != SymFunc {
		b.err(t.Pos, "%s %q is not a function", symStr(typ), t.Lit)
		return
	} else if pkg > 0 && !isPublic(s.Sym) {
		// for imported package, check if it is public
		b.err(t.Pos, "%q is not public", t.Lit)
		return
	}

	// save the link
	f.AddLink(s.Fill, pkg, index)
}

// makeFuncObj converts a function AST node f into a function object. It
// resolves the symbols of fillLink, fillHigh and fillLow into <pack, sym>
// index pairs, using the symbol scope and the curPkg context in the Builder b.
func makeFuncObj(b *builder, f *ast.FuncDecl) *link8.Func {
	ret := link8.NewFunc()
	for _, s := range f.Stmts {
		if s.IsLabel() {
			continue // skip labels
		}
		ret.AddInst(s.Inst.Inst)

		if !(s.Fill > ast.FillNone && s.Fill < ast.FillLabel) {
			continue // only care about fillHigh, fillLow and fillLink
		}

		linkSymbol(b, s, ret)
	}

	if ret.TooLarge() {
		b.err(f.Name.Pos, "too many instructions in func", f.Name.Lit)
	}

	return ret
}
