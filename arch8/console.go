package arch8

import (
	"io"
	"log"
	"os"
)

// Console is a simple console that can output/input a single
// byte at a time
type console struct {
	intBus intBus
	p      *page

	Core   byte
	IntIn  byte
	IntOut byte

	Output io.Writer
}

// NewConsole creates a new simple console.
func newConsole(p *page, i intBus) *console {
	ret := new(console)
	ret.intBus = i
	ret.p = p

	ret.Core = 0
	ret.IntIn = 8
	ret.IntOut = 9

	ret.Output = os.Stdout
	return ret
}

var _ device = new(console)

const (
	consoleBase = 0

	consoleOut      = consoleBase + 0
	consoleOutValid = consoleBase + 1
	consoleIn       = consoleBase + 2
	consoleInValid  = consoleBase + 3
)

func (c *console) interrupt(code byte) {
	c.intBus.Interrupt(code, c.Core)
}

// Tick flushes the buffered byte to the console.
func (c *console) Tick() {
	outValid := c.p.ReadByte(consoleOutValid)
	if outValid != 0 {
		out := c.p.ReadByte(consoleOut)
		_, e := c.Output.Write([]byte{out})
		if e != nil {
			log.Print(e)
		}
		c.p.WriteByte(consoleOutValid, 0)
		c.interrupt(c.IntOut) // out available
	}

	// TODO: input part
}
