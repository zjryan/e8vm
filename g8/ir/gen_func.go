package ir

import (
	"lonnie.io/e8vm/arch8"
)

/**
## E8VM calling convention

- r0, always keep it as zero, never touch it
- r1, the first arg or return
- r2, the second arg or return
- r3, the third arg or return
- r4, the forth arg or return
- sp, stack pointer
- ret, return address
- pc, the program counter

other args are pushed on the stack

## Function Prologue
- push ret to the stack
- push r1-r4 to the stack, for archive

## Function Epilogue
- save ret values to the stack
-

**/

func layoutArgs(f *Func, args []*stackVar) []*stackVar {
	nreg := uint32(0)
	frameSize := int32(0)
	var regArgs []*stackVar

	for _, arg := range args {
		if nreg >= 4 || arg.size != regSize {
			arg.offset = frameSize
			frameSize += arg.size
			continue
		}

		// this arg is now sent in via a register
		nreg++
		arg.viaReg = nreg
		regArgs = append(regArgs, arg)
	}

	if frameSize > f.callerFrameSize {
		f.callerFrameSize = frameSize
	}

	return regArgs
}

// pushLocal allocates a frame slot for the local var
func pushVar(f *Func, vars ...*stackVar) {
	for _, v := range vars {
		v.offset = f.frameSize
		f.frameSize += v.size
	}
}

func layoutLocals(f *Func) {
	regArgs := layoutArgs(f, f.args)
	regRets := layoutArgs(f, f.rets)

	// layout the variables in the function
	f.frameSize = f.callerFrameSize
	f.retAddr = f.newVar("", regSize)
	f.retAddr.viaReg = arch8.RET // the return address

	pushVar(f, f.retAddr)
	// if all args and rets are via register
	// then f.retAddr.offset should be 0

	pushVar(f, regArgs...)
	pushVar(f, regRets...)
}

func makePrologue(f *Func) *Block {
	b := f.newBlock()
	saveVar(b, arch8.RET, f.retAddr)
	for _, v := range f.args {
		if v.viaReg == 0 {
			continue // skip args not sent in via register
		}
		saveVar(b, v.viaReg, v)
	}
	return b
}

func makeEpilogue(f *Func) *Block {
	b := f.newBlock()
	for _, v := range f.rets {
		if v.viaReg == 0 {
			continue
		}
		loadVar(b, v.viaReg, v)
	}
	// back to the caller
	loadVar(b, arch8.PC, f.retAddr)
	return b
}

func genFunc(p *Pkg, f *Func) {
	layoutLocals(f)

	f.prologue = makePrologue(f)
	f.epilogue = makeEpilogue(f)
	for _, b := range f.blocks {
		genBlock(b)
	}
}
