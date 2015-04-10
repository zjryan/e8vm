package build8

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"lonnie.io/e8vm/lex8"
)

func listSrcFiles(dir string, lang Lang) ([]string, error) {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}

	var ret []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if lang.IsSrc(name) {
			ret = append(ret, name)
		}
	}

	return ret, nil
}

// DirHome is a file system basd building home.
type DirHome struct {
	path  string
	langs *langPicker

	fileList map[string][]string

	Quiet bool
}

var _ Home = new(DirHome)

// NewDirHome creates a file system home storage with
// a particualr default language for compiling.
func NewDirHome(path string, lang Lang) *DirHome {
	if lang == nil {
		panic("must specify a default language")
	}

	ret := new(DirHome)
	ret.path = path
	ret.fileList = make(map[string][]string)
	ret.langs = newLangPicker(lang)

	return ret
}

func (h *DirHome) sub(pre, p string) string {
	return filepath.Join(h.path, pre, p)
}

func (h *DirHome) subFile(pre, p, f string) string {
	return filepath.Join(h.path, pre, p, f)
}

// ClearCache clears the file list cache
func (h *DirHome) ClearCache() {
	h.fileList = make(map[string][]string)
}

// AddLang registers a language with a particular path prefix
func (h *DirHome) AddLang(prefix string, lang Lang) {
	h.langs.addLang(prefix, lang)
}

// Pkgs lists all the packages inside this home folder.
func (h *DirHome) Pkgs(prefix string) []string {
	root := filepath.Join(h.path, "src")
	start := filepath.Join(root, prefix)
	var pkgs []string

	walkFunc := func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if !info.IsDir() {
			return nil
		} else if !lex8.IsPkgName(info.Name()) {
			return filepath.SkipDir
		}

		if root == p {
			return nil
		}

		path, e := filepath.Rel(root, p)
		if e != nil {
			panic(e)
		} else if path == "." {
			return nil
		}

		lang := h.Lang(path)
		if lang == nil {
			panic(path)
		}

		files, e := listSrcFiles(p, lang)
		if e != nil {
			return e
		}

		if len(files) > 0 {
			h.fileList[path] = files // caching
			pkgs = append(pkgs, path)
		}

		return nil
	}

	e := filepath.Walk(start, walkFunc)
	if e != nil && !h.Quiet {
		log.Fatal("error", e)
	}

	sort.Strings(pkgs)
	return pkgs
}

// Src lists all the source files inside this package.
func (h *DirHome) Src(p string) map[string]*File {
	if !isPkgPath(p) {
		panic("not package path")
	}

	lang := h.Lang(p)
	if lang == nil {
		return nil
	}

	files, found := h.fileList[p]
	if !found {
		files, e := listSrcFiles(p, lang)
		if e != nil && !h.Quiet {
			log.Fatal("error", e)
		}

		h.fileList[p] = files
	}

	if len(files) == 0 {
		return nil
	}

	ret := make(map[string]*File)
	for _, name := range files {
		filePath := h.subFile("src", p, name)
		ret[name] = &File{
			Path:       filePath,
			Name:       name,
			ReadCloser: newDirFile(filePath),
		}
	}

	return ret
}

// CreateBin returns the writer to write the binary
func (h *DirHome) CreateBin(p string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newDirFile(h.sub("bin", p+".e8"))
}

// CreateLib returns the writer to write the linkable library
func (h *DirHome) CreateLib(p string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newDirFile(h.sub("pkg", p+".e8a"))
}

// CreateLog returns the log writer for the particular name
func (h *DirHome) CreateLog(p, name string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newDirFile(h.subFile("log", p, name))
}

// Lang returns the language for the particular path.
// It searches for the longest prefix match
func (h *DirHome) Lang(p string) Lang { return h.langs.lang(p) }
