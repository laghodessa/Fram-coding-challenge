package digraph

func New() Graph {
	return Graph{
		vertices: map[string]struct{}{},
		adj:      map[string][]string{},
	}
}

type Graph struct {
	vertices map[string]struct{} // vertice set
	adj      map[string][]string // adj edges
}

// V returns number of vertices
func (g Graph) V() int {
	return len(g.vertices)
}

func (g Graph) Adj(v string) []string {
	return g.adj[v]
}

func (g Graph) AddEdge(from, to string) {
	g.adj[from] = append(g.adj[from], to)
	g.vertices[from] = struct{}{}
	g.vertices[to] = struct{}{}
}
