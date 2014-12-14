package asm8

func (f *Func) parseLines(p *Parser) {
	for !p.see(Rbrace) {
		var ops []*Token
		for !p.see(Endl) {
			t := p.expect(Operand)
			if t != nil {
				ops = append(ops, t)
			} else if p.see(EOF) {
				return
			} else {
				p.next()
			}
		}

		if !p.skipErrLine() {
			p.expect(Endl)
			if ops != nil {
				f.Lines = append(f.Lines, &Line{ops})
			}
		}
	}
}

func parseFunc(p *Parser) interface{} {
	ret := new(Func)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)
	ret.lbrace = p.expect(Lbrace)
	p.expect(Endl)
	if p.skipErrLine() {
		return ret
	}

	ret.parseLines(p)

	ret.rbrace = p.expect(Rbrace)
	if p.skipErrLine() {
		return ret
	}

	return ret
}
