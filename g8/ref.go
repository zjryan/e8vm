package g8

import (
	"lonnie.io/e8vm/g8/ir"
)

type ref struct {
	typ typ
	ir  ir.Ref
}

func newRef(t typ, r ir.Ref) *ref { return &ref{t, r} }
