package g8

import (
	"bytes"
	"io"
	"os"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

func buildBareFunc(b *builder, stmts []ast.Stmt) *link8.Pkg {
	b.f = b.p.NewFunc("main", ir.VoidFuncSig)
	b.b = b.f.NewBlock()

	for _, stmt := range stmts {
		buildStmt(b, stmt)
	}

	e := ir.PrintFunc(os.Stdout, b.f)
	if e != nil {
		panic(e)
	}

	return ir.BuildPkg(b.p)
}

// BuildBareFunc builds a bare main function of signature func main()
func BuildBareFunc(f string, r io.Reader) ([]byte, []*lex8.Error) {
	stmts, es := parse.Stmts(f, r)
	if es != nil {
		return nil, es
	}

	b := newBuilder("_")
	pkg := buildBareFunc(b, stmts)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	buf := new(bytes.Buffer)
	e := link8.LinkMain(pkg, buf, "main")
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return buf.Bytes(), nil
}
