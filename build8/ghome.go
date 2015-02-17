package build8

import (
	"io"
	"path/filepath"
)

type ghome struct {
	path string
}

func (h *ghome) sub(pre, p string) string {
	return filepath.Join(h.path, pre, p)
}

func (h *ghome) src(p string) string { return h.sub("src", p) }
func (h *ghome) bin(p string) string { return h.sub("bin", p+".e8") }
func (h *ghome) pkg(p string) string { return h.sub("pkg", p+".e8a") }

func (h *ghome) makeBin(p string) io.WriteCloser {
	return newFile(h.bin(p))
}

func (h *ghome) makePkg(p string) io.WriteCloser {
	return newFile(h.pkg(p))
}

func (h *ghome) loadPkg(p string) io.ReadCloser {
	return newFile(h.pkg(p))
}
