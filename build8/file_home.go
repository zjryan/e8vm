package build8

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

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

// FileHome is a file system basd building home.
type FileHome struct {
	path string

	defaultLang Lang
	langs       map[string]Lang

	fileList map[string][]string

	Quiet bool
}

var _ Home = new(FileHome)

// NewFileHome creates a file system home storage with
// a particualr default language for compiling.
func NewFileHome(path string, lang Lang) *FileHome {
	if lang == nil {
		panic("must specify a default language")
	}

	ret := new(FileHome)
	ret.path = path
	ret.defaultLang = lang
	ret.fileList = make(map[string][]string)
	ret.langs = make(map[string]Lang)

	return ret
}

func (h *FileHome) sub(pre, p string) string {
	return filepath.Join(h.path, pre, p)
}

func (h *FileHome) subFile(pre, p, f string) string {
	return filepath.Join(h.path, pre, p, f)
}

// ClearCache clears the file list cache
func (h *FileHome) ClearCache() {
	h.fileList = make(map[string][]string)
}

// AddLang registers a language with a particular path prefix
func (h *FileHome) AddLang(prefix string, lang Lang) {
	if lang == nil {
		panic("language must not be nil")
	}

	if prefix == "" {
		h.defaultLang = lang
	}
	h.langs[prefix] = lang
}

// Pkgs lists all the packages inside this home folder.
func (h *FileHome) Pkgs(prefix string) []string {
	root := filepath.Join(h.path, "src")
	start := filepath.Join(root, prefix)
	var pkgs []string

	walkFunc := func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		name := info.Name()
		if !info.IsDir() {
			return nil
		} else if !lex8.IsPkgName(name) {
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
			h.fileList[path] = files
			pkgs = append(pkgs, path)
		}

		return nil
	}

	e := filepath.Walk(start, walkFunc)
	if e != nil && !h.Quiet {
		log.Fatal(e)
	}

	sort.Strings(pkgs)
	return pkgs
}

// Src lists all the source files inside this package.
func (h *FileHome) Src(p string) map[string]*File {
	if !isPkgPath(p) {
		panic("not package path")
	}

	lang := h.Lang(p)
	if lang == nil {
		return nil
	}

	files := h.fileList[p]
	if files == nil {
		files, e := listSrcFiles(p, lang)
		if e != nil && !h.Quiet {
			log.Fatal(e)
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
			ReadCloser: newFile(filePath),
		}
	}

	return ret
}

// Bin returns the writer to write the binary
func (h *FileHome) Bin(p string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newFile(h.sub("bin", p+".e8"))
}

// Lib returns the writer to write the linkable library
func (h *FileHome) Lib(p string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newFile(h.sub("pkg", p+".e8a"))
}

// Log returns the log writer for the particular name
func (h *FileHome) Log(p, name string) io.WriteCloser {
	if !isPkgPath(p) {
		panic("not package path")
	}
	return newFile(h.subFile("src", p, name))
}

// Lang returns the language for the particular path.
// It searches for the longest prefix match
func (h *FileHome) Lang(p string) Lang {
	if !isPkgPath(p) {
		panic("not package path")
	}

	nmax := -1
	var ret Lang
	for prefix, lang := range h.langs {
		n := len(prefix)
		if n < nmax || !strings.HasPrefix(p, prefix) {
			continue
		}

		nmax = n
		ret = lang
	}

	if ret == nil {
		ret = h.defaultLang
	}
	return ret
}
