package asm8

func parseAsm8(p *Parser) {
	if p.seeKeyword("func") {
		parseFunc(p)
	}

	p.err(p.t.Pos, "expect top-declaration")
}
