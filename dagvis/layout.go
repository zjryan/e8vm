package dagvis

// Layout layouts a DAG into a map.
func Layout(g *Graph) (*Map, error) {
	m, e := newMap(g)
	if e != nil {
		return nil, e
	}

	pushTight(m)

	return m, nil
}
