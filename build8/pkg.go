package build8

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/lex8"
)

type pkg struct {
	home *home
	path string

	src string

	imports map[string]*pkgImport
}

func newPkg(h *home, p string) (*pkg, error) {
	if !isPkgPath(p) {
		return nil, fmt.Errorf("invalid path: %q", p)
	}

	ret := new(pkg)
	ret.home = h
	ret.path = p
	ret.src = h.src(p)
	ret.imports = make(map[string]*pkgImport)

	ret.loadImport()

	return ret, nil
}

func (p *pkg) srcFile(f string) string {
	return filepath.Join(p.src, f)
}

func (p *pkg) openSrcFile(f string) io.ReadCloser {
	return newFile(p.srcFile(f))
}

func (p *pkg) loadImport() (*imports, []*lex8.Error) {
	path := p.srcFile("imports")
	return parseImports(path, newFile(path))
}

func (p *pkg) listSrcFiles(suffix string) ([]string, error) {
	dir, e := os.Open(p.src)
	if e != nil {
		return nil, e
	}

	files, e := dir.Readdir(0)
	if e != nil {
		return nil, e
	}

	e = dir.Close()
	if e != nil {
		return nil, e
	}

	var ret []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if strings.HasSuffix(name, suffix) {
			fullpath := p.srcFile(name)
			ret = append(ret, fullpath)
		}
	}

	return ret, nil
}
