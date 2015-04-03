package build8

import (
	"strings"
)

type langPicker struct {
	defaultLang Lang
	langs       map[string]Lang
}

func newLangPicker(def Lang) *langPicker {
	if def == nil {
		panic("default language must not be nil")
	}

	ret := new(langPicker)
	ret.defaultLang = def
	ret.langs = make(map[string]Lang)
	return ret
}

func (pick *langPicker) addLang(prefix string, lang Lang) {
	if lang == nil {
		panic("language must not be nil")
	}

	if prefix == "" {
		pick.defaultLang = lang
	}
	pick.langs[prefix] = lang
}

func (pick *langPicker) lang(path string) Lang {
	if !isPkgPath(path) {
		panic("not package path")
	}

	nmax := -1
	var ret Lang
	for prefix, lang := range pick.langs {
		n := len(prefix)
		if n < nmax || !strings.HasPrefix(path, prefix) {
			continue
		}

		nmax = n
		ret = lang
	}

	if ret == nil {
		ret = pick.defaultLang
	}
	return ret
}
