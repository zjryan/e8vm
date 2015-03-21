package build8

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/pkg8"
)

// Builder builds a bunch of packages.
type Builder struct {
	home  *home
	built map[string]*pkg
	// TODO: built should be something like an LRU cache
	// the libraries should be load back in only when linking

	lang    pkg8.Lang
	Verbose bool
}

// NewBuilder creates a new builder
func NewBuilder(homePath string) *Builder {
	ret := new(Builder)
	ret.home = &home{path: homePath}
	ret.built = make(map[string]*pkg)

	ret.AddLang(asm8.Lang)

	return ret
}

// AddLang registers a langauge into the building system
func (b *Builder) AddLang(lang pkg8.Lang) {
	b.lang = lang
}

func (b *Builder) getLang(p string) pkg8.Lang {
	// TODO:
	return b.lang
}

func (b *Builder) find(p string) (*pkg, bool) {
	ret, found := b.built[p]
	return ret, found
}

func (b *Builder) register(p string)       { b.built[p] = nil }
func (b *Builder) save(p string, pkg *pkg) { b.built[p] = pkg }

func (b *Builder) build(p string) (*pkg, []*lex8.Error) {
	saved, found := b.find(p)
	if found && saved == nil {
		e := fmt.Errorf("package %s has circular dependency", p)
		return nil, lex8.SingleErr(e)
	} else if found {
		return saved, nil // already built
	}

	b.register(p)

	// get the language
	lang := b.getLang(p)

	// setup the package folder
	ret, e := newPkg(b.home, p, lang)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	imports, es := lang.ListImport(ret.srcFiles())
	if es != nil {
		return nil, es
	}

	// build import first
	for _, imp := range imports {
		_, es = b.build(imp)
		if es != nil {
			return nil, es
		}
	}

	// ready to build this one
	if b.Verbose {
		fmt.Println(p)
	}

	// compile now
	built, es := lang.Compile(ret.srcFiles(), b)
	if es != nil {
		return nil, es
	}
	ret.built = built

	lib := built.Lib() // the linkable lib
	// a package with main entrance, build the bin
	if lib.HasFunc("main") {
		fout := b.home.makeBin(p)
		e := link8.LinkMain(lib, fout)
		if e != nil {
			return nil, lex8.SingleErr(e)
		}
	}

	// built library, save it into archive
	ret.lib = lib
	b.save(p, ret)

	return ret, nil
}

// Import imports a built linkable lib
func (b *Builder) Import(p string) pkg8.Linkable {
	saved, _ := b.find(p)
	if saved == nil {
		return nil
	}
	return saved.built
}

// Build builds a package
func (b *Builder) Build(p string) []*lex8.Error {
	_, es := b.build(p)
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
