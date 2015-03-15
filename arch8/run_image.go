package arch8

import (
	"os"

	"lonnie.io/e8vm/conv"
)

// RunImage loads and run a raw image on a single core machine
// with 1GB physical memory until it runs into an exception.
func RunImage(path string) error {
	f, e := os.Open(path)
	if e != nil {
		return e
	}

	m := NewMachine(uint32(1<<30), 1)
	e = m.LoadImage(f, conv.InitPC)
	if e != nil {
		return e
	}

	_, exp := m.Run(0)
	if exp != nil {
		return exp
	}

	return nil
}
