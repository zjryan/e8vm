package g8

import (
	"os"

	"lonnie.io/e8vm/link8"
)

func declareBuiltin(b *builder, builtin *link8.Pkg) {
	builtin.PrintSymbols(os.Stdout)

	o := func(name string, as string) {
		sym, index := builtin.SymbolByName(name)
		if sym == nil {
			b.Errorf(nil, "%s not found in builtin package", name)
			return
		}

		_ = index
	}

	o("PrintUint32", "printUint")
	o("PrintChar", "printChar")
}
