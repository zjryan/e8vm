package lex8

import (
	"fmt"
)

// Pos is the file line position in a file
type Pos struct {
	File string
	Line int
	Col  int
}

func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d:%d", p.File, p.Line, p.Col)
}
