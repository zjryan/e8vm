package ir

// Func is an IR function. It consists of a bunch of named
// or unamed local variables and also a set of basic blocks.
// it can generate a linkable function.
type Func struct {
	locals      []*stackVar
	namedLocals map[string]*stackVar
	blocks      []*Block

	frameSize uint32
}

func (f *Func) newLocalSize(name string, n int32) *stackVar {
	if name != "" {
		if f.namedLocals[name] != nil {
			panic("dup local name")
		}
	}

	ret := new(stackVar)
	ret.name = name
	ret.size = n
	ret.id = len(f.locals)

	f.locals = append(f.locals, ret)
	if name != "" {
		f.namedLocals[name] = ret
	}
	return ret
}

func (f *Func) newLocal(name string) *stackVar {
	return f.newLocalSize(name, 4)
}

func (f *Func) newTemp() *stackVar {
	return f.newLocal("")
}
