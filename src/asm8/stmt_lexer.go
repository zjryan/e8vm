package asm8

import (
	"io"

	"lex8"
)

// StmtLexer replaces end-lines with semicolons
type StmtLexer struct {
	x          *lex8.Lexer
	save       *lex8.Token
	insertSemi bool

	ParseComment bool
}

// NewStmtLexer creates a new statement lexer.
func NewStmtLexer(file string, r io.ReadCloser) *StmtLexer {
	ret := new(StmtLexer)
	ret.x = NewLexer(file, r)

	return ret
}

// Token returns the next token of lexing
func (sx *StmtLexer) Token() *lex8.Token {
	if sx.save != nil {
		ret := sx.save
		sx.save = nil
		return ret
	}

	for {
		t := sx.x.Token()
		switch t.Type {
		case Lbrace, Semi:
			sx.insertSemi = false
		case EOF:
			if sx.insertSemi {
				sx.insertSemi = false
				sx.save = t
				return &lex8.Token{Semi, t.Lit, t.Pos}
			}
		case Rbrace:
			if sx.insertSemi {
				sx.save = t
				return &lex8.Token{Semi, t.Lit, t.Pos}
			}
			sx.insertSemi = true
		case Endl:
			if sx.insertSemi {
				sx.insertSemi = false
				return &lex8.Token{Semi, "\n", t.Pos}
			}
			continue // ignore this end line
		case Comment:
			if !sx.ParseComment {
				continue
			}
		default:
			sx.insertSemi = true
		}

		return t
	}
}

// Errs returns the list of lexing errors.
func (sx *StmtLexer) Errs() []*lex8.Error {
	return sx.x.Errs()
}
