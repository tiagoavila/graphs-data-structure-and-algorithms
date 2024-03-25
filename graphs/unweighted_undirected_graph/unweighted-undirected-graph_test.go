package unweightedundirectedgraph

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph(5)

	if graph.Vertices != 5 {
		t.Errorf("Expected 5 vertices, got %d", graph.Vertices)
	}

	if len(graph.Adjacency) != 0 {
		t.Errorf("Expected empty adjacency list, got %v", graph.Adjacency)
	}
}

func TestAddEdge(t *testing.T) {
	graph := NewGraph(2)
	graph.AddEdge(0, 1)

	if len(graph.Adjacency) != 2 {
		t.Errorf("Expected adjacency list to have 2 entries, got %d", len(graph.Adjacency))
	}

	if !reflect.DeepEqual(graph.Adjacency[0], []Edge{{Src: 0, Dest: 1}}) {
		t.Errorf("Expected adjacency list of vertex 0 to be [{0 1}], got %v", graph.Adjacency[0])
	}

	if !reflect.DeepEqual(graph.Adjacency[1], []Edge{{Src: 1, Dest: 0}}) {
		t.Errorf("Expected adjacency list of vertex 1 to be [{1 0}], got %v", graph.Adjacency[1])
	}
}

// Additional tests can be written to further validate the functionality
// such as adding multiple edges, edges with non-existent vertices, etc.
