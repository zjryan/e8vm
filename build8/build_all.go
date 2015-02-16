package build8

import (
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/lex8"
)

func IsPkgName(s string) bool {
	for i, r := range s {
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		return false
	}

	return true
}

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
			if !IsPkgName(name) {
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
		es := b.BuildAsm(pkg)
		if es != nil {
			return es
		}
	}

	return nil
}
