package build8

// Build presents interface for querying packages under a build repo.
// It serves like an abstract file system for building.
type Build interface {
	// Packages list all packages
	Packages() []string

	// Files list all the files in a package
	Files(p string) ([]string, error)

	// Info returns the package meta info structure
	Info(p string) (*Info, error)

	// Imports lists the package dependencies
	Imports(p string) ([]string, error)

	// File fetches the content of a file
	File(p, f string) ([]byte, error)
}
