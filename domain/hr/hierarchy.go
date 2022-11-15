package hr

import (
	"context"
	"personia/domain"
	"personia/pkg/digraph"
)

func NewHierarchy(req map[string]string) (Hierarchy, error) {
	g := buildHierarchyGraph(req)
	loop := checkHierarchyLoop(g)
	roots := checkHierarchyMultipleRoots(g, req)
	if loop != nil || roots != nil {
		return nil, NewInvalidHierarchyError(loop, roots)
	}
	return Hierarchy(req), nil
}

func buildHierarchyGraph(req map[string]string) digraph.Graph {
	g := digraph.New()
	for empl, sup := range req {
		g.AddEdge(empl, sup)
	}
	return g
}

func checkHierarchyLoop(g digraph.Graph) []string {
	dc := digraph.NewDirectedCycle(g)
	if dc.HasCycle() {
		return dc.Cycle()
	}
	return nil
}

func checkHierarchyMultipleRoots(g digraph.Graph, req map[string]string) []string {
	cc := digraph.NewConnectedComponent(g)
	if cc.Count() < 2 {
		return nil
	}

	roots := make([]string, cc.Count())
	for i := 0; i < cc.Count(); i++ {
		comp := cc.Component(i)
		sup := comp[0]
		marked := make(map[string]struct{}, len(comp))

		for req[sup] != "" { // while employee still has supervisor
			marked[sup] = struct{}{}
			sup = req[sup] // get supervisor of this employee
			if _, ok := marked[sup]; ok {
				// loop detected, so there won't be a root supervisor
				sup = "?"
				break
			}
		}
		roots[i] = sup
	}
	return roots
}

type Hierarchy map[string]string

func (h Hierarchy) SupervisorOf(name string) string {
	return h[name]
}

// Topology returns values from top to bottom of the hierarchy
func (h Hierarchy) Topology() []string {
	g := digraph.New()
	for k, v := range h {
		g.AddEdge(v, k)
	}
	dfo := digraph.NewDepthFirstOrder(g)
	order := dfo.ReversePost()
	return order
}

//go:generate go run github.com/golang/mock/mockgen -package=hrmock -destination=./hrmock/mock.go . HierarchyRepo
type HierarchyRepo interface {
	Get(context.Context) (Hierarchy, error)
	Update(context.Context, Hierarchy) error
}

func NewInvalidHierarchyError(loop, roots []string) *domain.Error {
	meta := make(map[string]interface{}, 2)
	if loop != nil {
		meta["loop"] = loop
	}
	if roots != nil {
		meta["roots"] = roots
	}

	return &domain.Error{
		Code:    "hierarchy_invalid",
		Message: "hierarchy can't contains loop or multiple roots",
		Meta:    meta,
	}
}
