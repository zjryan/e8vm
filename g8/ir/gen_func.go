package ir

import (
	"lonnie.io/e8vm/arch8"
)

/**
## E8VM calling convention

- r0, always keep it as zero, never touch it
- r1, the first arg or return, if not used, should keep the value
- r2, the second arg or return
- r3, the third arg or return
- r4, free form temp
- sp, stack pointer
- ret, return address
- pc, the program counter

other args are pushed on the stack

## Function Prologue
- push ret to the stack
- push r1-r3 to the stack, for archive

## Function Epilogue
- save ret values to the stack

**/

// pushVar allocates a frame slot for the local var
func pushVar(f *Func, vars ...*stackVar) {
	for _, v := range vars {
		size := alignUp(v.size, regSize)
		v.offset = f.frameSize
		f.frameSize += size
	}
}

func layoutLocals(f *Func) {
	for i, used := range f.sig.argRegUsed {
		if used {
			continue
		}

		// the caller is not using this reg for sending
		// the argument, the callee hence needs to
		// save this register
		v := newVar(regSize, "")
		v.viaReg = uint32(i)
		f.savedRegs = append(f.savedRegs, v)
	}

	// layout the variables in the function
	f.frameSize = f.sig.frameSize
	f.retAddr = newVar(regSize, "")
	f.retAddr.viaReg = arch8.RET // the return address

	// if all args and rets are via register
	// then f.retAddr.offset should be 0, it is the nearest to SP
	pushVar(f, f.retAddr)
	pushVar(f, f.sig.regArgs...)
	pushVar(f, f.sig.regRets...)
	pushVar(f, f.savedRegs...)
	pushVar(f, f.locals...)
}

func makeMainPrologue(f *Func) *Block {
	b := f.newBlock()
	b.inst(asm.xor(_0, _0, _0))
	b.inst(asm.lui(_sp, 0x1000))
	b.inst(asm.addi(_sp, _sp, -f.frameSize))
	return b
}

func makeMainEpilogue(f *Func) *Block {
	b := f.newBlock()
	b.inst(asm.addi(_sp, _sp, f.frameSize))
	b.inst(asm.halt())
	return b
}

func makePrologue(f *Func) *Block {
	b := f.newBlock()
	b.frameSize = f.frameSize

	saveRetAddr(b, f.retAddr)
	// move the sp
	b.inst(asm.addi(_sp, _sp, -f.frameSize))

	for _, v := range f.sig.args {
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
	b.frameSize = f.frameSize

	for _, v := range f.savedRegs {
		loadVar(b, v.viaReg, v) // restoring the registers
	}

	for _, v := range f.sig.rets {
		if v.viaReg == 0 {
			continue
		}
		loadVar(b, v.viaReg, v)
	}

	b.inst(asm.addi(_sp, _sp, f.frameSize))
	// back to the caller
	loadRetAddr(b, f.retAddr)
	return b
}

func genFunc(p *Pkg, f *Func) {
	layoutLocals(f)

	if f.isMain {
		f.prologue = makeMainPrologue(f)
		f.epilogue = makeMainEpilogue(f)
	} else {
		f.prologue = makePrologue(f)
		f.epilogue = makeEpilogue(f)
	}

	for _, b := range f.body {
		b.frameSize = f.frameSize
		genBlock(b)
	}
}
