package arch8

import (
	"errors"
)

// Excep defines an exception error with a code
type Excep struct {
	Code byte
	Err  error
}

// NewExcep creates a new Exception with a particular code and message.
func NewExcep(c byte, s string) *Excep {
	ret := new(Excep)
	ret.Code = c
	ret.Err = errors.New(s)
	return ret
}

func (e *Excep) Error() string {
	return e.Err.Error()
}

var (
	errHalt        = NewExcep(1, "halt")
	errTimeInt     = NewExcep(2, "time interrupt")
	errInvalidInst = NewExcep(3, "invalid instruction")

	errOutOfRange   = NewExcep(4, "out of range")
	errMisalign     = NewExcep(5, "address misalign")
	errPageFault    = NewExcep(6, "page fault") // page invalid
	errPageReadonly = NewExcep(7, "page read-only")

	errSyscall = NewExcep(8, "system call")
)
