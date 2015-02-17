package build8

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/lex8"
)

func isPkgName(s string) bool {
	if s == "" {
		return false
	}

	for i, r := range s {
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		return false
	}

	return true
}

func isPkgPath(p string) bool {
	if p == "" {
		return false
	}
	subs := strings.Split(p, "/")
	for _, sub := range subs {
		if !isPkgName(sub) {
			return false
		}
	}
	return true
}

type pkg struct {
	home *home
	path string

	src string

	imports map[string]*pkg
}

func newPkg(h *home, p string) (*pkg, error) {
	if !isPkgPath(p) {
		return nil, fmt.Errorf("invalid path: %q", p)
	}

	ret := new(pkg)
	ret.home = h
	ret.path = p
	ret.src = h.src(p)
	ret.imports = make(map[string]*pkg)

	ret.loadImport()

	return ret, nil
}

func (p *pkg) srcFile(f string) string {
	return filepath.Join(p.src, f)
}

func (p *pkg) openSrcFile(f string) io.ReadCloser {
	return newFile(p.srcFile(f))
}

func (p *pkg) loadImport() []*lex8.Error {
	path := p.srcFile("imports")
	_, es := parseImports(path, newFile(path))
	if es != nil {
		return es
	}
	// p.imports = imports
	return nil
}
