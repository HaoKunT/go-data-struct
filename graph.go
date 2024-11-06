package datastruct

import (
	"fmt"
)

// Vertex is a node in graph
type Vertex struct {
	Name  string
	Value interface{}
}

// Edge is a connection between two verties
type Edge struct {
	Cost float64
	From string // the name of start vertex
	To   string // the name of end vertex
}

// Graph is a collection of verties and edges
type Graph struct {
	//we need use map instead of list, because we need make sure the node name/id will not be duplicated
	verties map[string]*Vertex
	edges   []*Edge
}

// NewGraph creates a new graph
func NewGraph(verties []*Vertex, edges []*Edge) (*Graph, error) {
	g := new(Graph)
	vx := make(map[string]*Vertex)
	for _, v := range verties {
		vx[v.Name] = v
	}
	for _, e := range edges {
		_, okf := vx[e.From]
		_, okt := vx[e.To]
		if !okf || !okt {
			return nil, fmt.Errorf("invalid edge: %v", e)
		}
	}
	g.verties = vx
	g.edges = edges
	return g, nil
}

// AddVertext adds a vertex to the graph
func (g *Graph) AddVertext(vertex *Vertex) error {
	if _, ok := g.verties[vertex.Name]; ok {
		return fmt.Errorf("duplicated vertex name: %s", vertex.Name)
	}
	g.verties[vertex.Name] = vertex
	return nil
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(edge *Edge) error {
	if _, ok := g.verties[edge.From]; !ok {
		return fmt.Errorf("invalid edge: %v", edge)
	}
	if _, ok := g.verties[edge.To]; !ok {
		return fmt.Errorf("invalid edge: %v", edge)
	}
	g.edges = append(g.edges, edge)
	return nil
}
