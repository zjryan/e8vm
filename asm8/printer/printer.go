package printer

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

type printer struct {
	out io.Writer

	comments []*lex8.Token

	indent  int
	lineBuf *bytes.Buffer
	lineMid bool

	last *lex8.Token
}

func (p *printer) printEndl() {
	p.lineMid = false
}

func (p *printer) puts(s string) {

}

func (p *printer) printToken(t *lex8.Token) {
	o := func(s string) {
		fmt.Fprint(p.lineBuf, s)
		p.lineMid = true
	}

	switch t.Type {
	case parse.Lbrace:
		o("{")
	case parse.Rbrace:
		o("}")
	case parse.Keyword:
		o(t.Lit)
	case parse.Operand:
		o(t.Lit)
	case lex8.Comment:
		fmt.Fprint(p.lineBuf, t.Lit)
		if strings.HasPrefix(t.Lit, "//") {
			p.printEndl()
		}
	}
}

func posBefore(p1, p2 *lex8.Pos) bool {
	if p1.Line < p2.Line {
		return true
	}
	if p1.Line == p2.Line && p1.Col < p2.Col {
		return true
	}
	return false
}

func (p *printer) printComment(until *lex8.Pos) {
	if len(p.comments) == 0 {
		return
	}

	for {
		c := p.comments[0]
		if !posBefore(c.Pos, until) {
			break
		}
		p.printToken(c)
	}
}

func newPrinter(out io.Writer, comments []*lex8.Token) *printer {
	ret := new(printer)
	ret.out = out
	ret.comments = comments
	ret.lineBuf = new(bytes.Buffer)

	return ret
}

func (p *printer) Print(node interface{}) {
	switch node := node.(type) {
	case *ast.File:
		for _, d := range node.Decls {
			_ = &d
		}
	}
}

// PrintFile prints a file in its AST form
func PrintFile(out io.Writer, f *ast.File) {
	p := newPrinter(out, f.Comments)
	p.Print(f)
}
