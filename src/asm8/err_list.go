package asm8

import (
	"fmt"
	"io"
)

// ErrList saves a list of error
type ErrList struct {
	Errs []*Error

	Max int
}

// NewErrList creates a new error list with default (20) maximum
// lines of errors.
func NewErrList() *ErrList {
	ret := new(ErrList)
	ret.Max = 20

	return ret
}

// Add appends the error to the list
func (lst *ErrList) Add(e *Error) {
	if len(lst.Errs) >= lst.Max {
		return
	}

	lst.Errs = append(lst.Errs, e)
}

// Addf appends a new error with particular position and format.
func (lst *ErrList) Addf(p *Pos, f string, args ...interface{}) {
	lst.Add(&Error{p, fmt.Errorf(f, args...)})
}

// Print prints to the writer (maximume lst.MaxPrint errors).
func (lst *ErrList) Print(w io.Writer) error {
	for _, e := range lst.Errs {
		_, pe := fmt.Fprintln(w, e)
		if pe != nil {
			return pe
		}
	}

	return nil
}
