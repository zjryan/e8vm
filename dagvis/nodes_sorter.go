package dagvis

type mapNodes []*MapNode

func (s mapNodes) Len() int      { return len(s) }
func (s mapNodes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type byLayer struct{ mapNodes }

func (s byLayer) Less(i, j int) bool {
	a := s.mapNodes[i]
	b := s.mapNodes[j]

	if a.layer < b.layer {
		return true
	}
	if a.layer > b.layer {
		return false
	}
	return a.Name < b.Name
}

type byNcritOuts struct{ mapNodes }

func (s byNcritOuts) Less(i, j int) bool {
	n1 := s.mapNodes[i]
	n2 := s.mapNodes[j]

	nout1 := len(n1.CritOuts)
	nout2 := len(n2.CritOuts)

	if nout1 < nout2 {
		return true
	}
	if nout1 < nout2 {
		return false
	}

	return n1.Name < n2.Name
}
