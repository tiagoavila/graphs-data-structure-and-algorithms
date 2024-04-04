package eulerian

import (
	unweightedundirectedgraph "graphs_data_structure_and_algorithms/graphs/unweighted_undirected_graph"
)

// https://www.geeksforgeeks.org/eulerian-path-and-circuit/
// Eulerian Path is a path in a graph that visits every edge exactly once.
// Eulerian Circuit is an Eulerian Path that starts and ends on the same vertex.

// Eulerian Cycle: An undirected graph has Eulerian cycle if following two conditions are true.
// 1 - All vertices with non-zero degree are connected. We don’t care about vertices with zero degree because they don’t belong to Eulerian Cycle or Path (we only consider all edges).
// 2 - All vertices have even degree.

// Eulerian Path: An undirected graph has Eulerian Path if following two conditions are true.
// 1 - Same as condition (1) for Eulerian Cycle.
// 2 - If zero or two vertices have odd degree and all other vertices have even degree. Note that only one vertex with odd degree is not possible in an undirected graph (sum of all degrees is always even in an undirected graph)

// isEulerian checks the Eulerian status of the graph.
// Returns 0 if the graph is non-Eulerian,
// 1 if the graph has an Euler path (semi-Eulerian),
// and 2 if the graph has an Euler Cycle (Eulerian).
func IsEulerian(g *unweightedundirectedgraph.UnWeightedUndirectedGraph) int {
	if !IsConnected(g) {
		return 0 // Non-Eulerian
	}

	odd := 0 // count of vertices with odd degree

	// Check the degree of each vertex
	for i := 0; i < g.Vertices; i++ {
		if len(g.Adjacency[i])%2 != 0 {
			odd++
		}
	}

	// Based on the count of odd degree vertices, determine the Eulerian status
	if odd > 2 {
		return 0 // Non-Eulerian
	} else if odd == 2 {
		return 1 // Semi-Eulerian
	} else {
		return 2 // Eulerian
	}
}

func IsConnected(g *unweightedundirectedgraph.UnWeightedUndirectedGraph) bool {
	// Mark all the vertices as not visited
	visited := make([]bool, g.Vertices)

	// Find a vertex with non-zero degree
	var i int
	for i = 0; i < g.Vertices; i++ {
		if len(g.Adjacency[i]) != 0 {
			break
		}
	}

	// If there are no edges in the graph, return true
	if i == g.Vertices {
		return true
	}

	// Start DFS traversal from a vertex with non-zero degree
	dfs(g, i, visited)

	// Check if all non-zero degree vertices are visited
	for i = 0; i < g.Vertices; i++ {
		if !visited[i] && len(g.Adjacency[i]) > 0 {
			return false
		}
	}

	return true
}

func dfs(g *unweightedundirectedgraph.UnWeightedUndirectedGraph, u int, visited []bool) {
	visited[u] = true
	for _, v := range g.Adjacency[u] {
		if !visited[v] {
			dfs(g, v, visited)
		}
	}
}
