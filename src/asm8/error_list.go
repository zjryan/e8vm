package asm8

import (
	"fmt"
	"io"
)

// ErrorList saves a list of error
type ErrorList struct {
	Errs []*Error

	Max int
}

// NewErrList creates a new error list with default (20) maximum
// lines of errors.
func NewErrList() *ErrorList {
	ret := new(ErrorList)
	ret.Max = 20

	return ret
}

// Add appends the error to the list
func (lst *ErrorList) Add(e *Error) {
	if len(lst.Errs) >= lst.Max {
		return
	}

	lst.Errs = append(lst.Errs, e)
}

// Addf appends a new error with particular position and format.
func (lst *ErrorList) Addf(p *Pos, f string, args ...interface{}) {
	lst.Add(&Error{p, fmt.Errorf(f, args...)})
}

// Print prints to the writer (maximume lst.MaxPrint errors).
func (lst *ErrorList) Print(w io.Writer) error {
	for _, e := range lst.Errs {
		_, pe := fmt.Fprintln(w, e)
		if pe != nil {
			return pe
		}
	}

	return nil
}
