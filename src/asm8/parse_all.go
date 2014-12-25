package asm8

func parseAsm8(p *Parser) {
	// TODO:
	if p.seeKeyword("func") {
		parseFunc(p)
	} else {
		p.err(p.t.Pos, "expect top-declaration")
	}
}
