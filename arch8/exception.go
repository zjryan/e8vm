package arch8

import (
	"errors"
)

// Excep defines an exception error with a code
type Excep struct {
	Code byte
	Arg  uint32
	Err  error
}

// NewExcep creates a new Exception with a particular code and message.
func newExcep(c byte, s string) *Excep {
	ret := new(Excep)
	ret.Code = c
	ret.Err = errors.New(s)
	return ret
}

func (e *Excep) Error() string {
	return e.Err.Error()
}

// Exception codes
const (
	ErrHalt         = 1
	ErrTimer        = 2
	ErrInvalidInst  = 3
	ErrOutOfRange   = 4
	ErrMisalign     = 5
	ErrPageFault    = 6
	ErrPageReadonly = 7
)

var (
	errHalt        = newExcep(ErrHalt, "halt")
	errTimeInt     = newExcep(ErrTimer, "time interrupt")
	errInvalidInst = newExcep(ErrInvalidInst, "invalid instruction")

	errOutOfRange = newExcep(ErrOutOfRange, "out of range")
	errMisalign   = newExcep(ErrMisalign, "address misalign")
)

func newPageFault(va uint32) *Excep {
	ret := newExcep(ErrPageFault, "page fault")
	ret.Arg = va
	return ret
}

func newPageReadonly(va uint32) *Excep {
	ret := newExcep(ErrPageReadonly, "page read-only")
	ret.Arg = va
	return ret
}
