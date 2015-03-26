package ir

// Func is an IR function. It consists of a bunch of named
// or unamed local variables and also a set of basic blocks.
// it can generate a linkable function.
type Func struct {
	id      int
	name    string
	args    []*stackVar
	rets    []*stackVar
	locals  []*stackVar
	retAddr *stackVar

	vars      []*stackVar
	namedVars map[string]*stackVar
	blocks    []*Block
	body      []*Block

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

// AddArg adds an arg stack variable for the function.
func (f *Func) AddArg(name string, n int32) Ref {
	ret := f.newVar(name, n)
	f.args = append(f.args, ret)
	return ret
}

// AddRet adds a return value for the function.
func (f *Func) AddRet(name string, n int32) Ref {
	ret := f.newVar(name, n)
	f.rets = append(f.rets, ret)
	return ret
}

// NewLocal creates a new named local variable of size n on stack.
func (f *Func) NewLocal(name string, n int32) Ref {
	ret := f.newVar(name, n)
	f.locals = append(f.rets, ret)
	return ret
}

// NewTemp creates a new temp variable of size n on stack.
func (f *Func) NewTemp(n int32) Ref {
	return f.NewLocal("", n)
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
