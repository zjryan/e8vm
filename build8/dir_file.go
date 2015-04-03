package build8

import (
	"os"
	"path/filepath"
)

type dirFile struct {
	path string
	f    *os.File
}

func newDirFile(p string) *dirFile {
	ret := new(dirFile)
	ret.path = p
	return ret
}

// Read will open the file for reading on
// its first read
func (f *dirFile) Read(buf []byte) (int, error) {
	if f.f == nil {
		var e error
		f.f, e = os.Open(f.path)
		if e != nil {
			return 0, e
		}
	}

	return f.f.Read(buf)
}

// Write will open the file for writing on
// its first write.
func (f *dirFile) Write(buf []byte) (int, error) {
	if f.f == nil {
		var e error
		folder := filepath.Dir(f.path)
		if folder != "." {
			e = os.MkdirAll(folder, 0755)
			if e != nil {
				return 0, e
			}
		}

		f.f, e = os.Create(f.path)
		if e != nil {
			return 0, e
		}
	}

	return f.f.Write(buf)
}

// Close will close the file if the file
// has already been opened.
func (f *dirFile) Close() error {
	if f.f == nil {
		return nil
	}
	return f.f.Close()
}
