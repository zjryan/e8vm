package arch8

import (
	"errors"
)

type Excep struct {
	Code uint32
	Err  error
}

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
	errMisalign     = NewExcep(4, "misalign")
	errPageFault    = NewExcep(3, "page fault")
	errPageReadonly = NewExcep(4, "page read-only")
)
