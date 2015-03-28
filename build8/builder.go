package build8

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Builder builds a bunch of packages.
type Builder struct {
	home *home
	pkgs map[string]*pkg

	// TODO: built should be something like an LRU cache
	// the libraries should be load back in only when linking

	lang    Lang
	Verbose bool
}

// NewBuilder creates a new builder
func NewBuilder(homePath string) *Builder {
	ret := new(Builder)
	ret.home = &home{path: homePath}
	ret.pkgs = make(map[string]*pkg)

	return ret
}

// AddLang registers a langauge into the building system
func (b *Builder) AddLang(lang Lang)     { b.lang = lang }
func (b *Builder) getLang(p string) Lang { return b.lang }

func (b *Builder) prepare(p string) (*pkg, []*lex8.Error) {
	saved := b.pkgs[p]
	if saved != nil {
		return saved, nil // already prepared
	}

	lang := b.getLang(p)
	pkg := newPkg(b.home, p, lang)
	b.pkgs[p] = pkg

	if pkg.err != nil {
		return pkg, nil
	}

	es := lang.Import(pkg)
	if es != nil {
		return pkg, es
	}

	for _, imp := range pkg.imports {
		impPkg, es := b.prepare(imp.Path)
		if es != nil {
			return pkg, es
		}

		if impPkg.err != nil {
			return pkg, []*lex8.Error{{
				Pos: imp.Pos,
				Err: impPkg.err,
			}}
		}
	}

	return pkg, nil
}

func (b *Builder) build(p string) (*pkg, []*lex8.Error) {
	ret := b.pkgs[p]
	if ret == nil {
		panic("build without prepare")
	}

	// already compiled
	if ret.compiled != nil {
		return ret, nil
	}

	if ret.buildStarted {
		e := fmt.Errorf("package %q circular depends itself", p)
		return ret, lex8.SingleErr(e)
	}

	ret.buildStarted = true
	lang := ret.lang

	for _, imp := range ret.imports {
		built, es := b.build(imp.Path)
		if es != nil {
			return nil, es
		}
		imp.Pkg = built
	}

	// ready to build this one
	if b.Verbose {
		fmt.Println(p)
	}

	// compile now
	es := lang.Compile(ret)
	if es != nil {
		return nil, es
	}

	lib := ret.Compiled().Lib() // the linkable lib
	// a package with main entrance, build the bin
	if lib.HasFunc("main") {
		log := lex8.NewErrorList()

		fout := b.home.makeBin(p)
		lex8.LogError(log, link8.LinkMain(lib, fout))
		lex8.LogError(log, fout.Close())

		if es := log.Errs(); es != nil {
			return nil, es
		}
	}

	return ret, nil
}

// Build builds a package
func (b *Builder) Build(p string) []*lex8.Error {
	_, es := b.prepare(p)
	if es != nil {
		return es
	}

	_, es = b.build(p)
	return es
}

// ListPkgs list all the packages under a home folder
func (b *Builder) ListPkgs() ([]string, error) {
	h := b.home

	src := h.srcRoot()

	m := make(map[string]struct{})

	e := filepath.Walk(src, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		name := info.Name()
		if info.IsDir() {
			if !lex8.IsPkgName(name) {
				return filepath.SkipDir
			}
		} else {
			path, e := filepath.Rel(src, filepath.Dir(p))
			lang := b.getLang(path)
			if lang.IsSrc(name) {
				if e != nil {
					panic(e)
				}
				m[path] = struct{}{}
			}
		}

		return nil
	})

	if e != nil {
		return nil, e
	}

	var ret []string
	for p := range m {
		ret = append(ret, p)
	}

	sort.Strings(ret)

	return ret, nil
}
