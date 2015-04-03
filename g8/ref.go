package g8

import (
	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/types"
)

// ref is a struct that
type ref struct {
	typ []types.Type
	ir  []ir.Ref // this is essentially anything
}

// newRef creates a simple single ref
func newRef(t types.Type, r ir.Ref) *ref {
	return &ref{[]types.Type{t}, []ir.Ref{r}}
}

func (r *ref) Len() int       { return len(r.typ) }
func (r *ref) IsSingle() bool { return len(r.typ) == 1 }

func (r *ref) Type() types.Type {
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

func (r *ref) String() string {
	if len(r.typ) == 0 {
		return "<nil>"
	}

	return fmt8.Join(r.typ, ",")
}

func (r *ref) IsBool() bool {
	if !r.IsSingle() {
		return false
	}
	return types.IsBasic(r.Type(), types.Bool)
}
