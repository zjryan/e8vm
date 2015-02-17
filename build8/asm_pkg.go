package build8

import (
	"io"
	"time"

	"lonnie.io/e8vm/asm8"
)

type asmPkg struct {
	gpath  string // gpath, package archive root
	path   string // package path
	folder string

	importFile io.ReadCloser
	files      map[string]io.ReadCloser
	lastUpdate time.Time
	lastBuild  *time.Time
	imports    []string
}

func (p *asmPkg) build() *asm8.PkgBuild {
	return &asm8.PkgBuild{
		Path:   p.path,
		Import: nil, // TODO
		Files:  p.files,
	}
}
