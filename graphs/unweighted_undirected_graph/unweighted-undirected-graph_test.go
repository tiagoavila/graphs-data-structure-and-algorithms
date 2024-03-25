package unweightedundirectedgraph

import (
	"reflect"
	"testing"
)

// TestNewGraph tests the creation of a new graph
func TestNewGraph(t *testing.T) {
	vertices := 4
	graph := NewGraph(vertices)

	if graph.Vertices != vertices {
		t.Errorf("NewGraph(%d): expected %d vertices, got %d", vertices, vertices, graph.Vertices)
	}

	if graph.Adjacency == nil {
		t.Errorf("NewGraph(%d): adjacency list is nil", vertices)
	}

	if len(graph.Adjacency) != 0 {
		t.Errorf("NewGraph(%d): expected empty adjacency list, got %v", vertices, graph.Adjacency)
	}
}

// TestAddEdge tests adding edges to the graph
func TestAddEdge(t *testing.T) {
	graph := NewGraph(5)
	edges := []Edge{
		{Src: 1, Dest: 2},
		{Src: 1, Dest: 3},
		{Src: 2, Dest: 4},
		{Src: 3, Dest: 4},
	}

	for _, edge := range edges {
		graph.AddEdge(edge.Src, edge.Dest)
	}

	// Check if the edges are added correctly
	expectedAdjacency := map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1, 4},
		4: {2, 3},
	}

	if !reflect.DeepEqual(graph.Adjacency, expectedAdjacency) {
		t.Errorf("AddEdge: expected adjacency list %v, got %v", expectedAdjacency, graph.Adjacency)
	}

	// Check for the undirected property
	for src, destinations := range graph.Adjacency {
		for _, dest := range destinations {
			found := false
			for _, srcCheck := range graph.Adjacency[dest] {
				if srcCheck == src {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("AddEdge: edge from %d to %d is not undirected", src, dest)
			}
		}
	}
}
