package link8

import (
	"fmt"

	"lonnie.io/e8vm/arch8"
)

func layout(used []pkgSym) (funcs, vars []pkgSym, e error) {
	// code
	codeStart := uint32(arch8.InitPC)
	codeEnd := codeStart
	codeMax := uint32(0x40000000)

	for _, ps := range used {
		typ := ps.Type()
		switch typ {
		case SymFunc:
			funcs = append(funcs, ps)

			f := ps.Func()
			f.addr = codeEnd
			codeEnd += f.Size()
			if codeEnd > codeMax {
				return nil, nil, fmt.Errorf("code section too large")
			}
		case SymVar:
			vars = append(vars, ps)
		default:
			panic("bug")
		}
	}

	// data
	for _, ps := range vars {
		if ps.Type() != SymVar {
			continue
		}

		panic("todo")
	}

	return
}
