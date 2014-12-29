package asm8

type linker struct {
	pkgs     []*pkgObj
	pkgIndex map[string]int
}

func newLinker() *linker {
	ret := new(linker)
	ret.pkgIndex = make(map[string]int)
	return ret
}

func (lnk *linker) addPkg(p *pkgObj) (int, bool) {
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

func (lnk *linker) addPkgs(p *pkgObj) int {
	index, isNew := lnk.addPkg(p)
	if isNew {
		for _, req := range p.requires {
			lnk.addPkg(req)
		}
	}

	return index
}

// PkgIndex returns the index of the particular package
func (lnk *linker) PkgIndex(path string) int {
	ret, found := lnk.pkgIndex[path]
	if !found {
		panic("not found")
	}
	return ret
}

// Npkg retunrs the total number of packages
func (lnk *linker) Npkg() int {
	return len(lnk.pkgs)
}
