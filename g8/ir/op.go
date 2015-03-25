package ir

type op interface{}

type arithOp struct {
	dest ref
	a    ref
	op   string
	b    ref
}

type jump struct {
	a  ref
	op string
	b  ref
	to int // block id
}
