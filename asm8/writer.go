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

func (w *writer) writeBareFunc(f *funcObj) {
	for _, i := range f.insts {
		w.writeU32(i)
	}
}

func (w *writer) bytes() []byte {
	return w.buf.Bytes()
}

func writeFunc(w *writer, s pkgSym) {
	p := s.p
	f := p.FuncObj(s.sym)

	cur := 0
	var curLink *link
	var curIndex int
	updateCur := func() {
		if cur < len(f.links) {
			curLink = f.links[cur]
			curIndex = int(curLink.offset >> 2)
		}
	}

	updateCur()
	for i, inst := range f.insts {
		if i == curIndex {
			fill := curLink.offset & 0x4
			if fill == fillLink {
				if (inst >> 31) != 0x1 {
					panic("not a jump")
				}
				if (inst & 0x3fffffff) != 0 {
					panic("already filled")
				}

				pc := f.addr + uint32(i)*4 + 4
				target := p.requires[curLink.pkg].FuncObj(curLink.sym).addr
				inst |= (target - pc) >> 2
			} else if fill == fillHigh || fill == fillLow {
				if (inst & 0xffff) != 0 {
					panic("already filled")
				}

				pkg := p.requires[curLink.pkg]
				sym := pkg.symbols[curLink.sym]
				var v uint32
				switch sym.Type {
				case SymFunc:
					v = pkg.FuncObj(curLink.sym).addr
				case SymVar:
					panic("todo")
				case SymConst:
					panic("todo")
				default:
					panic("bug")
				}

				if fill == fillHigh {
					inst |= v >> 16
				} else { // fillLow
					inst |= v & 0xffff
				}
			} else {
				panic("invalid fill")
			}

			cur++
			updateCur()
		}

		w.writeU32(inst)
	}
}
