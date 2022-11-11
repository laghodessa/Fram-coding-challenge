package hr

import (
	"context"
	"errors"
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

type HierarchyRepo interface {
	Update(context.Context, Hierarchy) error
}

var (
	ErrHierarchyHasLoop          = errors.New("hierarchy has loop")
	ErrHierarchyHasMultipleRoots = errors.New("hieracrchy has multiple roots")
)
