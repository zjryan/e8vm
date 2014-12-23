package lex8

import (
	"bufio"
	"io"
)

// RuneScanner is a rune scanner that scans runes from a file,
// and at the same time tracks the reading position.
type RuneScanner struct {
	file string
	line int

	rc io.ReadCloser
	r  *bufio.Reader

	Err  error // any error encountered
	Rune rune  // the rune just read
}

// NewRuneScanner creates a scanner.
func NewRuneScanner(file string, rc io.ReadCloser) *RuneScanner {
	ret := new(RuneScanner)
	ret.file = file
	ret.rc = rc
	ret.r = bufio.NewReader(rc)
	ret.line = 1 // natural counting

	return ret
}

// Scan reads in the next rune to s.Rune.
// It closes the reader automatically when it reaches the end of file
// or when an error occurs.
func (s *RuneScanner) Scan() bool {
	wasEndline := s.Rune == '\n'

	s.Rune, _, s.Err = s.r.ReadRune()

	if s.Err != nil {
		closeErr := s.rc.Close()
		if s.Err == io.EOF {
			s.Err = closeErr
		}
		return false
	}

	if wasEndline {
		s.line++
	}

	return true
}

// Pos returns the current position in the file.
func (s *RuneScanner) Pos() *Pos {
	return &Pos{s.file, s.line}
}
