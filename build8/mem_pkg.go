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

func (p *MemPkg) Path() string { return p.path }
func (p *MemPkg) AddFile(name, content string) {
	f := newMemFile()
	f.WriteString(content)
	p.files[name] = f
}
