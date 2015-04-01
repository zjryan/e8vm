package g8

import (
	"lonnie.io/e8vm/g8/ir"
)

func getFuncSig(f *typFunc) *ir.FuncSig {
	if f.sig == nil {
		f.sig = makeFuncSig(f)
	}
	return f.sig
}

// converts a langauge function signature into a IR function signature
func makeFuncSig(f *typFunc) *ir.FuncSig {
	ret := new(ir.FuncSig)

	for i, t := range f.argTypes {
		name := ""
		if f.argNames != nil {
			name = f.argNames[i]
		}
		ret.AddArg(t.Size(), name)
	}

	for i, t := range f.retTypes {
		name := ""
		if f.retNames != nil {
			name = f.retNames[i]
		}
		ret.AddRet(t.Size(), name)
	}

	return ret
}
