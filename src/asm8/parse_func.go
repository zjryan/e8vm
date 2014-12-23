package asm8

func (f *Func) parseLines(p *Parser) {
	for !(p.see(Rbrace) || p.see(EOF)) {
		line := parseAsmLine(p)
		if line != nil {
			f.Lines = append(f.Lines, line.(*Line))
		}
	}
}

func parseFunc(p *Parser) interface{} {
	ret := new(Func)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)
	ret.lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	ret.parseLines(p)

	ret.rbrace = p.expect(Rbrace)
	ret.semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
