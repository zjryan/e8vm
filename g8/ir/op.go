package ir

type op interface{}

type arithOp struct {
	dest Ref
	a    Ref
	op   string
	b    Ref
}

type callOp struct {
	dest Ref
	f    Ref
	args []Ref
}

type jump struct {
	a  Ref
	op string
	b  Ref
	to int // block id
}

func arith(dest Ref, a Ref, op string, b Ref) *arithOp {
	return &arithOp{dest, a, op, b}
}

func call(dest Ref, f Ref, args ...Ref) *callOp {
	return &callOp{dest, f, args}
}
