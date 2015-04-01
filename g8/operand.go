package g8

import (
	"math"
	"strconv"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
)

func buildInt(b *builder, op *lex8.Token) *ref {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		b.Errorf(op.Pos, "invalid integer: %s", e)
		return nil
	}

	if ret < math.MinInt32 {
		b.Errorf(op.Pos, "integer too small to fit in 32-bit")
		return nil
	} else if ret > math.MaxUint32 {
		b.Errorf(op.Pos, "integer too large to fit in 32-bit")
		return nil
	} else if ret > math.MaxInt32 {
		// must be unsigned integer
		return newRef(types.Uint, ir.Num(uint32(ret)))
	}

	return newRef(types.Int, ir.Snum(int32(ret)))
}

func buildIdent(b *builder, op *lex8.Token) *ref {
	s := b.scope.Query(op.Lit)
	if s == nil {
		b.Errorf(op.Pos, "undefined identifer %s", op.Lit)
		return nil
	}

	switch s.Type {
	case symVar:
		v := s.Item.(*objVar)
		return v.ref
	case symFunc:
		v := s.Item.(*objFunc)
		return v.ref
	default:
		panic("todo")
	}
}

func buildOperand(b *builder, op *ast.Operand) *ref {
	switch op.Token.Type {
	case parse.Int:
		return buildInt(b, op.Token)
	case parse.Ident:
		return buildIdent(b, op.Token)
	default:
		panic("invalid or not implemented")
	}
}
