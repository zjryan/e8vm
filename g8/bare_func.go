package g8

import (
	"fmt"
	"os"

	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// because bare function also uses builtin functions that comes from the
// building system, we also need to make it a simple language: a
// language with only one (implicit) main function
// In fact, we can simple "inherit" the basic lang
type bareFunc struct{ lang }

// BareFunc is a language where it only contains an implicit main function.
func BareFunc() build8.Lang { return bareFunc{lang{}} }

func buildBareFunc(b *builder, stmts []ast.Stmt) *link8.Pkg {
	b.f = b.p.NewFunc("main", ir.VoidFuncSig)
	b.b = b.f.NewBlock()

	for _, stmt := range stmts {
		buildStmt(b, stmt)
	}

	ir.PrintPkg(os.Stdout, b.p) // just for debugging...

	return ir.BuildPkg(b.p) // do the code gen
}

func (bareFunc) Prepare(
	src map[string]*build8.File, importer build8.Importer,
) []*lex8.Error {
	importer.Import("$", "asm/builtin", nil)
	return nil
}

func (bareFunc) Compile(
	path string, src map[string]*build8.File, imp map[string]*build8.Import,
) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	b := newBuilder(path)
	initBuilder(b, imp)
	if es = b.Errs(); es != nil {
		return nil, es
	}

	if len(src) == 0 {
		panic("no source file")
	}
	if len(src) > 1 {
		e := fmt.Errorf("bare func %q has too many files", path)
		return nil, lex8.SingleErr(e)
	}

	for f, r := range src {
		stmts, es := parse.Stmts(f, r)
		if es != nil {
			return nil, es
		}

		lib := buildBareFunc(b, stmts)
		if es = b.Errs(); es != nil {
			return nil, es
		}

		return &pkg{lib}, nil
	}

	panic("unreachable")
}
