package g8

import (
	"os"
	"strings"

	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
)

type lang struct{}

// Lang returns the G language builder for the building system
func Lang() build8.Lang { return lang{} }

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".g")
}

func (lang) Prepare(
	src map[string]*build8.File, importer build8.Importer,
) []*lex8.Error {
	importer.Import("$", "asm/builtin", nil)
	// TODO: parse import statements
	return nil
}

func initBuilder(b *builder, imp map[string]*build8.Import) {
	b.exprFunc = buildExpr
	b.stmtFunc = buildStmt

	builtin, ok := imp["$"]
	if !ok {
		b.Errorf(nil, "builtin import missing for %q", b.path)
		return
	}

	declareBuiltin(b, builtin.Compiled.Lib())
}

func parsePkg(pinfo *build8.PkgInfo) (map[string]*ast.File, []*lex8.Error) {
	var parseErrs []*lex8.Error
	asts := make(map[string]*ast.File)
	for name, src := range pinfo.Src {
		f, es := parse.File(src.Path, src)
		if es != nil {
			parseErrs = append(parseErrs, es...)
		}
		asts[name] = f
	}
	if len(parseErrs) > 0 {
		return nil, parseErrs
	}

	return asts, nil
}

func addStart(b *builder) {
	s := b.scope.Query("main")
	f, isFunc := s.Item.(*objFunc)
	if !isFunc { // main is not a function
		return
	}
	if !types.SameType(f.ref.Type(), types.MainFuncSig) {
		// main is not of "func main()"
		return
	}

	b.f = b.p.NewFunc(":start", ir.VoidFuncSig)
	b.f.SetAsMain()
	b.b = b.f.NewBlock(nil)
	b.b.Call(nil, f.IR(), ir.VoidFuncSig)
}

func (lang) Compile(pinfo *build8.PkgInfo) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	asts, es := parsePkg(pinfo)
	if es != nil {
		return nil, es
	}

	// need to load these two builtin functions here
	b := newBuilder(pinfo.Path)
	initBuilder(b, pinfo.Import)
	if es = b.Errs(); es != nil {
		return nil, es
	}

	b.scope.Push() // package scope
	defer b.scope.Pop()
	for _, fileAST := range asts {
		buildFile(b, fileAST)
	}

	if es = b.Errs(); es != nil {
		return nil, es
	}

	addStart(b)

	ir.PrintPkg(os.Stdout, b.p)
	lib := ir.BuildPkg(b.p)
	return &pkg{lib}, nil
}
