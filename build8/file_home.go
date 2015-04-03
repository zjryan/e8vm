package build8

import (
	"io"
	"path/filepath"
	"strings"
)

// FileHome is a file system basd building home.
type FileHome struct {
	path string

	defaultLang Lang
	langs       map[string]Lang
}

var _ Home = new(FileHome)

func (h *FileHome) sub(pre, p string) string {
	return filepath.Join(h.path, pre, p)
}

func (h *FileHome) subFile(pre, p, f string) string {
	return filepath.Join(h.path, pre, p, f)
}

// Pkgs lists all the packages inside this home folder.
func (h *FileHome) Pkgs() []string {
	panic("todo")
}

// Src lists all the source files inside this package.
func (h *FileHome) Src(p string) (map[string]*File, Lang) {
	lang := h.lang(p)
	_ = lang
	panic("todo")
}

// Bin returns the writer to write the binary
func (h *FileHome) Bin(p string) io.WriteCloser {
	return newFile(h.sub("bin", p+".e8"))
}

// Lib returns the writer to write the linkable library
func (h *FileHome) Lib(p string) io.WriteCloser {
	return newFile(h.sub("pkg", p+".e8a"))
}

// Log returns the log writer for the particular name
func (h *FileHome) Log(p, name string) io.WriteCloser {
	return newFile(h.subFile("src", p, name))
}

// Lang returns the language for the particular path.
// It searches for the longest prefix match
func (h *FileHome) lang(p string) Lang {
	nmax := -1
	var ret Lang
	for prefix, lang := range h.langs {
		n := len(prefix)
		if n < nmax || !strings.HasPrefix(p, prefix) {
			continue
		}

		nmax = n
		ret = lang
	}

	if ret == nil {
		ret = h.defaultLang
	}
	return ret
}

// MemHome is a memory based building home.
type MemHome struct {
}
