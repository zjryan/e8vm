package build8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

func isImportChar(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	if r == '/' {
		return true
	}
	return false
}

func lexImports(x *lex8.Lexer) *lex8.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '\n':
		return x.MakeToken(endl)
	case '/':
		return lex8.LexComment(x)
	case ';':
		return x.MakeToken(semi)
	}

	if r >= 'a' && r <= 'z' {
		for {
			x.Next()
			if x.Ended() || !isImportChar(x.Rune()) {
				break
			}
		}

		return x.MakeToken(operand)
	}

	return x.MakeToken(lex8.Illegal)
}

type importsLexer struct {
	x          *lex8.Lexer
	save       *lex8.Token
	insertSemi bool
}

func newImportsLexer(file string, r io.Reader) *importsLexer {
	ret := new(importsLexer)
	ret.x = lex8.NewLexer(file, r)
	ret.x.LexFunc = lexImports

	return ret
}

func (ix *importsLexer) Token() *lex8.Token {
	if ix.save != nil {
		ret := ix.save
		ix.save = nil
		return ret
	}

	for {
		t := ix.x.Token()
		switch t.Type {
		case semi:
			ix.insertSemi = false
		case lex8.EOF:
			if ix.insertSemi {
				ix.insertSemi = false
				ix.save = t
				return &lex8.Token{
					Type: semi,
					Lit:  t.Lit,
					Pos:  t.Pos,
				}
			}
		case endl:
			if ix.insertSemi {
				ix.insertSemi = false
				return &lex8.Token{
					Type: semi,
					Lit:  "\n",
					Pos:  t.Pos,
				}
			}
			continue
		case lex8.Comment:
			continue
		default:
			ix.insertSemi = true
		}

		return t
	}

}

func (ix *importsLexer) Errs() []*lex8.Error {
	return ix.x.Errs()
}
