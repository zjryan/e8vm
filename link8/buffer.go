package link8

import (
	"bytes"
)

// Buf is just a wrapper around bytes.Buffer
// that gives it a Close() method
type Buf struct {
	bytes.Buffer
}

// Close is an noop
func (b Buf) Close() error {
	return nil
}
