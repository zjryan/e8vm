package lex8

import (
	"fmt"
	"io"
)

// ErrorList saves a list of error
type ErrorList struct {
	errs []*Error

	Max int

	inJail bool
}

var _ Logger = new(ErrorList)

// NewErrorList creates a new error list with default (20) maximum
// lines of errors.
func NewErrorList() *ErrorList {
	ret := new(ErrorList)
	ret.Max = 20

	return ret
}

// Add appends the error to the list. Change the state to "in jail".
func (lst *ErrorList) Add(e *Error) {
	if e == nil {
		panic("nil error")
	}

	lst.inJail = true
	if len(lst.errs) >= lst.Max {
		return
	}

	lst.errs = append(lst.errs, e)
}

// Jail puts it in jail without generating a new error message
func (lst *ErrorList) Jail() { lst.inJail = true }

// InJail checks if a new error has been added since created or last bail out
func (lst *ErrorList) InJail() bool { return lst.inJail }

// BailOut clears the "in jail" state.
func (lst *ErrorList) BailOut() { lst.inJail = false }

// Errorf appends a new error with particular position and format.
func (lst *ErrorList) Errorf(p *Pos, f string, args ...interface{}) {
	lst.Add(&Error{p, fmt.Errorf(f, args...)})
}

// Print prints to the writer (maximume lst.MaxPrint errors).
func (lst *ErrorList) Print(w io.Writer) error {
	for _, e := range lst.errs {
		_, pe := fmt.Fprintln(w, e)
		if pe != nil {
			return pe
		}
	}

	return nil
}

// SingleErr returns an error array with one error.
func SingleErr(e error) []*Error {
	return []*Error{{Err: e}}
}

// Errs retunrs the errors in the list
func (lst *ErrorList) Errs() []*Error {
	ret := lst.errs
	if len(ret) == 0 {
		return nil
	}
	return ret
}
