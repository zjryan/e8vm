package build8

import (
	"fmt"

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

// ListPkgs list all the packages under a home folder
func (b *Builder) ListPkgs() ([]string, error) {
	return listPkgs(b.home)
}
