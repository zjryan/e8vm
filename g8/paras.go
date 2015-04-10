package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
)

func buildNamedParaList(b *builder, lst *ast.ParaList) []*types.Arg {
	ret := make([]*types.Arg, lst.Len())
	// named typeed list
	for i, para := range lst.Paras {
		if para.Ident == nil {
			b.Errorf(ast.ExprPos(para.Type),
				"expect identifer as argument name",
			)
			return nil
		}

		name := para.Ident.Lit
		if name == "_" {
			name = ""
		}
		ret[i] = &types.Arg{Name: name}

		if para.Type == nil {
			continue
		}

		t := buildType(b, para.Type)
		if t == nil {
			return nil
		}

		// go back and assign types
		for j := i; j > 0 && ret[j].T == nil; j-- {
			ret[j].T = t
		}
	}

	// check that everything has a type
	if len(ret) > 0 && ret[len(ret)-1].T == nil {
		b.Errorf(lst.Rparen.Pos, "missing type in argument list")
	}

	return ret
}

func buildAnonyParaList(b *builder, lst *ast.ParaList) []*types.Arg {
	ret := make([]*types.Arg, lst.Len())
	for i, para := range lst.Paras {
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
			return nil
		}

		ret[i] = &types.Arg{T: t}
	}

	return ret
}

func buildParaList(b *builder, lst *ast.ParaList) []*types.Arg {
	if lst == nil {
		return make([]*types.Arg, 0)
	}
	if lst.Named() {
		return buildNamedParaList(b, lst)
	}
	return buildAnonyParaList(b, lst)
}
