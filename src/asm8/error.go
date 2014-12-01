package asm8

import (
	"fmt"
)

// Error is a parsing error
type Error struct {
	Pos *Pos  // Pos can be null for error not related to any position
	Err error // Err is the error message
}

// Error returns the error string.
func (e *Error) Error() string {
	if e.Pos == nil {
		return e.Err.Error()
	}

	return fmt.Sprintf("%s:%d: %s",
		e.Pos.File, e.Pos.Line,
		e.Err.Error(),
	)
}
