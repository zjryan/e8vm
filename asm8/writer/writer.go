package writer

import (
	"fmt"
	"io"
	"strings"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

type Writer struct {
	out io.Writer

	file *ast.File
	f    *ast.FuncDecl
	v    *ast.VarDecl
}

func makeTok(lit string, t int) *lex8.Token {
	ret := new(lex8.Token)
	ret.Lit = lit
	ret.Type = t
	return ret
}

func makeOp(lit string) *lex8.Token {
	return makeTok(lit, parse.Operand)
}

// Func declares a new function.
func (w *Writer) Func(name string) {
	if !parse.IsIdent(name) {
		panic("invalid function name")
	}

	w.v = nil
	w.f = new(ast.FuncDecl)
	w.f.Name = makeTok(name, parse.Operand)

	w.file.Funcs = append(w.file.Funcs, w.f)
}

func (w *Writer) Instf(f string, args ...interface{}) {
	w.instLine(fmt.Sprintf(f, args...))
}

func (w *Writer) Inst(args ...interface{}) {
	w.instLine(fmt.Sprint(args...))
}

func (w *Writer) instLine(line string) {
	fields := strings.Fields(line)
	ops := make([]*lex8.Token, 0, len(fields))
	for i, field := range fields {
		ops[i] = makeOp(field)
	}

	stmt := new(ast.FuncStmt)
	stmt.Ops = ops

	w.f.Stmts = append(w.f.Stmts, stmt)
}

func (w *Writer) Label(label string) {
	stmt := new(ast.FuncStmt)
	stmt.Label = label

	w.f.Stmts = append(w.f.Stmts, stmt)
}

// Var
func (w *Writer) Var(name string) {
}

// Str adds a string
func (w *Writer) Str(strs ...string) {

}

func (w *Writer) Int(nums ...int32) {
}

func (w *Writer) Uint(nums ...int32) {

}

func (w *Writer) Hex(bs []byte) {

}

func (w *Writer) Int8(nums ...int8) {

}

func (w *Writer) Uint8(nums ...uint8) {

}
