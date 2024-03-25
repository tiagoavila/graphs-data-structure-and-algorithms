package unweightedundirectedgraph

type Edge struct {
	Src  int // Source vertex
	Dest int // Destination vertex
}

type UnWeightedUndirectedGraph struct {
	Vertices  int           // Number of vertices in the graph
	Adjacency map[int][]int // Adjacency list to store edges
}

func NewGraph(vertices int) *UnWeightedUndirectedGraph {
	return &UnWeightedUndirectedGraph{
		Vertices:  vertices,
		Adjacency: make(map[int][]int),
	}
}

func (g *UnWeightedUndirectedGraph) AddEdge(src, dest int) {
	// Add an undirected edge between src and dest with the given weight
	g.Adjacency[src] = append(g.Adjacency[src], dest)

	// Since it's an undirected graph, we add an edge in the opposite direction as well
	g.Adjacency[dest] = append(g.Adjacency[dest], src)
}
