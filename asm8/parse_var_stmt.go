package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func parseArgs(p *parser) (typ *lex8.Token, args []*lex8.Token) {
	typ = p.expect(Operand)
	if typ == nil {
		p.skipErrStmt()
		return nil, nil
	}

	for !p.acceptType(Semi) {
		if !p.hasErr() {
			t := p.token()
			if t.Type == Operand || t.Type == String {
				args = append(args, t)
			} else {
				p.err(t.Pos, "expect operand or string, got %s", typeStr(t.Type))
			}
		}
		if p.see(lex8.EOF) {
			break
		}
		p.next()
	}

	p.clearErr()

	return typ, args
}

func parseData(p *parser, t *lex8.Token, args []*lex8.Token) ([]byte, uint32) {
	switch t.Lit {
	case "str":
		return parseDataStr(p, args)
	case "u32", "i32", "u8", "i8", "byte", "f32":
		p.err(t.Pos, "data type %s not implemented", t.Lit)
		return nil, 0
	default:
		p.err(t.Pos, "unknown data type %q", t.Lit)
		return nil, 0
	}
}

func parseVarStmt(p *parser) *varStmt {
	typ, args := parseArgs(p)
	if typ == nil {
		return nil
	}

	ret := new(varStmt)
	ret.typ = typ
	ret.args = args
	ret.data, ret.align = parseData(p, typ, args)

	return ret
}
