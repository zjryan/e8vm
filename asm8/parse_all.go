package asm8

type File struct {
	Funcs []*Func
}

func parseFile(p *Parser) *File {
	ret := new(File)

	for p.seeKeyword("import") {
		panic("todo")
	}

	for {
		if p.seeKeyword("func") {
			if f := parseFunc(p); f != nil {
				ret.Funcs = append(ret.Funcs, f)
			}
		} else if p.seeKeyword("var") {
			panic("todo")
		} else if p.seeKeyword("const") {
			panic("todo")
		} else {
			p.err(p.t.Pos, "expect top-declaration: func, var or const")
			return nil
		}
	}

	return ret
}
