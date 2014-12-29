package asm8

import (
	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/lex8"
)

func layout(used []pkgSym) (funcs, vars []pkgSym, e *lex8.Error) {
	// code
	codeStart := uint32(arch8.InitPC)
	codeEnd := codeStart
	codeMax := uint32(0x40000000)

	for _, u := range used {
		s := u.p.symbols[u.sym]
		switch s.Type {
		case SymFunc:
			funcs = append(funcs, u)
			f := u.p.FuncObj(u.sym)
			f.addr = codeEnd
			codeEnd += f.Size()
			if codeEnd > codeMax {
				return nil, nil, lex8.Errorf("code section too large")
			}
		case SymVar:
			vars = append(vars, u)
		case SymConst:
			// nothing to do
		default:
			panic("bug")
		}
	}

	// data
	for _, u := range vars {
		s := u.p.symbols[u.sym]
		if s.Type != SymVar {
			continue
		}

		// todo
	}

	return
}
