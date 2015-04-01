package ir

// FuncArg is a function arg
type FuncArg struct {
	Name string
	Size int32
}

// FuncSig describes the function signature of a callable
// function
type FuncSig struct {
	args []*stackVar
	rets []*stackVar

	regArgs []*stackVar // args sent in via registers
	regRets []*stackVar // return values sent out via registers

	argRegUsed []bool
	frameSize  int32
}

// NewFuncSig creates a new function call signature
func NewFuncSig(args, rets []*FuncArg) *FuncSig {
	ret := new(FuncSig)
	for _, arg := range args {
		v := newVar(arg.Size, arg.Name)
		ret.args = append(ret.args, v)
	}
	for _, arg := range rets {
		v := newVar(arg.Size, arg.Name)
		ret.rets = append(ret.rets, v)
	}

	ret.regArgs, ret.argRegUsed = layoutFuncArgs(ret, ret.args)
	ret.regRets, _ = layoutFuncArgs(ret, ret.rets)

	return ret
}

func layoutFuncArgs(f *FuncSig, args []*stackVar) ([]*stackVar, []bool) {
	const viaRegMax = 3

	frameSize := int32(0)
	nreg := uint32(0)
	regUsed := make([]bool, viaRegMax+1) // only track r1-r3
	regArgs := make([]*stackVar, 0, viaRegMax)

	for _, arg := range args {
		if nreg >= viaRegMax || arg.size > regSize {
			size := alignUp(arg.size, regSize)
			arg.offset = frameSize
			frameSize += size
			continue
		}

		nreg++
		arg.viaReg = nreg
		regUsed[nreg] = true
		regArgs = append(regArgs, arg)
	}

	if frameSize > f.frameSize {
		f.frameSize = frameSize
	}

	return regArgs, regUsed
}

// VoidFuncSig is the signature for a function that has no
// parameters and no return values.
var VoidFuncSig = NewFuncSig(nil, nil)
