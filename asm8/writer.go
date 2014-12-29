package asm8

import (
	"bytes"

	"encoding/binary"
)

type writer struct {
	buf *bytes.Buffer
}

func newWriter() *writer {
	ret := new(writer)
	ret.buf = new(bytes.Buffer)
	return ret
}

func (w *writer) writeU32(u uint32) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], u)
	_, e := w.buf.Write(b[:])
	if e != nil {
		panic("buf write")
	}
}

func (w *writer) writeFunc(f *funcObj) {
	for _, i := range f.insts {
		w.writeU32(i)
	}
}

func (w *writer) bytes() []byte {
	return w.buf.Bytes()
}
