package build8

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Builder builds a bunch of packages.
type Builder struct {
	home *home
	pkgs map[string]*pkg

	// TODO: built should be something like an LRU cache
	// the libraries should be load back in only when linking

	langs       map[string]Lang
	defaultLang Lang

	Verbose bool
}

// NewBuilder creates a new builder
func NewBuilder(homePath string) *Builder {
	ret := new(Builder)
	ret.home = &home{path: homePath}
	ret.pkgs = make(map[string]*pkg)
	ret.langs = make(map[string]Lang)

	return ret
}

// AddLang registers a langauge into the building system
func (b *Builder) AddLang(prefix string, lang Lang) error {
	// TODO: use a prefix tree
	// for now, we just force no two prefixes conflict.

	if prefix == "" {
		b.defaultLang = lang
		return nil
	}

	for other := range b.langs {
		if strings.HasPrefix(other, prefix) ||
			strings.HasPrefix(prefix, other) {
			return fmt.Errorf("prefix %q conflict with %q", prefix, other)
		}
	}

	b.langs[prefix] = lang
	return nil
}

func (b *Builder) getLang(p string) Lang {
	for prefix, lang := range b.langs {
		if strings.HasPrefix(p, prefix) {
			return lang
		}
	}

	return b.defaultLang
}

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

	es := lang.Prepare(pkg.srcMap(), pkg)
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
		imp.Compiled = built.compiled
	}

	// ready to build this one
	if b.Verbose {
		fmt.Println(p)
	}

	// compile now
	compiled, es := lang.Compile(p, ret.srcMap(), ret.imports)
	if es != nil {
		return nil, es
	}
	ret.compiled = compiled

	lib := ret.compiled.Lib() // the linkable lib
	// a package with main entrance, build the bin

	main := ret.compiled.Main()
	if main != "" && lib.HasFunc(main) {
		log := lex8.NewErrorList()

		fout := b.home.makeBin(p)
		lex8.LogError(log, link8.LinkMain(lib, fout, main))
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
