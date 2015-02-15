package fs8

import (
	"os"
	"path/filepath"
)

// IsPackageName checks if the string is a valid package name.
// A package name is valid when it only contains 'a'-'z' and '0'-'9',
// where it cannot start with a digit.
func IsPackageName(name string) bool {
	if name == "" {
		return false
	}

	for i, r := range name {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		return false
	}
	return true
}

// AllPackages return all the packages under a certain root.
func AllPackages(root string) ([]string, error) {
	var ret []string
	walk := func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if !info.IsDir() {
			return nil // skip all non-dirs
		}

		p, e = filepath.Rel(root, p)
		if e != nil {
			panic(e)
		}

		base := filepath.Base(p)

		if base == "." {
			return nil // this is the root
		}

		if !IsPackageName(base) {
			return filepath.SkipDir // skip the entire dir
		}

		ret = append(ret, p)
		return nil
	}

	err := filepath.Walk(root, walk)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
