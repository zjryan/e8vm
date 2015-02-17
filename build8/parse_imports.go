package build8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

func parseImports(f string, rc io.ReadCloser) ([]*pkgImport, []*lex8.Error) {
	p := newImportsParser(f, rc)

	var ret []*pkgImport

	for !p.see(lex8.EOF) {
		if !p.see(Operand) {
			p.err(p.t.Pos, "expect operand")
			p.skipErrStmt()
			continue
		}

		imp := new(pkgImport)

		imp.path = p.t.Lit
		p.next()

		if p.see(Operand) {
			imp.as = p.t.Lit
			p.next()
		}

		p.expect(Semi)
		p.skipErrStmt()

		ret = append(ret, imp)
	}

	es := p.Errs()
	if es != nil {
		return nil, es
	}
	return ret, nil
}
