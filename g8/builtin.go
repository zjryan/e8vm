package g8

import (
	"lonnie.io/e8vm/link8"
)

func declareBuiltin(b *builder, builtin *link8.Pkg) {
	o := func(name string, as string) {
		sym, index := builtin.SymbolByName(name)
		if sym == nil {
			b.Errorf(nil, "builtin symbol %s missing", name)
			return
		}

		_ = index
	}

	o("PrintUint32", "printUint")
	o("PrintChar", "printChar")
}
