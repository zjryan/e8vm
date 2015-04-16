package g8

import (
	"fmt"
	"path/filepath"

	"lonnie.io/e8vm/asm8"
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
	b.f = b.p.NewFunc(":start", ir.VoidFuncSig)
	b.fretNamed = false
	b.fretRef = nil
	b.f.SetAsMain()
	b.b = b.f.NewBlock(nil)

	b.scope.Push()
	b.buildStmts(stmts)
	b.scope.Pop()

	// write this to log file
	if b.irLog != nil {
		ir.PrintPkg(b.irLog, b.p)
	}

	return ir.BuildPkg(b.p) // do the code gen
}

func (bareFunc) Prepare(
	src map[string]*build8.File, importer build8.Importer,
) []*lex8.Error {
	importer.Import("$", "asm/builtin", nil)
	return nil
}

func (bareFunc) Compile(pinfo *build8.PkgInfo) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	b := newBuilder(pinfo.Path)

	initBuilder(b, pinfo.Import)
	if es = b.Errs(); es != nil {
		return nil, es
	}

	if len(pinfo.Src) == 0 {
		panic("no source file")
	} else if len(pinfo.Src) > 1 {
		e := fmt.Errorf("bare func %q has too many files", pinfo.Path)
		return nil, lex8.SingleErr(e)
	}

	b.irLog = pinfo.CreateLog("ir")
	defer b.irLog.Close()

	for _, r := range pinfo.Src {
		stmts, es := parse.Stmts(r.Path, r)
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

func buildSingle(fname, s string, lang build8.Lang) (
	image []byte, es []*lex8.Error, log []byte,
) {
	home := build8.NewMemHome(lang)
	home.AddLang("asm", asm8.Lang())

	pkg := home.NewPkg("main")
	name := filepath.Base(fname)
	pkg.AddFile(fname, name, s)

	builtin := home.NewPkg("asm/builtin")
	builtin.AddFile("", "builtin.s", builtInSrc)

	b := build8.NewBuilder(home)
	es = b.BuildAll()
	if es != nil {
		return nil, es, nil
	}

	return home.Bin("main"), nil, home.Log("main", "ir")
}

// CompileBareFunc compiles a bare function into a bare-metal E8 image
func CompileBareFunc(fname, s string) ([]byte, []*lex8.Error, []byte) {
	lang := BareFunc()
	return buildSingle(fname, s, lang)
}
