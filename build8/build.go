package build8

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Build is a build folder for our language.
type Build struct {
	path string

	built map[string]bool
}

// NewBuild returns a build based on a build path
func NewBuild(path string) *Build {
	ret := new(Build)
	ret.path = path
	ret.built = make(map[string]bool)
	return ret
}

func (b *Build) join(pre, p string) string {
	return filepath.Join(b.path, pre, p)
}

func (b *Build) src(p string) string { return b.join("src", p) }
func (b *Build) bin(p string) string { return b.join("bin", p+".e8") }
func (b *Build) pkg(p string) string { return b.join("pkg", p+".e8a") }

// AsmPkg creates an asm pkg build for our asm package.
func (b *Build) prepareAsm(path string) (*asm8.PkgBuild, error) {
	folder := b.src(path)

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

// BuildAsm builds an assembly package into a binary
func (b *Build) BuildAsm(path string) []*lex8.Error {
	pb, e := b.prepareAsm(path)
	if e != nil {
		return lex8.SingleErr(e)
	}

	p, es := pb.Build()
	if es != nil {
		return es
	}

	// TODO: save the lib

	if p.HasFunc("main") {
		fout := newFile(b.bin(path))
		e := link8.LinkMain(p, fout)
		if e != nil {
			return lex8.SingleErr(e)
		}
	}

	return nil
}
