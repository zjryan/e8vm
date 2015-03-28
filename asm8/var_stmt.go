package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type varStmt struct {
	*ast.VarStmt

	align uint32
	data  []byte
}

func resolveVarStmt(log lex8.Logger, v *ast.VarStmt) *varStmt {
	ret := new(varStmt)
	ret.VarStmt = v
	ret.data, ret.align = resolveData(log, v.Type, v.Args)
	return ret
}

func resolveData(p lex8.Logger, t *lex8.Token, args []*lex8.Token) (
	[]byte, uint32,
) {
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
