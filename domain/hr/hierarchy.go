package hr

import (
	"context"
	"personia/domain"
	"personia/pkg/digraph"
)

func NewHierarchy(req map[string]string) (Hierarchy, error) {
	g := digraph.New()
	for empl, sup := range req {
		g.AddEdge(empl, sup)
	}

	dc := digraph.NewDirectedCycle(g)
	if dc.HasCycle() {
		return nil, ErrHierarchyHasLoop
	}

	cc := digraph.NewConnectedComponent(g)
	if cc.Count() > 1 {
		return nil, ErrHierarchyHasMultipleRoots
	}
	return Hierarchy(req), nil
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

var (
	ErrHierarchyHasLoop          = domain.NewError("hierarchy_has_loop", "hierarchy has loop")
	ErrHierarchyHasMultipleRoots = domain.NewError("hierarchy_has_multi_roots", "hieracrchy has multiple roots")
)
