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

func (lang) Prepare(
	src map[string]*build8.File, imp build8.Importer,
) []*lex8.Error {
	if len(src) == 1 {
		for _, f := range src {
			return listImport(f.Path, f, imp)
		}
	}

	f := src["import.s"]
	if f == nil {
		return nil
	}
	return listImport(f.Path, f, imp)
}

func (lang) Compile(pinfo *build8.PkgInfo) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	// resolve pass, will also parse the files
	pkg, es := resolvePkg(pinfo.Path, pinfo.Src)
	if es != nil {
		return nil, es
	}

	// import
	errs := lex8.NewErrorList()
	if pkg.imports != nil {
		for _, stmt := range pkg.imports.stmts {
			imp := pinfo.Import[stmt.as]
			if imp == nil || imp.Compiled == nil {
				errs.Errorf(stmt.Path.Pos, "import missing")
				continue
			}

			stmt.linkable = imp.Compiled
			if stmt.linkable == nil {
				panic("import missing")
			}

			stmt.lib = stmt.linkable.Lib()
		}

		if es := errs.Errs(); es != nil {
			return nil, es
		}
	}

	// library building
	b := newBuilder()
	lib := buildLib(b, pkg)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	return lib, nil
}

// Lang returns the assembly language builder for the building system
func Lang() build8.Lang { return lang{} }
