package lex8

import (
	"bytes"
	"io"
)

// LexScanner parses a file input stream into tokens.
type LexScanner struct {
	s     *RuneScanner
	errs  *ErrorList
	valid bool

	pos *Pos
	buf *bytes.Buffer
}

// NewLexScanner creates a new lexer.
func NewLexScanner(file string, r io.Reader) *LexScanner {
	ret := new(LexScanner)
	ret.s = NewRuneScanner(file, r)
	ret.errs = NewErrorList()

	ret.buf = new(bytes.Buffer)
	ret.pos = ret.s.Pos()

	return ret
}

// Next pushes the current rune (if valid) into the buffer,
// and returns the next rune or error from scanning the input
// stream.
func (s *LexScanner) Next() (rune, error) {
	if s.valid {
		s.buf.WriteRune(s.s.Rune) // push into the buffer
		s.valid = false
	}

	if !s.s.Scan() {
		if s.s.Err != nil {
			return 0, s.s.Err
		}

		return 0, io.EOF // signal end of file
	}

	s.valid = true
	return s.s.Rune, nil
}

// Accept returns the string buffered, and the starting position
// of the string.
func (s *LexScanner) Accept() (string, *Pos) {
	ret := s.buf.String()
	s.buf.Reset()
	pos := s.pos

	s.pos = s.s.Pos()

	return ret, pos
}

// Buffered returns the current buffered string in the
// scanner
func (s *LexScanner) Buffered() string {
	return s.buf.String()
}

// Pos returns the position of the buffer start.
func (s *LexScanner) Pos() *Pos {
	return s.pos
}
