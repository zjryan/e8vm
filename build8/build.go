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
}

// NewBuild returns a build based on a build path
func NewBuild(path string) *Build {
	ret := new(Build)
	ret.path = path
	return ret
}

func (b *Build) join(pre, p string) string {
	return filepath.Join(b.path, pre, p)
}

func (b *Build) src(p string) string { return b.join("src", p) }
func (b *Build) bin(p string) string { return b.join("bin", p+".e8") }
func (b *Build) pkg(p string) string { return b.join("pkg", p+".e8a") }

// AsmPkg creates an asm pkg build for our asm package.
func (b *Build) newAsmPkg(path string) (*asmPkg, error) {
	folder := b.src(path)

	f, e := os.Open(folder)
	if e != nil {
		return nil, e
	}

	files, e := f.Readdir(0)
	if e != nil {
		return nil, e
	}

	srcFiles := make(map[string]io.ReadCloser)

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

	ret := &asmPkg{
		path:       path,
		importFile: nil,
		files:      srcFiles,
	}
	return ret, nil
}

// needRebuild checks if the package requires a rebuild
func (b *Build) needRebuild(pb *asm8.PkgBuild) (bool, error) {
	// TODO:
	return true, nil
}

// BuildAsm builds an assembly package into a binary
func (b *Build) BuildAsm(path string) (int, []*lex8.Error) {
	asm, e := b.newAsmPkg(path)
	if e != nil {
		return 0, lex8.SingleErr(e)
	}

	pb := asm.build()

	nbuilt := 0
	if pb.Import != nil {
		var es []*lex8.Error
		for importPkg := range pb.Import {
			nbuilt, es = b.BuildAsm(importPkg)
			if es != nil {
				return nbuilt, es
			}
			// TODO: bind the lib imported
		}
	}

	rebuild, e := b.needRebuild(pb)
	if e != nil {
		return nbuilt, lex8.SingleErr(e)
	}
	if !rebuild {
		return nbuilt, nil
	}

	p, es := pb.Build()
	if es != nil {
		return nbuilt, es
	}

	// TODO: save the lib

	if p.HasFunc("main") {
		fout := newFile(b.bin(path))
		e := link8.LinkMain(p, fout)
		if e != nil {
			return nbuilt, lex8.SingleErr(e)
		}
	}

	return nbuilt, nil
}
