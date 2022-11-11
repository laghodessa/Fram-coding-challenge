package digraph

func NewConnectedComponent(g Graph) ConnectedComponent {
	cc := ConnectedComponent{
		marked: map[string]bool{},
		count:  0,
	}
	for v := range g.vertices {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}
	return cc
}

type ConnectedComponent struct {
	marked map[string]bool
	count  int // number of connected components
}

func (cc *ConnectedComponent) dfs(digraph Graph, v string) {
	undirected := New()
	for k, v := range digraph.adj {
		for _, e := range v {
			undirected.AddEdge(k, e)
			undirected.AddEdge(e, k)
		}
	}
	cc.marked[v] = true

	for _, w := range undirected.Adj(v) {
		if !cc.marked[w] {
			cc.dfs(undirected, w)
		}
	}
}

func (cc ConnectedComponent) Count() int {
	return cc.count
}
