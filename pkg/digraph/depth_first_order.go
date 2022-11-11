package digraph

func NewDepthFirstOrder(g Graph) DepthFirstOrder {
	dfo := DepthFirstOrder{
		marked:      map[string]bool{},
		reversePost: nil,
	}
	for v := range g.vertices {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}
	return dfo
}

type DepthFirstOrder struct {
	marked      map[string]bool
	reversePost []string // vertices stack in reverse postorder
}

func (dfo *DepthFirstOrder) dfs(g Graph, v string) {
	dfo.marked[v] = true

	for _, w := range g.Adj(v) {
		if !dfo.marked[w] {
			dfo.dfs(g, w)
		}
	}
	dfo.reversePost = append([]string{v}, dfo.reversePost...)
}

func (dfo DepthFirstOrder) ReversePost() []string {
	return dfo.reversePost
}
