package asm8

// GenFunc is a function semantics checker and code generator
type GenFunc struct {
	funcs []*Func
	objs  []*FuncObj
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

// Gen gererates the functions, and return a list of function objects.
// It retunrs with type []*FuncObj.
func (g *GenFunc) Gen(b *Builder) interface{} {
	for _, f := range g.funcs {
		obj := g.genFunc(b, f)
		g.objs = append(g.objs, obj)
	}

	return g.objs
}

func (g *GenFunc) genFunc(b *Builder, f *Func) *FuncObj {
	ret := new(FuncObj)

	return ret
}
