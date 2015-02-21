package link8

// Linker provides the framework to link packages together.
type Linker struct {
	pkgs     []*Pkg
	pkgIndex map[string]int
}

// NewLinker creates a new linker with no packages.
func NewLinker() *Linker {
	ret := new(Linker)
	ret.pkgIndex = make(map[string]int)
	return ret
}

// AddPkg adds a package into the linker.
// It returns the index of the package and if the package
// is new.
func (lnk *Linker) AddPkg(p *Pkg) (index int, isNew bool) {
	path := p.path
	index, found := lnk.pkgIndex[path]
	if found {
		return index, false
	}

	index = len(lnk.pkgs)
	lnk.pkgs = append(lnk.pkgs, p)
	lnk.pkgIndex[path] = index
	return index, true
}

// AddPkgs adds the package and recursively adds the packages
// it requires. It returns the package index.
func (lnk *Linker) AddPkgs(p *Pkg) int {
	index, isNew := lnk.AddPkg(p)
	if isNew {
		for _, req := range p.requires {
			lnk.AddPkgs(req)
		}
	}

	return index
}

// PkgIndex returns the index of the particular package.
func (lnk *Linker) PkgIndex(path string) int {
	ret, found := lnk.pkgIndex[path]
	if !found {
		panic("not found")
	}
	return ret
}

// Pkg returns the package of index i.
func (lnk *Linker) Pkg(i int) *Pkg {
	return lnk.pkgs[i]
}

// Npkg retunrs the total number of packages.
func (lnk *Linker) Npkg() int {
	return len(lnk.pkgs)
}
