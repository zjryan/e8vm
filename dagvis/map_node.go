package dagvis

// MapNode is a node in the DAG graph
type MapNode struct {
	Name string
	X, Y int // layout position

	Ins  map[string]*MapNode // direct input nodes
	Outs map[string]*MapNode // direct output nodes

	AllIns  map[string]*MapNode // direct and indirect input nodes
	AllOuts map[string]*MapNode // direct and indirect output nodes

	// critical nodes is the minimum set of nodes that keeps
	// the same partial order of the nodes in the DAG graph
	CritIns  map[string]*MapNode // critical input nodes
	CritOuts map[string]*MapNode // critical output nodes

	layer    int // min layer
	newLayer int // new layer after pushing

	nhit int // for counting on layers
}

func newMapNode(name string) *MapNode {
	ret := new(MapNode)
	ret.Name = name

	ret.AllIns = make(map[string]*MapNode)
	ret.AllOuts = make(map[string]*MapNode)

	ret.CritIns = make(map[string]*MapNode)
	ret.CritOuts = make(map[string]*MapNode)

	ret.Ins = make(map[string]*MapNode)
	ret.Outs = make(map[string]*MapNode)

	return ret
}
