package build8

import (
	"fmt"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Builder builds a bunch of packages.
type Builder struct {
	home  *home
	built map[string]*pkg
	// TODO: built should be something like an LRU cache
	// the libraries should be load back in only when linking

	Verbose bool
}

// NewBuilder creates a new builder
func NewBuilder(homePath string) *Builder {
	ret := new(Builder)
	ret.home = &home{path: homePath}
	ret.built = make(map[string]*pkg)

	return ret
}

func (b *Builder) buildImports(p *pkg) (*imports, []*lex8.Error) {
	imports, es := p.loadImports()
	if es != nil {
		return nil, es
	}

	if imports != nil {
		for _, imp := range imports.m {
			imported, es := b.build(imp.path)
			if es != nil {
				return nil, es
			}

			imp.lib = imported.lib
		}
	}

	return imports, nil
}

func (b *Builder) build(p string) (*pkg, []*lex8.Error) {
	ret, found := b.built[p]
	if found {
		if ret == nil {
			// registered but not built, must be circular
			// dependency
			e := fmt.Errorf("package %q has circular dependency", p)
			return nil, lex8.SingleErr(e)
		}
		return ret, nil
	}

	// register the package slot
	b.built[p] = nil

	// setup the package folder
	ret, e := newPkg(b.home, p)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	// load imports
	// will recursively call b.build if the imported lib has
	// not been built.
	imports, es := b.buildImports(ret)
	if es != nil {
		return nil, es
	}

	// ready
	if b.Verbose {
		fmt.Println(p)
	}

	// go build
	lib, es := ret.build(imports)
	if es != nil {
		return nil, es
	}

	// a package with main entrance, build the bin
	if lib.HasFunc("main") {
		fout := b.home.makeBin(p)
		e := link8.LinkMain(lib, fout)
		if e != nil {
			return nil, lex8.SingleErr(e)
		}
	}

	// built library, save it into archive
	b.built[p] = ret

	return ret, nil
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
