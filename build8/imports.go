package build8

type imports struct {
	m map[string]*pkgImport
}

func newImports() *imports {
	ret := new(imports)
	ret.m = make(map[string]*pkgImport)
	return ret
}
