package build8

import (
	"fmt"
	"io/ioutil"
	// "os"
	"path/filepath"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	home *home
	path string
	src  string

	lang         Lang
	files        []string
	imports      map[string]*Import
	err          error // an error package, required but not exists
	buildStarted bool

	compiled Linkable
	lib      *link8.Pkg
}

func newPkg(h *home, p string, lang Lang) *pkg {
	ret := new(pkg)

	if !isPkgPath(p) {
		ret.err = fmt.Errorf("invalid path: %q", p)
		return ret
	}

	ret.home = h
	ret.path = p
	ret.lang = lang
	ret.src = h.src(p)

	var e error
	ret.files, e = listSrcFiles(ret.src, lang)
	if e != nil {
		ret.err = e
		return ret
	}
	if len(ret.files) == 0 {
		ret.err = fmt.Errorf("empty package: %q", p)
		return ret
	}

	ret.imports = make(map[string]*Import)
	return ret
}

func (p *pkg) srcFilePath(f string) string { return filepath.Join(p.src, f) }

func listSrcFiles(dir string, lang Lang) ([]string, error) {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}

	var ret []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if lang.IsSrc(name) {
			ret = append(ret, name)
		}
	}

	return ret, nil
}

func (p *pkg) srcMap() map[string]*File {
	ret := make(map[string]*File)

	for _, f := range p.files {
		path := p.srcFilePath(f)
		file := &File{
			Name:       f,
			Path:       path,
			ReadCloser: newFile(path),
		}

		ret[f] = file
	}
	return ret
}

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
