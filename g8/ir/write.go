package ir

import (
	"lonnie.io/e8vm/link8"
)

func writeBlock(f *link8.Func, b *Block) {
	for _, inst := range b.insts {
		f.AddInst(inst.inst)
		if inst.sym != nil {
			panic("todo")
		}
	}
}

func writeFunc(p *Pkg, f *Func) {
	lfunc := link8.NewFunc()

	writeBlock(lfunc, f.prologue)
	for _, b := range f.body {
		writeBlock(lfunc, b)
	}
	writeBlock(lfunc, f.epilogue)

	p.lib.DefineFunc(f.index, lfunc)
}
