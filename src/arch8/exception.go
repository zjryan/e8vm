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
func NewExcep(c byte, s string) *Excep {
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
	errHalt        = NewExcep(ErrHalt, "halt")
	errTimeInt     = NewExcep(ErrTimer, "time interrupt")
	errInvalidInst = NewExcep(ErrInvalidInst, "invalid instruction")

	errOutOfRange = NewExcep(ErrOutOfRange, "out of range")
	errMisalign   = NewExcep(ErrMisalign, "address misalign")
)

func newPageFault(va uint32) *Excep {
	ret := NewExcep(ErrPageFault, "page fault")
	ret.Arg = va
	return ret
}

func newPageReadonly(va uint32) *Excep {
	ret := NewExcep(ErrPageReadonly, "page read-only")
	ret.Arg = va
	return ret
}
