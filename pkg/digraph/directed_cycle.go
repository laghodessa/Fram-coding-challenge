package digraph

func NewDirectedCycle(g Graph) DirectedCycle {
	dc := DirectedCycle{
		onStack: map[string]bool{},
		edgeTo:  map[string]string{},
		marked:  map[string]bool{},
		cycle:   nil,
	}

	for v := range g.vertices {
		if !dc.marked[v] && !dc.HasCycle() {
			dc.dfs(g, v)
		}
	}
	return dc
}

type DirectedCycle struct {
	onStack map[string]bool
	edgeTo  map[string]string
	marked  map[string]bool
	cycle   []string
}

func (dc *DirectedCycle) dfs(g Graph, v string) {
	dc.onStack[v] = true
	dc.marked[v] = true

	for _, w := range g.Adj(v) {
		if dc.HasCycle() {
			return
		}

		if !dc.marked[w] {
			dc.edgeTo[w] = v
			dc.dfs(g, w)
		} else if dc.onStack[w] {
			for x := v; x != w; x = dc.edgeTo[x] {
				dc.cycle = append(dc.cycle, x)
			}
			dc.cycle = append(dc.cycle, w, v)
		}
	}
	dc.onStack[v] = false
}

func (dc DirectedCycle) HasCycle() bool {
	return len(dc.cycle) > 0
}
func (dc DirectedCycle) Cycle() []string {
	return dc.cycle
}
