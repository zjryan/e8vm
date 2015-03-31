package g8

import (
	"lonnie.io/e8vm/link8"
)

type pkg struct {
	lib *link8.Pkg
}

func (p *pkg) Lib() *link8.Pkg { return p.lib }
