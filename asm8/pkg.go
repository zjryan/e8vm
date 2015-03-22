package asm8

import (
	"path"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/pkg8"
)

type pkg struct {
	path  string
	files []*file

	imports *importDecl
}

func resolvePkg(p string, src pkg8.Files) (*pkg, []*lex8.Error) {
	res := newResolver()
	ret := new(pkg)
	ret.path = p

	for f, rc := range src {
		// parse the file first
		astFile, es := parse.File(f, rc)
		if es != nil {
			return nil, es
		}

		// then resolve the file
		file := resolveFile(res, astFile)
		ret.files = append(ret.files, file)

		// import policy
		if len(src) == 1 || path.Base(f) == "import.s" {
			if ret.imports != nil {
				res.Errorf(file.imports.Kw.Pos,
					"double valid import stmt; two import.s?",
				)
			} else {
				ret.imports = file.imports
			}
		} else if file.imports != nil {
			res.Errorf(file.imports.Kw.Pos,
				"invalid import outside import.s in a multi-file package",
			)
		}
	}

	if es := res.Errs(); es != nil {
		return nil, es
	}
	return ret, nil
}
