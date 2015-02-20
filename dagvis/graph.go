// Package dagvis visualizes a DAG graph into 
// a structured, layered planer map.
package dagvis

// Graph is a directed graph
type Graph struct {
	Nodes map[string][]string
}
