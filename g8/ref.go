package g8

import (
	"lonnie.io/e8vm/g8/ir"
)

// ref is a struct that 
type ref struct {
	typ []typ
	ir  []ir.Ref // this is essentially anything
}

// newRef creates a simple single ref
func newRef(t typ, r ir.Ref) *ref { 
	return &ref{[]typ{t}, []ir.Ref{r}} 
}

var voidRef = newRef(typVoid, nil)

func (r *ref) Len() int { return len(r.typ) }
func (r *ref) IsSingle() bool { return len(r.typ) == 1 }

func (r *ref) Typ() typ {
	if !r.IsSingle() {
		panic("not single")
	}
	return r.typ[0]
}

func (r *ref) IR() ir.Ref {
	if !r.IsSingle() {
		panic("not single")
	}
	return r.ir[0]
}

