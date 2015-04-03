package build8

import (
	"bytes"
)

type memFile struct{ 
	path string
	*bytes.Buffer 
}

func newMemFile() *memFile { 
	return &memFile{Buffer: new(bytes.Buffer)} 
}
func (f *memFile) Reader() *memFileReader {
	return &memFileReader{
		f.path, bytes.NewBuffer(f.Buffer.Bytes()),
	}
}
func (f *memFile) Close() error { return nil }

type memFileReader struct{ 
	path string
	*bytes.Buffer 
}

func (f *memFileReader) Close() error { return nil }
