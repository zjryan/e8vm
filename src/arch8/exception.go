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
	errTimeInt      = NewExcep(1, "time interrupt")
	errOutOfRange   = NewExcep(2, "out of range")
	errMisalign     = NewExcep(3, "address misalign")
	errPageFault    = NewExcep(4, "page fault") // page invalid
	errPageReadonly = NewExcep(5, "page read-only")
)
