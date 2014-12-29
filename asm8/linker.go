package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type linker struct {
	pkgs     []*pkgObj
	pkgIndex map[string]int

	used [][]bool
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

func (lnk *linker) trackUsed(p *pkgObj, f *funcObj) {
	lnk.used = make([][]bool, len(lnk.pkgs))
	for i, p := range lnk.pkgs {
		lnk.used[i] = make([]bool, len(p.symbols))
	}

	type ref struct {
		p *pkgObj
		f *funcObj
	}

	cur := []ref{{p, f}}
	var next []ref

	// BFS traverse all the symbols used by f
	for len(cur) > 0 {
		for _, r := range cur {
			for _, link := range r.f.links {
				pkg := r.p.requires[link.pkg]
				pindex := lnk.pkgIndex[pkg.path]
				sindex := link.sym
				pt := &lnk.used[pindex][sindex]
				if !*pt {
					*pt = true
					next = append(next, ref{pkg, pkg.funcs[sindex]})
				}
			}
		}

		cur = next
		next = nil
	}
}

func linkPkg(main *pkgObj) ([]byte, *lex8.Error) {
	lnk := newLinker()
	lnk.addPkgs(main)

	funcMain, index := main.Query("main")
	if funcMain.Type != SymFunc {
		return nil, lex8.Errorf("main functiongg missing in package %s", main.path)
	}
	f, found := main.funcs[index]
	if !found {
		panic("main function lost")
	}
	lnk.trackUsed(main, f)

	return nil, nil
}
