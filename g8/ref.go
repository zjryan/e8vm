package g8

import (
	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/types"
)

// ref is a struct that
type ref struct {
	typ []types.T
	ir  []ir.Ref // this is essentially anything
}

// newRef creates a simple single ref
func newRef(t types.T, r ir.Ref) *ref {
	return &ref{[]types.T{t}, []ir.Ref{r}}
}

func (r *ref) Len() int       { return len(r.typ) }
func (r *ref) IsSingle() bool { return len(r.typ) == 1 }

func (r *ref) Type() types.T {
	if !r.IsSingle() {
		panic("not single")
	}
	return r.typ[0]
}

func (r *ref) IsType() bool {
	if !r.IsSingle() {
		return false
	}
	_, ok := r.Type().(*types.Type)
	return ok
}

func (r *ref) TypeType() types.T {
	return r.Type().(*types.Type).T
}

func (r *ref) IR() ir.Ref {
	if !r.IsSingle() {
		panic("not single")
	}
	return r.ir[0]
}

func (r *ref) String() string {
	if r == nil {
		return "void"
	}

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
