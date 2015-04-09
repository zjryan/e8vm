package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
)

func buildFunc(b *builder, f *ast.Func) {
	name := f.Name.Lit

	b.scope.Push()
	defer b.scope.Pop()

	args := make([]*types.Arg, f.Args.Len())
	if f.Args.Named() {
		// named typeed list
		for i, para := range f.Args.Paras {
			if para.Ident == nil {
				b.Errorf(ast.ExprPos(para.Type),
					"expect identifer as argument name",
				)
				return
			}

			name := para.Ident.Lit
			if name == "_" {
				name = ""
			}
			args[i] = &types.Arg{Name: name}

			if para.Type == nil {
				continue
			}

			t := buildType(b, para.Type)
			if t == nil {
				return
			}

			// go back and assign types
			for j := i; j > 0 && args[j].T == nil; j-- {
				args[j].T = t
			}
		}

		// check that everything has a type
		if len(args) > 0 && args[len(args)-1].T == nil {
			b.Errorf(f.Args.Rparen.Pos, "missing type in argument list")
		}
	} else {
		// anonymous typed list
		for i, para := range f.Args.Paras {
			if para.Ident != nil && para.Type != nil {
				// anonymous typed list must all be single
				panic("bug")
			}

			var t types.T
			expr := para.Type
			if expr == nil {
				expr = &ast.Operand{para.Ident}
			}

			t = buildType(b, expr)
			if t == nil {
				return
			}

			args[i] = &types.Arg{T: t}
		}
	}

	// got the return values
	var rets []*types.Arg
	if f.RetType == nil {
		rets = make([]*types.Arg, f.Rets.Len())

	} else {
		retType := buildType(b, f.RetType)
		if retType == nil {
			return
		}

		rets = []*types.Arg{{T: retType}}
	}

	ftype := types.NewFuncNamed(args, rets)
	b.f = b.p.NewFunc(name, ftype.Sig) // switch to this function
	// NewFunc() will create the variables required for the sigs

	argRefs := b.f.ArgRefs() // we need these refs to declare
	if len(argRefs) != len(args) {
		panic("bug")
	}

	_ = argRefs

	retRefs := b.f.RetRefs() // we need these refs to declare
	if len(retRefs) != len(rets) {
		panic("bug")
	}

	_ = retRefs
}
