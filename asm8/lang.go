package asm8

import (
	"strings"

	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
)

type lang struct{}

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".s")
}

func (lang) Import(pkg build8.Pkg) []*lex8.Error {
	src := pkg.Src()

	if len(src) == 1 {
		for _, f := range src {
			return listImport(f.Path, f, pkg)
		}
	}

	f := src["import.s"]
	if f == nil {
		return nil
	}
	return listImport(f.Path, f, pkg)
}

func (lang) Compile(p build8.Pkg) []*lex8.Error {
	// resolve pass, will also parse the files
	pkg, es := resolvePkg(p.Path(), p.Src())
	if es != nil {
		return es
	}

	imports := p.Imports()
	// import
	errs := lex8.NewErrorList()
	if pkg.imports != nil {
		for _, stmt := range pkg.imports.stmts {
			imp := imports[stmt.as]
			if imp == nil || imp.Pkg == nil {
				errs.Errorf(stmt.Path.Pos, "import missing")
				continue
			}

			stmt.linkable = imp.Pkg.Compiled()
			if stmt.linkable == nil {
				panic("import missing")
			}

			stmt.lib = stmt.linkable.Lib()
		}

		if es := errs.Errs(); es != nil {
			return es
		}
	}

	// library building
	b := newBuilder()
	lib := buildLib(b, pkg)
	if es := b.Errs(); es != nil {
		return es
	}

	p.SetCompiled(lib)
	return nil
}

// Lang returns the assembly language builder for the building system
func Lang() build8.Lang { return lang{} }
