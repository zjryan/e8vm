// Package conv defines the common conventions of the architecture
package conv

import (
	"encoding/binary"
)

const (
	// InitPC points the default starting program counter
	InitPC = 0x8000
)

// The machine's endian (byte order).
var Endian = binary.LittleEndian
