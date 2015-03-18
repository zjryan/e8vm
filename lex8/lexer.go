package lex8

import (
	"io"
)

// LexFunc is a function type that takes a lexer and returns the next token.
type LexFunc func(x *Lexer) *Token

// WhiteFunc is a function type that checks if a rune is white space.
type WhiteFunc func(r rune) bool

// Lexer parses a file input stream into tokens.
type Lexer struct {
	s *LexScanner

	e    error
	errs *ErrorList

	r       rune

	IsWhite WhiteFunc
	LexFunc
}

// NewLexer creates a new lexer.
func NewLexer(file string, r io.Reader) *Lexer {
	ret := new(Lexer)
	ret.s = NewLexScanner(file, r)
	ret.errs = NewErrList()

	ret.IsWhite = func(r rune) bool {
		return r == ' ' || r == '\t'
	}

	ret.Next()

	return ret
}

// Next pushes the current rune into the scanning buffer,
// and returns the next rune.
func (x *Lexer) Next() (rune, error) {
	x.r, x.e = x.s.Next()
	return x.r, x.e
}

// Rune returns the current rune.
func (x *Lexer) Rune() rune {
	return x.r
}

// See returns true when the current rune is r.
func (x *Lexer) See(r rune) bool {
	return x.r == r
}

// Discard clears the scanning buffer
func (x *Lexer) Discard() { x.s.Accept() }

// Ended returns true when the scanning stops.
func (x *Lexer) Ended() bool { return x.e != nil }

// SkipWhite is a helper function that skips
// any rune that returns true by IsWhite function.
// The buffer is discarded after the skipping.
func (x *Lexer) SkipWhite() {
	for {
		if x.Ended() || !x.IsWhite(x.r) {
			break
		}
		x.Next()
	}
	x.Discard()
}

// MakeToken accepts the runes in the scanning buffer
// and returns it as a token of type t.
func (x *Lexer) MakeToken(t int) *Token {
	s, p := x.s.Accept()
	return &Token{t, s, p}
}

// Token returns the next parsed token.
// It ends with a token with type EOF.
func (x *Lexer) Token() *Token {
	x.SkipWhite()

	if x.Ended() {
		return x.MakeToken(EOF)
	}

	if x.LexFunc == nil {
		x.Next()
		return x.MakeToken(Illegal)
	}

	return x.LexFunc(x)
}

// Err adds an error into the error list with current postion.
func (x *Lexer) Err(f string, args ...interface{}) {
	x.errs.Addf(x.s.Pos(), f, args...)
}

// Errs returns the lexing errors.
func (x *Lexer) Errs() []*Error {
	if x.e != nil && x.e != io.EOF {
		return []*Error{{Err: x.e}}
	}

	return x.errs.Errs
}
