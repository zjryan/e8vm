package build8

import (
	"io"
	"sort"
)

// MultiHome is a stack of homes. When called for a package, it will search
// for each home respectively
type MultiHome struct {
	homes []Home
}

// NewMultiHome creates a new stack of homes
func NewMultiHome(homes ...Home) *MultiHome {
	if len(homes) == 0 {
		panic("must have at least one home")
	}

	ret := new(MultiHome)
	ret.homes = homes
	return ret
}

// Pkgs lists the pacakges
func (h *MultiHome) Pkgs(prefix string) []string {
	ret := make(map[string]struct{})
	for _, home := range h.homes {
		pkgs := home.Pkgs(prefix)
		for _, pkg := range pkgs {
			ret[pkg] = struct{}{}
		}
	}

	var lst []string
	for p := range ret {
		lst = append(lst, p)
	}
	sort.Strings(lst)
	return lst
}

// HomeFor returns the home for a particular path.
func (h *MultiHome) HomeFor(path string) Home {
	for _, home := range h.homes {
		src := home.Src(path)
		if src != nil {
			return home
		}
	}
	return nil
}

// Src lists the source files in a package. It returns nil
// when the package does not exist.
func (h *MultiHome) Src(path string) map[string]*File {
	for _, home := range h.homes {
		src := home.Src(path)
		if len(src) > 0 {
			return src
		}
	}
	return nil
}

// CreateLib creates the writer for writing the linkable package library
func (h *MultiHome) CreateLib(path string) io.WriteCloser {
	return h.HomeFor(path).CreateLib(path)
}

// CreateLog creates the logger
func (h *MultiHome) CreateLog(path, name string) io.WriteCloser {
	return h.HomeFor(path).CreateLog(path, name)
}

// CreateBin creates the writer for writing the E8 binary
func (h *MultiHome) CreateBin(path, name string) io.WriteCloser {
	return h.HomeFor(path).CreateBin(path)
}

// Lang returns the language of a path. If the package exists in a home
// it will return the language in the package. If the package does not exist
// if any of the homes, it will return the language from the first home.
func (h *MultiHome) Lang(path string) Lang {
	home := h.HomeFor(path)
	if home == nil {
		home = h.homes[0]
	}
	return home.Lang(path)
}
