package asm8

func traceUsed(lnk *linker, p *pkgObj, index uint32) []pkgSym {
	// create a hit map for all symbols
	used := make([][]bool, lnk.Npkg())
	for i, p := range lnk.pkgs {
		used[i] = make([]bool, len(p.symbols))
	}

	cur := []pkgSym{{p, index}}
	var next []pkgSym
	var ret []pkgSym

	// BFS traverse all the symbols used by the symbol
	for len(cur) > 0 {
		for _, u := range cur {
			ret = append(ret, u)

			s := u.p.symbols[u.sym]
			if s.Type == SymFunc {
				f := u.p.FuncObj(u.sym)
				for _, link := range f.links {
					pkg := u.p.requires[link.pkg] // the pacakge
					i := lnk.PkgIndex(pkg.path)   // fetch the index
					pt := &used[i][link.sym]      // locate the bool ptr
					if *pt {
						continue
					}

					*pt = true
					item := pkgSym{pkg, link.sym}
					next = append(next, item)
				}
			}
		}

		cur = next
		next = nil
	}

	return ret
}
