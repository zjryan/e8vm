package asm8

func parseAsmLine(p *Parser) interface{} {
	ret := new(Line)

	// a good assembly line is a series of ops that ends with
	// a semicolon or a right-brace
	for {
		if p.acceptType(Semi) || p.see(Rbrace) || p.see(EOF) {
			break
		}

		if p.see(Lbrace) {
			p.expect(Operand)
			return nil
		}

		t := p.expect(Operand)
		if t != nil {
			ret.Ops = append(ret.Ops, t)
		} else {
			ret = nil // error now
			if p.see(Lbrace) {
				break
			}
			p.next() // proceed anyway for other stuff
		}
	}

	p.clearErr()

	if ret == nil {
		return nil
	}

	return ret
}
