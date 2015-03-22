package build8

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	home *home
	path string
	src  string

	lang    Lang
	files   []string
	imports map[string]*Import

	compiled Linkable
	lib      *link8.Pkg
}

var _ Pkg = new(pkg)

func newPkg(h *home, p string, lang Lang) (*pkg, error) {
	if !isPkgPath(p) {
		return nil, fmt.Errorf("invalid path: %q", p)
	}

	ret := new(pkg)
	ret.home = h
	ret.path = p
	ret.lang = lang
	ret.src = h.src(p)

	var e error
	ret.files, e = ret.listSrcFiles()
	if e != nil {
		return nil, e
	}

	ret.imports = make(map[string]*Import)

	return ret, nil
}

func (p *pkg) srcFile(f string) string { return filepath.Join(p.src, f) }
func (p *pkg) openSrcFile(f string) io.ReadCloser {
	return newFile(p.srcFile(f))
}

func (p *pkg) listSrcFiles() ([]string, error) {
	files, e := ioutil.ReadDir(p.src)
	if e != nil {
		return nil, e
	}

	var ret []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if p.lang.IsSrc(name) {
			ret = append(ret, name)
		}
	}

	return ret, nil
}

func (p *pkg) Path() string { return p.path }

// Src returns the map of the source files
func (p *pkg) Src() map[string]*File {
	ret := make(map[string]*File)

	for _, f := range p.files {
		path := p.srcFile(f)
		file := &File{
			Name:       f,
			Path:       path,
			ReadCloser: newFile(path),
		}

		ret[f] = file
	}
	return ret
}

func (p *pkg) AddImport(name, path string, pos *lex8.Pos) {
	p.imports[name] = &Import{Path: path, Pos: pos}
}

func (p *pkg) Imports() map[string]*Import {
	return p.imports
}

func (p *pkg) SetCompiled(lib Linkable) {
	p.compiled = lib
}

func (p *pkg) Compiled() Linkable {
	return p.compiled
}

func (p *pkg) lastUpdate(suffix string) (*timeStamp, error) {
	ts := new(timeStamp)

	dirInfo, e := os.Stat(p.src)
	if e != nil {
		return nil, e
	}
	ts.update(dirInfo.ModTime())

	files, e := ioutil.ReadDir(p.src)
	if e != nil {
		return nil, e
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if p.lang.IsSrc(name) {
			ts.update(file.ModTime())
		}
	}

	return ts, nil
}

func (p *pkg) lastBuild() (*timeStamp, error) {
	ts := new(timeStamp)

	info, e := os.Stat(p.home.pkg(p.path))
	if !os.IsNotExist(e) {
		return nil, e
	}
	ts.update(info.ModTime())

	info, e = os.Stat(p.home.bin(p.path))
	if !os.IsNotExist(e) {
		return nil, e
	}
	ts.update(info.ModTime())

	return ts, nil
}
