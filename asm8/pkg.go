package asm8

import (
	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
)

type pkg struct {
	path  string
	files []*file

	imports *importDecl
}

func resolvePkg(p string, src map[string]*build8.File) (*pkg, []*lex8.Error) {
	log := lex8.NewErrorList()
	ret := new(pkg)
	ret.path = p

	for name, f := range src {
		// parse the file first
		astFile, es := parse.File(f.Path, f)
		if es != nil {
			return nil, es
		}

		// then resolve the file
		file := resolveFile(log, astFile)
		ret.files = append(ret.files, file)

		// enforce import policy
		if len(src) == 1 || name == "import.s" {
			if ret.imports != nil {
				log.Errorf(file.imports.Kw.Pos,
					"double valid import stmt; two import.s?",
				)
			} else {
				ret.imports = file.imports
			}
		} else if file.imports != nil {
			log.Errorf(file.imports.Kw.Pos,
				"invalid import outside import.s in a multi-file package",
			)
		}
	}

	if es := log.Errs(); es != nil {
		return nil, es
	}
	return ret, nil
}
