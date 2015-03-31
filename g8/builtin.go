package g8

import (
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/sym8"
)

// makes a function type that returns nothing
func noRetFunc(ts ...typ) *typFunc { return &typFunc{argTypes: ts} }

func declareBuiltin(b *builder, builtin *link8.Pkg) {
	pindex := b.p.Require(builtin)

	o := func(name string, as string, t *typFunc) {
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

	o("PrintUint32", "printUint", noRetFunc(typUint))
	o("PrintChar", "printChar", noRetFunc(typUint8))
}
