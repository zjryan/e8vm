package build8

import (
	"bytes"
)

type memFile struct{ *bytes.Buffer }

func newMemFile() *memFile { return &memFile{new(bytes.Buffer)} }
func (f *memFile) Reader() *memFileReader {
	return &memFileReader{
		bytes.NewBuffer(f.Buffer.Bytes()),
	}
}
func (f *memFile) Close() error { return nil }

type memFileReader struct{ *bytes.Buffer }

func (f *memFileReader) Close() error { return nil }
