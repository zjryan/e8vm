package build8

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/pkg8"
)

type pkg struct {
	home *home
	path string
	src  string

	lang  pkg8.Lang
	files []string

	built pkg8.Linkable
	lib   *link8.Pkg
}

func newPkg(h *home, p string, lang pkg8.Lang) (*pkg, error) {
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
			ret = append(ret, p.srcFile(name))
		}
	}

	return ret, nil
}

func (p *pkg) srcFiles() pkg8.Files {
	ret := make(map[string]io.ReadCloser)
	for _, f := range p.files {
		ret[f] = newFile(f)
	}
	return pkg8.Files(ret)
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
