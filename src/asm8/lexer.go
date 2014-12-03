package asm8

import (
	"io"
)

// Lexer parses a file input stream into tokens.
type Lexer struct {
	s *LexScanner

	r    rune
	e    error
	errs *ErrorList

	isWhite func(r rune) bool
	lexFunc func(x *Lexer) *Token
}

func newLexer(file string, r io.ReadCloser) *Lexer {
	ret := new(Lexer)
	ret.s = NewLexScanner(file, r)
	ret.errs = NewErrList()

	ret.isWhite = func(r rune) bool {
		return r == ' ' || r == '\t'
	}

	ret.next()

	return ret
}

// some helper functions
func (x *Lexer) next() (rune, error) {
	x.r, x.e = x.s.Next()
	return x.r, x.e
}
func (x *Lexer) token(t int) *Token {
	s, p := x.s.Accept()
	return &Token{t, s, p}
}
func (x *Lexer) discard()  { x.s.Accept() }
func (x *Lexer) eof() bool { return x.e != nil }

func (x *Lexer) skipWhite() {
	for {
		if x.eof() || !x.isWhite(x.r) {
			break
		}
		x.next()
	}
	x.discard()
}

// Token returns the next parsed token.
// It ends with a token with type EOF.
func (x *Lexer) Token() *Token {
	x.skipWhite()

	if x.eof() {
		return x.token(EOF)
	}

	if x.lexFunc == nil {
		x.next()
		return x.token(Illegal)
	}

	return x.lexFunc(x)
}

func (x *Lexer) err(f string, args ...interface{}) {
	x.errs.Addf(x.s.Pos(), f, args...)
}

// Errs returns the lexing errors.
func (x *Lexer) Errs() []*Error {
	if x.e != nil && x.e != io.EOF {
		return []*Error{{Err: x.e}}
	}

	return x.errs.Errs
}

// Tokens breaks a file input stream into tokens or errors.
func Tokens(f string, rc io.ReadCloser) ([]*Token, []*Error) {
	x := NewLexer(f, rc)
	var ret []*Token

	for {
		t := x.Token()
		ret = append(ret, t)
		if t.Type == EOF {
			break
		}
	}

	errs := x.Errs()
	if errs != nil {
		return nil, errs
	}

	return ret, nil
}
