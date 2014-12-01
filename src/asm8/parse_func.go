package asm8

// Line is an assembly line.
type Line struct {
	toks []*Token
}

// Func is an assembly function.
type Func struct {
	lines []*Line

	kw     *Token
	name   *Token
	lbrace *Token
	rbrace *Token
}

func parseFunc(p *Parser) *Func {
	ret := new(Func)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)
	ret.lbrace = p.expect(Lbrace)
	p.expect(Endl)

	// parse lines here

	ret.rbrace = p.expect(Rbrace)
	p.expect(Endl)

	p.inErr = false

	return ret
}
