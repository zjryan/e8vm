package build8

import (
	"io"
	"path"

	"lonnie.io/e8vm/lex8"
)

func parseImports(f string, r io.Reader) (*imports, []*lex8.Error) {
	p := newImportsParser(f, r)

	ret := newImports()

	for !p.See(lex8.EOF) {
		t := p.Expect(operand)
		if t == nil {
			p.SkipErrStmt(semi)
			continue
		}

		imp := new(pkgImport)
		imp.pathToken = t
		imp.path = t.Lit

		if p.See(operand) {
			t := p.Shift()
			imp.asToken = t
			imp.as = t.Lit
		}

		p.Expect(semi)
		if p.SkipErrStmt(semi) {
			continue
		}

		if !isPkgPath(imp.path) {
			p.Errorf(imp.pathToken.Pos,
				"invalid package path: %q",
				imp.path,
			)
			continue
		} else if imp.as != "" && !lex8.IsPkgName(imp.as) {
			p.Errorf(imp.pathToken.Pos,
				"invalid package alias: %q",
				imp.as,
			)
			continue
		}

		if imp.as == "" {
			imp.as = path.Base(imp.path)
		}
		if _, found := ret.m[imp.as]; found {
			p.Errorf(imp.pathToken.Pos, "duplicate import %q", imp.as)
			continue
		}

		ret.m[imp.as] = imp
	}

	es := p.Errs()
	if es != nil {
		return nil, es
	}
	return ret, nil
}
