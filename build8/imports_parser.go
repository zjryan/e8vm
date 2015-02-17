package build8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

type importsParser struct {
	x    *importsLexer
	errs *lex8.ErrorList

	t     *lex8.Token
	inErr bool
}

func newImportsParser(file string, r io.ReadCloser) *importsParser {
	ret := new(importsParser)
	ret.x = newImportsLexer(file, r)
	ret.errs = lex8.NewErrList()
	ret.next()

	return ret
}

func (p *importsParser) see(t int) bool { return p.t.Type == t }
func (p *importsParser) clearErr()      { p.inErr = false }

func (p *importsParser) next() *lex8.Token {
	p.t = p.x.Token()
	return p.t
}

func (p *importsParser) err(pos *lex8.Pos, f string, args ...interface{}) {
	p.inErr = true
	p.errs.Addf(pos, f, args...)
}

func (p *importsParser) skipStmt() {
	for !(p.see(Semi) || p.see(lex8.EOF)) {
		p.next()
	}
	if p.see(Semi) {
		p.next()
	}
}

func (p *importsParser) skipErrStmt() bool {
	if !p.inErr {
		return false
	}

	p.skipStmt()
	p.clearErr()
	return true
}

func (p *importsParser) Errs() []*lex8.Error {
	ret := p.x.Errs()
	if ret != nil {
		return ret
	}
	return p.errs.Errs
}

func typeStr(t int) string {
	switch t {
	case lex8.EOF:
		return "eof"
	case lex8.Comment:
		return "comment"
	case Operand:
		return "operand"
	case Semi:
		return "';'"
	case Endl:
		return "end-line"
	case lex8.Illegal:
		return "illegal"
	}
	return "unknown"
}

func (p *importsParser) expect(t int) *lex8.Token {
	if p.inErr {
		return nil
	}

	if p.t.Type != t {
		p.err(p.t.Pos, "expect %s, got %s", typeStr(t), typeStr(p.t.Type))
		return nil
	}

	ret := p.t
	p.next()
	return ret
}
