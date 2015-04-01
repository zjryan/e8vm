package g8

import (
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/sym8"
)

func declareBuiltin(b *builder, builtin *link8.Pkg) {
	pindex := b.p.Require(builtin)

	o := func(name string, as string, t *types.Func) {
		sym, index := builtin.SymbolByName(name)
		if sym == nil {
			b.Errorf(nil, "builtin symbol %s missing", name)
			return
		}

		ref := ir.FuncSym(pindex, index, nil) // a reference to the function
		obj := &objFunc{as, newRef(t, ref)}
		pre := b.scope.Declare(sym8.Make(as, symFunc, obj, nil))
		if pre != nil {
			b.Errorf(nil, "builtin symbol %s declare failed", name)
			return
		}
	}

	o("PrintInt32", "printInt", types.NewVoidFunc(types.Int))
	o("PrintUint32", "printUint", types.NewVoidFunc(types.Uint))
	o("PrintChar", "printChar", types.NewVoidFunc(types.Uint8))
}
