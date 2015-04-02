package ir

type op interface{}

type arithOp struct {
	dest Ref
	a    Ref
	op   string
	b    Ref
}

type callOp struct {
	dest []Ref
	f    Ref
	sig  *FuncSig
	args []Ref
}
