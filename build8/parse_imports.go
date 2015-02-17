package build8

import (
	"io"
	"path"

	"lonnie.io/e8vm/lex8"
)

func parseImports(f string, rc io.ReadCloser) (*imports, []*lex8.Error) {
	p := newImportsParser(f, rc)

	ret := newImports()

	for !p.see(lex8.EOF) {
		if !p.see(operand) {
			p.err(p.t.Pos, "expect operand")
			p.skipErrStmt()
			continue
		}

		imp := new(pkgImport)

		imp.pathToken = p.t
		imp.path = p.t.Lit
		p.next()

		if p.see(operand) {
			imp.asToken = p.t
			imp.as = p.t.Lit
			p.next()
		}

		p.expect(semi)

		if p.hasErr() {
			p.skipErrStmt()
			continue
		}

		if !isPkgPath(imp.path) {
			p.err(imp.pathToken.Pos,
				"invalid package path: %q",
				imp.path,
			)
			continue
		} else if imp.as != "" && !isPkgName(imp.as) {
			p.err(imp.pathToken.Pos,
				"invalid package alias: %q",
				imp.as,
			)
			continue
		}

		if imp.as == "" {
			imp.as = path.Base(imp.path)
		}
		if _, found := ret.m[imp.as]; found {
			p.err(imp.pathToken.Pos, "duplicate import %q", imp.as)
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
