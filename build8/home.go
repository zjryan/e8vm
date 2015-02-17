package build8

import (
	"io"
	"path/filepath"
)

type home struct {
	path string
}

func (h *home) sub(pre, p string) string {
	return filepath.Join(h.path, pre, p)
}

func (h *home) src(p string) string { return h.sub("src", p) }
func (h *home) bin(p string) string { return h.sub("bin", p+".e8") }
func (h *home) pkg(p string) string { return h.sub("pkg", p+".e8a") }

func (h *home) makeBin(p string) io.WriteCloser {
	return newFile(h.bin(p))
}

func (h *home) makePkg(p string) io.WriteCloser {
	return newFile(h.pkg(p))
}

func (h *home) loadPkg(p string) io.ReadCloser {
	return newFile(h.pkg(p))
}
