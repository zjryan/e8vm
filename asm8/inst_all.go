package asm8

import (
	"lonnie.io/e8vm/lex8"
)

var insts = []instParse{
	parseInstReg,
	parseInstImm,
	parseInstBr,
	parseInstJmp,
	parseInstSys,
}

func parseInst(p *parser, ops []*lex8.Token) (i *inst) {
	return instParsers(insts).parse(p, ops)
}
