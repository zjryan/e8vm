package build8

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"lonnie.io/e8vm/lex8"
)

// Builder builds a bunch of packages.
type Builder struct {
	home  *home
	built map[string]*pkg

	Verbose bool
}

// NewBuilder creates a new builder
func NewBuilder(homePath string) *Builder {
	ret := new(Builder)
	ret.home = &home{path: homePath}
	ret.built = make(map[string]*pkg)

	return ret
}

func (b *Builder) build(p string) (*pkg, []*lex8.Error) {
	ret, found := b.built[p]
	if found {
		return ret, nil
	}

	ret, e := newPkg(b.home, p)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	es := ret.build(b)
	if es != nil {
		return nil, es
	}

	return ret, nil
}

func (b *Builder) prebuild(p string) {
	if b.Verbose {
		fmt.Println(p)
	}
}

// Build builds a package
func (b *Builder) Build(p string) []*lex8.Error {
	_, es := b.build(p)
	return es
}

func (b *Builder) ListPkgs() ([]string, error) {
	src := b.home.srcRoot()

	m := make(map[string]struct{})

	e := filepath.Walk(src, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		name := info.Name()
		if info.IsDir() {
			if !isPkgName(name) {
				return filepath.SkipDir
			}
		} else {
			if isSrc(name) {
				path, e := filepath.Rel(src, filepath.Dir(p))
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
