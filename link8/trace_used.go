package link8

type tracer struct {
	lnk  *Linker
	hits [][]bool
}

func newTracer(lnk *Linker) *tracer {
	npkg := lnk.Npkg()
	ret := new(tracer)
	ret.lnk = lnk
	ret.hits = make([][]bool, npkg)
	for i := 0; i < npkg; i++ {
		p := lnk.Pkg(i)
		ret.hits[i] = make([]bool, len(p.symbols))
	}

	return ret
}

func (t *tracer) hit(pkg *Package, sym uint32) bool {
	i := t.lnk.PkgIndex(pkg.path)
	pt := &t.hits[i][sym]
	ret := *pt
	*pt = true
	return ret
}

func traceUsed(lnk *Linker, p *Package, index uint32) []pkgSym {
	t := newTracer(lnk)

	cur := []pkgSym{{p, index}}
	var next []pkgSym
	var ret []pkgSym

	// BFS traverse all the symbols used by the symbol
	for len(cur) > 0 {
		for _, ps := range cur {
			ret = append(ret, ps)

			typ := ps.Type()
			if typ == SymFunc {
				f := ps.Func()

				for _, link := range f.links {
					pkg := ps.Import(link.pkg)
					if !t.hit(pkg, link.sym) {
						continue
					}

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
