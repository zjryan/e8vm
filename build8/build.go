package build8

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/asm8"
)

// Build is a build folder for our language.
type Build struct {
	path string
}

// AsmPkg creates an asm pkg build for our asm package.
func (b *Build) AsmPkg(path string) (*asm8.PkgBuild, error) {
	folder := filepath.Join(b.path, "src", path)

	f, e := os.Open(folder)
	if e != nil {
		return nil, e
	}

	files, e := f.Readdir(0)
	if e != nil {
		return nil, e
	}

	var srcFiles map[string]io.ReadCloser

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !strings.HasSuffix(name, ".s") {
			continue
		}

		filename := filepath.Join(folder, name)
		srcFiles[filename] = newFile(filename)
	}

	ret := &asm8.PkgBuild{
		Path:   path,
		Import: nil, // TODO
		Files:  srcFiles,
	}
	return ret, nil
}
