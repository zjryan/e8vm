package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildFile(b *builder, f *ast.File) {
	// TODO: build imports

	var funcs []*objFunc

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.Func:
			f := declareFunc(b, d)
			if f != nil {
				funcs = append(funcs, f)
			}
		case *ast.VarDecls:
			b.Errorf(d.Kw.Pos, "var declaration not implemented")
		case *ast.ConstDecls:
			b.Errorf(d.Kw.Pos, "func declaration not implemented")
		default:
			b.Errorf(nil, "invlaid top declare: %T", d)
		}
	}
}
