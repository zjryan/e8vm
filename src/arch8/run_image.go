package arch8

import (
	"os"
)

// RunImage loads and run a raw image on a single core machine
// with 1GB physical memory until it runs into an exception.
func RunImage(image string) error {
	f, e := os.Open(image)
	if e != nil {
		return e
	}

	m := NewMachine(uint32(1<<30), 1)
	e = m.LoadImage(f, InitPC)
	if e != nil {
		return e
	}

	_, exp := m.Run(0)
	return exp
}
