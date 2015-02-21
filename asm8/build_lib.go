package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
)

func declareSymbol(b *builder, sym *symbol) bool {
	// declare in the scope
	exists := b.scope.Declare(sym)
	if exists != nil {
		b.err(sym.Pos, "%q already declared", sym.Name)
		b.err(exists.Pos, "  previously declared here")
		return false
	}
	return true
}

func declareFile(b *builder, pkg *ast.Pkg, file *ast.File) {
	// declare functions
	for _, fn := range file.Funcs {
		t := fn.Name
		sym := &symbol{
			t.Lit,
			SymFunc,
			fn,
			t.Pos,
			pkg.Path,
		}

		if !declareSymbol(b, sym) {
			continue
		}

		// declare in the lib
		fn.Index = b.curPkg.Declare(sym)
	}

	// declare variables
	for _, v := range file.Vars {
		t := v.Name
		sym := &symbol{
			t.Lit,
			SymVar,
			v,
			t.Pos,
			pkg.Path,
		}

		if !declareSymbol(b, sym) {
			continue
		}

		// declare in the lib
		v.Index = b.curPkg.Declare(sym)
	}
}

func buildPkgScope(b *builder, pkg *ast.Pkg) {
	// declare requires
	for as, p := range pkg.Imports {
		// p is the *PkgImport
		sym := &symbol{
			as,
			SymImport,
			p,
			p.Tok.Pos,
			p.Pkg.Path(),
		}
		if !declareSymbol(b, sym) {
			continue
		}

		p.Index = b.curPkg.Require(p.Pkg)
	}

	for _, file := range pkg.Files {
		declareFile(b, pkg, file)
	}
}

func checkUnusedImport(b *builder, pkg *ast.Pkg) {
	for _, imp := range pkg.Imports {
		if !imp.Use {
			b.err(imp.Tok.Pos, "package %q imported but not used", imp.As)
		}
	}
}

func buildLib(b *builder, pkg *ast.Pkg) *lib {
	ret := newLib(pkg.Path)
	b.curPkg = ret

	b.scope.Push()
	defer b.scope.Pop()

	buildPkgScope(b, pkg)
	if b.Errs() != nil {
		return nil // error on declaring, so just return
	}

	for _, file := range pkg.Files {
		buildFile(b, file)
	}

	checkUnusedImport(b, pkg)

	if b.Errs() != nil {
		return nil // error on building
	}

	return ret
}
