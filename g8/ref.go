package g8

import (
	"lonnie.io/e8vm/g8/ir"
)

type ref struct {
	typ typ
	ir  ir.Ref
}

func newRef(t typ, r ir.Ref) *ref { return &ref{t, r} }

var voidRef = newRef(typVoid, nil)

func irRefs(list []*ref) []ir.Ref {
	n := len(list)
	if n == 0 {
		return nil
	}

	ret := make([]ir.Ref, 0, n)
	for _, r := range list {
		ret = append(ret, r.ir)
	}
	return ret
}
