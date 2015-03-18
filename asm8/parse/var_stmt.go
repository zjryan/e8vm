package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseArgs(p *parser) (typ *lex8.Token, args []*lex8.Token) {
	typ = p.Expect(Operand)
	if typ == nil {
		p.skipErrStmt()
		return nil, nil
	}

	for !p.Accept(Semi) {
		if !p.InError() {
			t := p.Token()
			if t.Type == Operand || t.Type == String {
				args = append(args, t)
			} else {
				p.Errorf(t.Pos, "expect operand or string, got %s", p.TypeStr(t.Type))
			}
		}
		if p.See(lex8.EOF) {
			break
		}
		p.Next()
	}

	p.BailOut()

	return typ, args
}

func parseData(p *parser, t *lex8.Token, args []*lex8.Token) ([]byte, uint32) {
	switch t.Lit {
	case "str":
		return parseDataStr(p, args)
	case "x":
		return parseDataHex(p, args)
	case "u32":
		return parseDataNums(p, args, modeWord)
	case "i32":
		return parseDataNums(p, args, modeWord|modeSigned)
	case "u8", "byte":
		return parseDataNums(p, args, 0)
	case "i8":
		return parseDataNums(p, args, modeSigned)
	case "f32":
		return parseDataNums(p, args, modeWord|modeFloat)
	default:
		p.Errorf(t.Pos, "unknown data type %q", t.Lit)
		return nil, 0
	}
}

func parseVarStmt(p *parser) *ast.VarStmt {
	typ, args := parseArgs(p)
	if typ == nil {
		return nil
	}

	ret := new(ast.VarStmt)
	ret.Type = typ
	ret.Args = args
	ret.Data, ret.Align = parseData(p, typ, args)

	return ret
}
