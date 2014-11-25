package arch8

import (
	"os"
)

// RunRaw loads and run a raw image until it runs into an exception.
func RunRaw(image string) error {
	f, e := os.Open(image)
	if e != nil {
		return e
	}

	m := NewMachine(uint32(1<<30), 1)
	e = m.LoadImage(f, InitPC)
	if e != nil {
		return e
	}

	_, e = m.Run(0)
	return e
}
