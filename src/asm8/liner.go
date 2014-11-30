package asm8

import (
	"bufio"
	"io"
)

// Pos specifies a position in a file.
type Pos struct {
	File string
	Row  int
}

// Line specifies a raw line in a file
type Line struct {
	Pos  *Pos
	Text string
}

// Lines parses a file into lines.
func Lines(file string, r io.ReadCloser) ([]*Line, error) {
	s := bufio.NewScanner(r)
	var ret []*Line
	n := 0

	for s.Scan() {
		text := s.Text()
		n++
		line := &Line{
			&Pos{file, n},
			text,
		}

		ret = append(ret, line)
	}

	e := s.Err()
	if e != nil {
		return nil, e
	}

	return ret, nil
}
