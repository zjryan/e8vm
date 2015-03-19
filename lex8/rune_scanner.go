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
	col  int

	r *bufio.Reader

	Err  error // any error encountered
	Rune rune  // the rune just read

	closed bool
}

// NewRuneScanner creates a scanner.
func NewRuneScanner(file string, r io.Reader) *RuneScanner {
	ret := new(RuneScanner)
	ret.file = file
	ret.r = bufio.NewReader(r)
	ret.line = 1 // natural counting
	ret.col = 1

	return ret
}

// Scan reads in the next rune to s.Rune.
// It closes the reader automatically when it reaches the end of file
// or when an error occurs.
func (s *RuneScanner) Scan() bool {
	if s.closed {
		panic("scanning on closed rune scanner")
	}

	wasEndline := s.Rune == '\n'

	s.Rune, _, s.Err = s.r.ReadRune()

	if s.Err != nil {
		// closeErr := s.rc.Close()
		if s.Err == io.EOF {
			s.Err = nil
		}
		s.closed = true
		return false
	}

	if wasEndline {
		s.line++
		s.col = 1
	} else {
		s.col++
	}

	return true
}

// Pos returns the current position in the file.
func (s *RuneScanner) Pos() *Pos {
	return &Pos{s.file, s.line, s.col}
}
