package asm8

import (
	"fmt"
	"io"
)

// Lexer parses the a file input stream into tokens
type Lexer struct {
	s *LexScanner

	r    rune
	e    error
	errs *ErrList
}

// NewLexer creates a new lexer of a file stream.
func NewLexer(file string, r io.ReadCloser) *Lexer {
	ret := new(Lexer)
	ret.s = NewLexScanner(file, r)
	ret.errs = NewErrList()

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

func isWhite(r rune) bool { return r == ' ' || r == '\t' }

func (x *Lexer) skipWhite() {
	for {
		if x.eof() || !isWhite(x.r) {
			break
		}
		x.next()
	}
	x.discard()
}

func (x *Lexer) scanString() *Token {
	escaped := false

	for {
		x.next()
		if x.eof() {
			x.err("unexpected eof in string")
			return x.token(String)
		} else if x.r == '\n' {
			x.err("unexpected endl in string")
			return x.token(String)
		}

		if escaped {
			escaped = false
		} else {
			if x.r == '\\' {
				escaped = true
			} else if x.r == '"' {
				x.next()
				break
			}
		}
	}

	return x.token(String)
}

func (x *Lexer) scanLineComment() *Token {
	for {
		x.next()
		if x.eof() || x.r == '\n' {
			break
		}
	}
	return x.token(Comment)
}

func (x *Lexer) scanBlockComment() *Token {
	star := false
	for {
		x.next()
		if x.eof() {
			x.err("unexpected eof in block comment")
			return x.token(Comment)
		}

		if star && x.r == '/' {
			x.next()
			break
		}

		star = x.r == '*'
	}

	return x.token(Comment)
}

func isOperandChar(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}

	for _, x := range []rune{'_', '+', '-', '.', ':'} {
		if r == x {
			return true
		}
	}

	return false
}

func isKeyword(lit string) bool {
	switch lit {
	case "func", "var", "const", "import":
		return true
	}
	return false
}

func (x *Lexer) scanOperand() *Token {
	for {
		x.next()
		if x.eof() || !isOperandChar(x.r) {
			break
		}
	}

	ret := x.token(Operand)
	if isKeyword(ret.Lit) {
		ret.Type = Keyword
	}
	return ret
}

// Token returns the next parsed token.
// It ends with a token with type EOF.
func (x *Lexer) Token() *Token {
	x.skipWhite()

	if x.eof() {
		return x.token(EOF)
	}

	switch x.r {
	case '\n':
		x.next()
		return x.token(Endl)
	case '{':
		x.next()
		return x.token(Lbrace)
	case '}':
		x.next()
		return x.token(Rbrace)
	case '/':
		x.next()
		if x.r == '/' {
			return x.scanLineComment()
		} else if x.r == '*' {
			return x.scanBlockComment()
		}
		x.err("illegal char %q", x.r)
		return x.token(Illegal)
	case '"':
		return x.scanString()
	}

	if isOperandChar(x.r) {
		return x.scanOperand()
	}

	x.err("illegal char %q", x.r)
	x.next()
	return x.token(Illegal)
}

func (x *Lexer) err(f string, args ...interface{}) {
	x.errs.Add(&Error{x.s.Pos(), fmt.Errorf(f, args...)})
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
