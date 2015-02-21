package build8

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	home *home
	path string

	src string

	lib *link8.Pkg
}

func newPkg(h *home, p string) (*pkg, error) {
	if !isPkgPath(p) {
		return nil, fmt.Errorf("invalid path: %q", p)
	}

	ret := new(pkg)
	ret.home = h
	ret.path = p
	ret.src = h.src(p)

	return ret, nil
}

func (p *pkg) srcFile(f string) string {
	return filepath.Join(p.src, f)
}

func (p *pkg) openSrcFile(f string) io.ReadCloser {
	return newFile(p.srcFile(f))
}

func (p *pkg) loadImports() (*imports, []*lex8.Error) {
	path := p.srcFile(importFile)
	_, e := os.Stat(path)
	if os.IsNotExist(e) {
		return nil, nil
	}

	f := newFile(path)
	ret, es := parseImports(path, f)
	e = f.Close()
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return ret, es
}

func (p *pkg) listSrcFiles(suffix string) ([]string, error) {
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
		if strings.HasSuffix(name, suffix) {
			fullpath := p.srcFile(name)
			ret = append(ret, fullpath)
		}
	}

	return ret, nil
}

func (p *pkg) openSrcFiles(suffix string) (map[string]io.ReadCloser, error) {
	files, e := p.listSrcFiles(suffix)
	if e != nil {
		return nil, e
	}

	ret := make(map[string]io.ReadCloser)
	for _, f := range files {
		ret[f] = newFile(f)
	}
	return ret, nil
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
		if isSrc(name) {
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

func (p *pkg) build(imps *imports) (*link8.Pkg, []*lex8.Error) {
	files, e := p.openSrcFiles(".s")
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	imports := make(map[string]*ast.PkgImport)
	if imps != nil {
		for as, imp := range imps.m {
			if imp.lib == nil {
				panic("no lib binded")
			}
			imports[as] = &ast.PkgImport{
				As:    as,
				Pkg:   imp.lib,
				Tok:   imp.pathToken,
				AsTok: imp.asToken,
			}
		}
	}

	pb := asm8.Pkg{
		Path:    p.path,
		Imports: imports,
		Files:   files,
	}

	lib, es := pb.Build()
	if es != nil {
		return nil, es
	}

	p.lib = lib

	return lib, nil
}
