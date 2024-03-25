package articulationpoints

import (
	unweightedundirectedgraph "graphs_data_structure_and_algorithms/graphs/unweighted_undirected_graph"
)

// ArticulationPointFinder encapsulates the necessary parameters for Tarjan's algorithm
type ArticulationPointFinder struct {
	graph  *unweightedundirectedgraph.UnWeightedUndirectedGraph
	time   int
	disc   []int
	low    []int
	parent []int
	ap     []bool
}

// NewArticulationPointFinder initializes and returns a new instance of ArticulationPointFinder
func NewArticulationPointFinder(g *unweightedundirectedgraph.UnWeightedUndirectedGraph) *ArticulationPointFinder {
	vertices := g.Vertices
	return &ArticulationPointFinder{
		graph:  g,
		time:   0,
		disc:   make([]int, vertices),
		low:    make([]int, vertices),
		parent: make([]int, vertices),
		ap:     make([]bool, vertices),
	}
}

// FindAP finds the articulation points in the graph and returns them
func (a *ArticulationPointFinder) FindAP() []int {
	// Initialize the necessary arrays
	for i := range a.disc {
		a.disc[i] = -1
		a.low[i] = -1
		a.parent[i] = -1
	}

	// Find articulation points in DFS tree rooted with vertex 'i'
	for i := 0; i < a.graph.Vertices; i++ {
		if a.disc[i] == -1 {
			a.dfs(i)
		}
	}

	// Collecting the articulation points from the boolean array
	var aps []int
	for i, isAP := range a.ap {
		if isAP {
			aps = append(aps, i)
		}
	}
	return aps
}

// dfs is a recursive function that finds articulation points using DFS traversal
func (a *ArticulationPointFinder) dfs(u int) {
	children := 0
	a.disc[u] = a.time
	a.low[u] = a.time
	a.time++

	for _, v := range a.graph.Adjacency[u] {
		if a.disc[v] == -1 {
			children++
			a.parent[v] = u
			a.dfs(v)

			a.low[u] = min(a.low[u], a.low[v])

			// Check for articulation point
			if a.parent[u] == -1 && children > 1 {
				a.ap[u] = true
			}
			if a.parent[u] != -1 && a.low[v] >= a.disc[u] {
				a.ap[u] = true
			}
		} else if v != a.parent[u] {
			a.low[u] = min(a.low[u], a.disc[v])
		}
	}
}

// min returns the minimum of two integers
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
