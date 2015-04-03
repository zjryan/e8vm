package build8

import (
	"bytes"
	"sort"
)

type memFile struct {
	buf *bytes.Buffer
}

type memPkg struct {
	files map[string]*memFile
}

// MemHome is a memory based building home.
type MemHome struct {
	pkgs map[string]*memPkg
}

// Pkgs lists all the packages in this home
func (h *MemHome) Pkgs() []string {
	ret := make([]string, 0, len(h.pkgs))
	for p := range h.pkgs {
		ret = append(ret, p)
	}
	sort.Strings(ret)
	return ret
}

// Src lists all the source files in this home
