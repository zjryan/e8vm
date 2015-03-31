package g8

import (
	"lonnie.io/e8vm/lex8"
)

func addressable(r *ref) bool {
	// TODO:
	return true
}

func canAssignType(left, right typ) bool {
	// TODO:
	return true
}

func checkAssignable(b *builder, pos *lex8.Pos, left, right *ref) bool {
	if !addressable(left) {
		b.Errorf(pos, "assigning to non-addressable")
		return false
	}

	if !canAssignType(left.typ, right.typ) {
		b.Errorf(pos, "cannot assign %s to %s",
			typStr(left.typ), typStr(right.typ),
		)
		return false
	}

	return true
}
