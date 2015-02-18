package build8

import (
	"os"
	"path/filepath"
	"sort"

	"lonnie.io/e8vm/lex8"
)

func listPkgs(h *home) ([]string, error) {
	src := h.srcRoot()

	m := make(map[string]struct{})

	e := filepath.Walk(src, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		name := info.Name()
		if info.IsDir() {
			if !lex8.IsPkgName(name) {
				return filepath.SkipDir
			}
		} else {
			if isSrc(name) {
				path, e := filepath.Rel(src, filepath.Dir(p))
				if e != nil {
					panic(e)
				}
				m[path] = struct{}{}
			}
		}

		return nil
	})

	if e != nil {
		return nil, e
	}

	var ret []string
	for p := range m {
		ret = append(ret, p)
	}

	sort.Strings(ret)

	return ret, nil
}
