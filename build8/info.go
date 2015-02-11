package build8

import (
	"time"
)

// Info contains meta information of a package.
type Info struct {
	Language   string
	LastUpdate *time.Time
	LastBuild  *time.Time
}
