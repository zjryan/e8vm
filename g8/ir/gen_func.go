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

func alignUp(size, align int32) int32 {
	mod := size % align
	if mod == 0 {
		return size
	}
	return size + align - mod
}

func layoutArgs(f *Func, args []*stackVar) ([]*stackVar, []bool) {
	nreg := uint32(0)
	frameSize := int32(0)
	var regArgs []*stackVar
	regs := make([]bool, 5) // only need to track r1-r4

	for _, arg := range args {
		if nreg >= 4 || arg.size > regSize {
			size := alignUp(arg.size, regSize)
			if size%regSize != 0 {
				panic("invalid arg size")
			}
			arg.offset = frameSize
			frameSize += size
			continue
		}

		// this arg is now sent in via a register
		nreg++
		arg.viaReg = nreg
		regs[nreg] = true
		regArgs = append(regArgs, arg)
	}

	if frameSize > f.callerFrameSize {
		f.callerFrameSize = frameSize
	}

	return regArgs, regs
}

// pushVar allocates a frame slot for the local var
func pushVar(f *Func, vars ...*stackVar) {
	for _, v := range vars {
		size := alignUp(v.size, regSize)
		v.offset = f.frameSize
		f.frameSize += size
	}
}

func layoutLocals(f *Func) {
	regArgs, _ := layoutArgs(f, f.args)
	regRets, regUsed := layoutArgs(f, f.rets)

	var savedRegs []*stackVar
	for i := uint32(_1); i <= _4; i++ {
		if !regUsed[i] {
			v := f.newVar(regSize, "")
			v.viaReg = i
			savedRegs = append(savedRegs, v)
		}
	}
	f.savedRegs = savedRegs

	// layout the variables in the function
	f.frameSize = f.callerFrameSize
	f.retAddr = f.newVar(regSize, "")
	f.retAddr.viaReg = arch8.RET // the return address

	// if all args and rets are via register
	// then f.retAddr.offset should be 0

	pushVar(f, regArgs...)
	pushVar(f, regRets...)
	pushVar(f, savedRegs...)
	pushVar(f, f.locals...)

	// and we push retAddr at the end
	pushVar(f, f.retAddr)
}

func makePrologue(f *Func) *Block {
	b := f.newBlock()

	saveRet(f, b, f.retAddr)
	// move the sp
	b.inst(asm.addi(_sp, _sp, -f.frameSize))

	for _, v := range f.args {
		if v.viaReg == 0 {
			continue // skip args not sent in via register
		}
		saveVar(b, v.viaReg, v)
	}

	// this is for restoreing the registers
	for _, v := range f.savedRegs {
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

	for _, v := range f.savedRegs {
		loadVar(b, v.viaReg, v) // restoring the registers
	}

	b.inst(asm.addi(_sp, _sp, f.frameSize))
	// back to the caller
	loadRet(f, b, f.retAddr)
	return b
}

func genFunc(p *Pkg, f *Func) {
	layoutLocals(f)

	f.prologue = makePrologue(f)
	f.epilogue = makeEpilogue(f)

	for _, b := range f.body {
		genBlock(b)
	}
}
