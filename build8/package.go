package build8

type Package struct {
	lang string // language
	name string
	path string

	srcPath string
	libPath string
	files   []string
}

func (p *Package) listFiles() {

}
