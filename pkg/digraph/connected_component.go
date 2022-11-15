package digraph

func NewConnectedComponent(g Graph) ConnectedComponent {
	cc := ConnectedComponent{
		marked: map[string]bool{},
		id:     map[string]int{},
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
	id     map[string]int // id[v] = component id of vertex v, id starts from 0
	count  int            // number of connected components
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
	cc.id[v] = cc.count

	for _, w := range undirected.Adj(v) {
		if !cc.marked[w] {
			cc.dfs(undirected, w)
		}
	}
}

func (cc ConnectedComponent) Component(id int) []string {
	r := make([]string, 0, len(cc.id))
	for v, comp := range cc.id {
		if comp == id {
			r = append(r, v)
		}
	}
	return r
}

func (cc ConnectedComponent) Count() int {
	return cc.count
}
