package ir

// FuncSig describes the function signature of a callable
// function
type FuncSig struct {
	args []*stackVar
	rets []*stackVar
}

// AddArg adds an arg stack variable for the function.
func (f *FuncSig) AddArg(n int32, name string) Ref {
	ret := newVar(n, name)
	f.args = append(f.args, ret)
	return ret
}

// AddRet adds a return value for the function.
func (f *FuncSig) AddRet(n int32, name string) Ref {
	ret := newVar(n, name)
	f.rets = append(f.rets, ret)
	return ret
}

// VoidFuncSig is the signature for a function that has no
// parameters and no return values.
var VoidFuncSig = new(FuncSig)
