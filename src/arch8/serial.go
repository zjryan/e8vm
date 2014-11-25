package arch8

import (
	"bytes"
	"io"
	"log"
	"os"
)

// Serial is a serial console device
// It is basically two DMA pipes of ring buffered bytes:
// one pipe for input, one pipe for output.
type Serial struct {
	intBus IntBus
	p      *Page

	Core   byte
	IntIn  byte
	IntOut byte

	Output io.Writer
}

var _ Device = new(Serial) // Serial is a device

const (
	serialBase     = 0
	serialInHead   = serialBase + 0  // updated by hardware
	serialInTail   = serialBase + 4  // updated by cpu
	serialInWait   = serialBase + 8  // cycles to wait to raise an interrupt
	serialInThresh = serialBase + 12 // threashold to raise an input interrupt

	serialOutHead   = serialBase + 16 // updated by cpu
	serialOutTail   = serialBase + 20 // updated by hardware
	serialOutWait   = serialBase + 24 // cycles to wait to raise an interrupt
	serialOutThresh = serialBase + 28 // threashold to raise an output interrupt

	serialInBuf  = serialBase + 64
	serialOutBuf = serialBase + 96

	serialCap = 30 // 30 bytes maximum in each pipe
)

// NewSerial creates a new serial controller.
func NewSerial(p *Page, i IntBus) *Serial {
	ret := new(Serial)
	ret.intBus = i
	ret.p = p

	// default interrupts
	ret.Core = 0 // to core 0 only
	ret.IntIn = 8
	ret.IntOut = 9

	ret.Output = os.Stdout

	return ret
}

func (s *Serial) interrupt(code byte) {
	s.intBus.Interrupt(code, s.Core)
}

// WriteByte appends a byte into the input ring buffer.
func (s *Serial) WriteByte(b byte) bool {
	head := s.p.ReadWord(serialInHead)
	tail := s.p.ReadWord(serialInTail)
	n := head - tail
	if n >= serialCap {
		return false
	}

	s.p.WriteByte(serialInBuf+head%32, b)

	head++
	s.p.WriteWord(serialInTail, head)

	n++
	thresh := s.p.ReadWord(serialInThresh)
	if n >= thresh {
		s.interrupt(s.IntIn)
	}

	return true
}

// OutLen returns the current buffer length of the output ring buffer.
func (s *Serial) OutLen() uint32 {
	head := s.p.ReadWord(serialOutHead)
	tail := s.p.ReadWord(serialOutTail)
	ret := head - tail
	if ret > serialCap {
		log.Printf("error output buffer length")
		// on error set the length to full
		// to avoid trigger interrupt
		ret = serialCap
	}
	return ret
}

// InLen returns the current buffer length of the input ring buffer.
func (s *Serial) InLen() uint32 {
	head := s.p.ReadWord(serialInHead)
	tail := s.p.ReadWord(serialInTail)
	ret := head - tail
	if ret > serialCap {
		log.Printf("error input buffer length")
		// on error, set the length to empty
		// to avoid trigger interrupt
		ret = 0
	}
	return ret
}

// ReadByte reads a byte out of serial output ring buffer.
func (s *Serial) ReadByte() (byte, bool) {
	head := s.p.ReadWord(serialOutHead)
	tail := s.p.ReadWord(serialOutTail)
	n := head - tail
	if n == 0 || n > serialCap {
		return 0, false
	}

	b := s.p.ReadByte(serialOutBuf + tail%32)
	tail++
	s.p.WriteWord(serialOutTail, tail)

	n--
	thresh := s.p.ReadWord(serialOutThresh)
	if n <= thresh {
		s.interrupt(s.IntOut)
	}

	return b, true
}

func (s *Serial) countDown() {
	inWait := s.p.ReadWord(serialInWait)
	outWait := s.p.ReadWord(serialOutWait)

	if inWait > 0 {
		inWait--
	}
	if outWait > 0 {
		outWait--
	}

	if inWait == 0 {
		if s.InLen() > 0 {
			s.interrupt(s.IntIn)
		}
	}

	if outWait == 0 {
		if s.OutLen() < serialCap {
			s.interrupt(s.IntOut)
		}
	}

	s.p.WriteWord(serialInWait, inWait)
	s.p.WriteWord(serialOutWait, outWait)
}

func (s *Serial) flush() {
	buf := new(bytes.Buffer)
	for {
		b, valid := s.ReadByte()
		if valid {
			buf.WriteByte(b)
		} else {
			break
		}
	}

	if s.Output != nil {
		_, e := s.Output.Write(buf.Bytes())
		if e != nil {
			log.Print(e)
		}
	}
}

// Tick counts down the waiting counters and triggers
// interrupt when the count down reaches zero.
func (s *Serial) Tick() {
	s.flush()
	s.countDown()
}
