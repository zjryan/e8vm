package build8

type build struct {
	buildPath string
}

func (b *build) Packages() []string {
	panic("todo")
}

func (b *build) Files(p string) ([]string, error) {
	panic("todo")
}

func (b *build) Info(p string) (*Info, error) {
	panic("todo")
}

func (b *build) File(p, f string) ([]byte, error) {
	panic("todo")
}
