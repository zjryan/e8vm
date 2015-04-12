package ir

import (
	"fmt"
)

// Func is an IR function. It consists of a bunch of named
// or unamed local variables and also a set of basic blocks.
// it can generate a linkable function.
type Func struct {
	id   int
	name string

	sig *FuncSig

	savedRegs []*stackVar
	locals    []*stackVar // local variables
	retAddr   *stackVar   // saved return address register

	prologue *Block
	epilogue *Block
	nblock   int

	nvar            int
	callerFrameSize int32 // frame size where the caller pushed
	frameSize       int32

	index  uint32 // the index in the lib
	isMain bool
}

func newFunc(name string, id int, sig *FuncSig) *Func {
	ret := new(Func)
	ret.id = id
	ret.name = name
	ret.sig = sig

	ret.prologue = ret.newBlock(nil)
	ret.epilogue = ret.newBlock(ret.prologue)

	return ret
}

// ArgRefs returns the refs to the arguments of the function
func (f *Func) ArgRefs() []Ref {
	var ret []Ref
	for _, arg := range f.sig.args {
		ret = append(ret, arg)
	}
	return ret
}

// RetRefs returns the refs to the return values of the function
func (f *Func) RetRefs() []Ref {
	var ret []Ref
	for _, arg := range f.sig.rets {
		ret = append(ret, arg)
	}
	return ret
}

// NewLocal creates a new named local variable of size n on stack.
func (f *Func) NewLocal(n int32, name string) Ref {
	ret := newVar(n, name)
	f.locals = append(f.locals, ret)
	return ret
}

// NewTemp creates a new temp variable of size n on stack.
func (f *Func) NewTemp(n int32) Ref {
	s := fmt.Sprintf("<%d>", f.nvar)
	f.nvar++
	return f.NewLocal(n, s)
}

func (f *Func) newBlock(after *Block) *Block {
	ret := new(Block)
	ret.id = f.nblock
	ret.frameSize = &f.frameSize

	f.nblock++

	if after != nil {
		ret.next = after.next
		ret.jump = after.jump

		after.next = ret
		after.jump = nil // jump to natural next, which is ret
	}

	return ret
}

// End returns the ending block of the function (the epilogue).
func (f *Func) End() *Block { return f.epilogue }

// NewBlock creates a new basic block for the function
func (f *Func) NewBlock(after *Block) *Block {
	if after == nil {
		after = f.prologue
	}
	ret := f.newBlock(after)
	return ret
}

// SetAsMain marks the function as main function
// and this function will have a bare metal prologue and epilogue.
func (f *Func) SetAsMain() { f.isMain = true }

func (f *Func) String() string { return f.name }
