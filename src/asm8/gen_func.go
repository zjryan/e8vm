package asm8

// GenFunc is a function semantics checker and code generator
type GenFunc struct {
	funcs []*Func
}

func isIdent(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r >= 'Z' {
			continue
		}
		if r == '_' || r == ':' {
			continue
		}
		if i > 0 && r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}

// Register registers a function. The block must be of type (*Func).
func (g *GenFunc) Register(b *Builder, block interface{}) {
	f := block.(*Func)
	name := f.name.Lit
	if !isIdent(name) {
		b.err(f.name.Pos, "invalid function name %q", name)
		return
	}

	// declare the symbol
	pre := b.scope.Declare(&Symbol{name, SymFunc, f, f.name.Pos})
	if pre != nil {
		b.err(f.name.Pos, "%q already defined as a %s", name, symStr(pre.Type))
		if pre.Pos != nil {
			b.err(pre.Pos, "   previously defined here")
		}
		return
	}

	// now you are registered
	g.funcs = append(g.funcs, f)
}
