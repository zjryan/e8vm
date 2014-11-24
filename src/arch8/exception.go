package arch8

import (
	"errors"
)

// Excep defines an exception error with a code
type Excep struct {
	Code uint32
	Err  error
}

// NewExcep creates a new Exception with a particular code and message.
func NewExcep(c uint32, s string) *Excep {
	ret := new(Excep)
	ret.Code = c
	ret.Err = errors.New(s)
	return ret
}

func (e *Excep) Error() string {
	return e.Err.Error()
}

var (
	errOutOfRange   = NewExcep(1, "out of range")
	errMisalign     = NewExcep(2, "address misalign")
	errPageFault    = NewExcep(3, "page fault") // page invalid
	errPageReadonly = NewExcep(4, "page read-only")
)
