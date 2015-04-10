package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/sym8"
)

func buildFuncType(b *builder, f *ast.Func) *types.Func {
	// the arguments
	args := buildParaList(b, f.Args)
	if args == nil {
		return nil
	}

	// the return values
	var rets []*types.Arg
	if f.RetType == nil {
		rets = buildParaList(b, f.Rets)
	} else {
		retType := buildType(b, f.RetType)
		if retType == nil {
			return nil
		}
		rets = []*types.Arg{{T: retType}}
	}

	return types.NewFuncNamed(args, rets)
}

func declareFunc(b *builder, f *ast.Func) *objFunc {
	ftype := buildFuncType(b, f)
	if ftype == nil {
		return nil
	}

	// NewFunc() will create the variables required for the sigs
	name := f.Name.Lit
	ret := new(objFunc)
	ret.name = name
	ret.f = f
	irFunc := b.p.NewFunc(name, ftype.Sig)
	ret.ref = newRef(ftype, irFunc)

	// add this item to the top scope
	s := sym8.Make(name, symFunc, ret, f.Name.Pos)
	conflict := b.scope.Declare(s) // lets declare the function
	if conflict != nil {
		b.Errorf(f.Name.Pos, "%q already declared as a %s",
			name, symStr(conflict.Type),
		)
		b.Errorf(conflict.Pos, "previously declared here")
		return nil
	}

	return ret
}

func declareParas(b *builder,
	lst *ast.ParaList, ts []*types.Arg, irs []ir.Ref,
) {
	if len(ts) != len(irs) {
		panic("bug")
	}

	for i, t := range ts {
		if t.Name == "" {
			continue
		}

		r := newRef(t.T, irs[i])
		declareVar(b, lst.Paras[i].Ident, r)
	}
}

func buildFunc(b *builder, f *objFunc) {
	b.scope.Push() // func body scope
	defer b.scope.Pop()

	t := f.ref.Type().(*types.Func) // the signature of the function
	irFunc := f.ref.IR().(*ir.Func)
	b.f = irFunc

	declareParas(b, f.f.Args, t.Args, irFunc.ArgRefs())
	declareParas(b, f.f.Rets, t.Rets, irFunc.RetRefs())

	b.buildStmts(f.f.Body.Stmts)
}
