package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	verties := []*Vertex{
		{Name: "A", Value: 1},
		{Name: "B", Value: 2},
	}
	edges := []*Edge{
		{Cost: 1.0, From: "A", To: "B"},
	}

	graph, err := NewGraph(verties, edges)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(graph.verties))
	assert.Equal(t, 1, len(graph.edges))
}

func TestAddVertext(t *testing.T) {
	graph, _ := NewGraph([]*Vertex{}, []*Edge{})

	vertex := &Vertex{Name: "A", Value: 1}
	err := graph.AddVertext(vertex)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(graph.verties))

	// Test adding duplicate vertex
	err = graph.AddVertext(vertex)
	assert.Error(t, err)
}

func TestAddEdge(t *testing.T) {
	verties := []*Vertex{
		{Name: "A", Value: 1},
		{Name: "B", Value: 2},
	}
	graph, _ := NewGraph(verties, []*Edge{})

	edge := &Edge{Cost: 1.0, From: "A", To: "B"}
	err := graph.AddEdge(edge)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(graph.edges))

	// Test adding edge with invalid vertex
	invalidEdge := &Edge{Cost: 1.0, From: "A", To: "C"}
	err = graph.AddEdge(invalidEdge)
	assert.Error(t, err)
}
