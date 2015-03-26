package link8

import (
	"fmt"

	"lonnie.io/e8vm/arch8"
)

func layout(used []pkgSym) (funcs, vars []pkgSym, e error) {
	// code
	pt := uint32(arch8.InitPC)
	codeMax := uint32(0xffffffff)

	for _, ps := range used {
		typ := ps.Type()
		switch typ {
		case SymFunc:
			funcs = append(funcs, ps)

			f := ps.Func()
			f.addr = pt
			size := f.Size()
			if size > codeMax-pt {
				return nil, nil, fmt.Errorf("code section too large")
			}
			pt += size
		case SymVar:
			vars = append(vars, ps)
		default:
			panic("bug")
		}
	}

	dataMax := uint32(0xffffffff)

	for _, ps := range vars {
		if ps.Type() != SymVar {
			panic("bug")
		}

		v := ps.Var()

		if v.align > 1 && pt%v.align != 0 {
			v.prePad = v.align - pt%v.align
			pt += v.prePad
		}
		if v.align > 1 && pt%v.align != 0 {
			panic("bug")
		}

		v.addr = pt
		size := v.Size()
		if size > dataMax-pt {
			return nil, nil, fmt.Errorf("binary too large")
		}

		pt += size
	}

	return
}
