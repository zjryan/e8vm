package build8

import (
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/lex8"
)

// BuildAll packages under gpath
func BuildAll(gpath string) []*lex8.Error {
	b := NewBuild(gpath)
	src := b.src("")

	pkgs := make(map[string]struct{})

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
			if strings.HasSuffix(name, ".g") || strings.HasSuffix(name, ".s") {
				dir := filepath.Dir(p)
				pkg, e := filepath.Rel(src, dir)
				if e != nil {
					panic(e)
				}
				pkgs[pkg] = struct{}{}
			}
		}

		return nil
	})

	if e != nil {
		return lex8.SingleErr(e)
	}

	// TODO: build with dependency
	for pkg := range pkgs {
		_, es := b.BuildAsm(pkg)
		if es != nil {
			return es
		}
	}

	return nil
}
