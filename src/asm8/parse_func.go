package asm8

type Line struct {
	toks []*Token
}

type Func struct {
	lines []*Line

	kw     *Token
	name   *Token
	lbrace *Token
	rbrace *Token
}

func parseFunc(p *Parser) interface{} {
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
