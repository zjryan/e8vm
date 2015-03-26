package g8

import (
	"math"
	"strconv"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
)

// parseInt parses an signed or unsigned 32-bit integer
func parseInt(b *builder, op *lex8.Token) typ {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		b.Errorf(op.Pos, "invalid integer: %s", e)
		return nil
	}

	if ret < math.MinInt32 {
		b.Errorf(op.Pos, "integer too small, not fit in 32-bit")
		return nil
	} else if ret > math.MaxUint32 {
		b.Errorf(op.Pos, "integer too large, not fit in 32-bit")
		return nil
	} else if ret > math.MaxInt32 {
		// must be unsigned integer
		return constUint(uint32(ret))
	}

	return constInt(int32(ret))
}

func buildOperand(b *builder, op *ast.Operand) *ref {
	switch op.Token.Type {
	case parse.Int:
		t := parseInt(b, op.Token)
		if t == nil {
			return nil
		}

		switch t := t.(type) {
		case constUint:
			return newRef(t, ir.Num(uint32(t)))
		case constInt:
			return newRef(t, ir.Snum(int32(t)))
		default:
			panic("unknwon integer type")
		}
	default:
		panic("invalid or not implemented")
	}
}

func buildExpr(b *builder, expr ast.Expr) *ref {
	if expr == nil {
		return nil
	}

	switch expr := expr.(type) {
	default:
		_ = expr
		panic("invalid or not implemented")
	}
}

func buildExprList(b *builder, list *ast.ExprList) []*ref {
	ret := make([]*ref, 0, list.Len())
	for _, expr := range list.Exprs {
		ref := buildExpr(b, expr)
		if ref == nil {
			return nil
		}
		ret = append(ret, ref)
	}
	return ret
}
