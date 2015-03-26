package ir

// Func is an IR function. It consists of a bunch of named
// or unamed local variables and also a set of basic blocks.
// it can generate a linkable function.
type Func struct {
	id      int
	args    []*stackVar
	rets    []*stackVar
	locals  []*stackVar
	retAddr *stackVar

	vars      []*stackVar
	namedVars map[string]*stackVar
	blocks    []*Block

	prologue *Block
	epilogue *Block

	callerFrameSize int32
	frameSize       int32

	index uint32 // the index in the lib
}

func (f *Func) newVar(
	name string,
	n int32,
) *stackVar {
	if name != "" {
		if f.namedVars[name] != nil {
			panic("dup var name")
		}
	}

	ret := new(stackVar)
	ret.name = name
	ret.size = n
	ret.id = len(f.locals)

	f.vars = append(f.vars, ret)
	if name != "" {
		f.namedVars[name] = ret
	}

	return ret
}

const regSize = 4

func (f *Func) newArg(name string, n int32) *stackVar {
	ret := f.newVar(name, n)
	f.args = append(f.args, ret)
	return ret
}

func (f *Func) newRet(name string, n int32) *stackVar {
	ret := f.newVar(name, n)
	f.rets = append(f.rets, ret)
	return ret
}

func (f *Func) newLocal(name string, n int32) *stackVar {
	ret := f.newVar(name, n)
	f.locals = append(f.rets, ret)
	return ret
}

func (f *Func) newTemp(n int32) *stackVar {
	return f.newLocal("", n)
}

func (f *Func) newBlock() *Block {
	ret := new(Block)
	ret.id = len(f.blocks)
	return ret
}
