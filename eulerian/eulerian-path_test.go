package eulerian

import (
	unweightedundirectedgraph "graphs_data_structure_and_algorithms/graphs/unweighted_undirected_graph"
	"testing"
)

// helper function to create a graph
func createGraph(vertices int, edges [][]int) *unweightedundirectedgraph.UnWeightedUndirectedGraph {
	g := unweightedundirectedgraph.NewGraph(vertices)
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}
	return g
}

func TestIsConnected(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][]int
		want     bool
	}{
		{"empty graph", 0, nil, true},
		{"single vertex", 1, nil, true},
		{"connected graph", 4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true},
		{"disconnected graph", 4, [][]int{{0, 1}, {2, 3}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := createGraph(tt.vertices, tt.edges)
			if got := IsConnected(g); got != tt.want {
				t.Errorf("IsConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEulerian(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		edges    []unweightedundirectedgraph.Edge
		vertices int
		expected int
	}{
		{
			name:     "Eulerian graph",
			vertices: 4,
			edges: []unweightedundirectedgraph.Edge{
				{Src: 0, Dest: 1},
				{Src: 1, Dest: 2},
				{Src: 2, Dest: 3},
				{Src: 3, Dest: 0},
			},
			expected: 2,
		},
		{
			name:     "Semi-Eulerian graph",
			vertices: 3,
			edges: []unweightedundirectedgraph.Edge{
				{Src: 0, Dest: 1},
				{Src: 1, Dest: 2},
			},
			expected: 1,
		},
		{
			name:     "Non-Eulerian graph",
			vertices: 5,
			edges: []unweightedundirectedgraph.Edge{
				{Src: 0, Dest: 1},
				{Src: 1, Dest: 2},
				{Src: 1, Dest: 3},
				{Src: 0, Dest: 2},
				{Src: 0, Dest: 3},
				{Src: 3, Dest: 4},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := unweightedundirectedgraph.NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge.Src, edge.Dest)
			}

			result := IsEulerian(g)
			if result != tt.expected {
				t.Errorf("For %s, expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}
