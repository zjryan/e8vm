package ir

type op interface{}

type arithOp struct {
	dest ref
	a    ref
	op   string
	b    ref
}

type callOp struct {
	dest ref
	f    ref
	args []ref
}

type jump struct {
	a  ref
	op string
	b  ref
	to int // block id
}

func arith(dest ref, a ref, op string, b ref) *arithOp {
	return &arithOp{dest, a, op, b}
}

func call(dest ref, f ref, args ...ref) *callOp {
	return &callOp{dest, f, args}
}
