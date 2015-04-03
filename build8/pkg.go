package build8

import (
	"fmt"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	home Home
	path string
	src  string

	lang         Lang
	files        []string
	imports      map[string]*Import
	buildStarted bool

	compiled Linkable
	lib      *link8.Pkg

	err error
}

func newErrPkg(e error) *pkg { return &pkg{err: e} }

func newPkg(h Home, p string) *pkg {
	if !isPkgPath(p) {
		return newErrPkg(fmt.Errorf("invalid path: %q", p))
	}

	ret := new(pkg)
	ret.lang = h.Lang(p)
	if ret.lang == nil {
		return newErrPkg(fmt.Errorf("invalid pacakge: %q", p))
	} else if h.Src(p) == nil {
		return newErrPkg(fmt.Errorf("package has no source file: %q", p))
	}

	ret.home = h
	ret.path = p
	ret.imports = make(map[string]*Import)
	return ret
}

func (p *pkg) srcMap() map[string]*File { return p.home.Src(p.path) }

func (p *pkg) Import(name, path string, pos *lex8.Pos) {
	p.imports[name] = &Import{Path: path, Pos: pos}
}

var _ Importer = new(pkg)

/*
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
*/
