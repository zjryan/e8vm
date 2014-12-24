package asm8

import (
	"lex8"
)

var insts = []instParse{
	parseInstReg,
	parseInstImm,
	parseInstBr,
	parseInstJmp,
	parseInstSys,
}

func parseInst(p *Parser, ops []*lex8.Token) (i *inst) {
	return instParsers(insts).parse(p, ops)
}
