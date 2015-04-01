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

	blocks   []*Block
	prologue *Block
	epilogue *Block
	body     []*Block

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

func (f *Func) newBlock() *Block {
	ret := new(Block)
	ret.id = len(f.blocks)
	f.blocks = append(f.blocks, ret)
	return ret
}

// NewBlock creates a new basic block for the function
func (f *Func) NewBlock() *Block {
	ret := f.newBlock()
	f.body = append(f.body, ret)
	return ret
}

// SetAsMain marks the function as main function
// and this function will have a bare metal prologue and epilogue.
func (f *Func) SetAsMain() { f.isMain = true }
