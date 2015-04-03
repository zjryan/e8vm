package build8

// MemPkg is a package in memory
type MemPkg struct {
	path  string
	files map[string]*memFile
	logs  map[string]*memFile
	bin   *memFile
	lib   *memFile
}

func newMemPkg(path string) *MemPkg {
	ret := new(MemPkg)
	ret.path = path
	ret.logs = make(map[string]*memFile)
	ret.files = make(map[string]*memFile)
	return ret
}

// Path returns the path of the package
func (p *MemPkg) Path() string { return p.path }

// AddFile adds (or replaces) a source file in the package
func (p *MemPkg) AddFile(path, name, content string) {
	f := newMemFile()
	f.path = path
	f.WriteString(content)
	p.files[name] = f
}
