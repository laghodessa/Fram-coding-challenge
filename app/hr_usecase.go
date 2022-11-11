package app

import (
	"context"
	"fmt"
	"personia/domain/hr"
)

type HRUC struct {
	HierarchyRepo hr.HierarchyRepo
}

func (uc *HRUC) UpdateHierachy(ctx context.Context, req map[string]string) error {
	hier, err := hr.NewHierarchy(req)
	if err != nil {
		return err
	}

	if err := uc.HierarchyRepo.Update(ctx, hier); err != nil {
		return fmt.Errorf("update hierarchy: %w", err)
	}
	return nil
}

type GetSupervisorResp struct {
	Supervisor            string `json:"supervisor,omitempty"`
	SupervisorsSupervisor string `json:"supervisorsSupervisor,omitempty"`
}

func (uc *HRUC) GetSupervisor(ctx context.Context, name string) (GetSupervisorResp, error) {
	hier, err := uc.HierarchyRepo.Get(ctx)
	if err != nil {
		return GetSupervisorResp{}, err
	}

	sup := hier.SupervisorOf(name)
	return GetSupervisorResp{
		Supervisor:            sup,
		SupervisorsSupervisor: hier.SupervisorOf(sup),
	}, nil
}

func (uc *HRUC) GetHierarchy(ctx context.Context) (Hierarchy, error) {
	hier, err := uc.HierarchyRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return uc.FormatHierarchy(hier), nil
}

func (uc *HRUC) FormatHierarchy(hier hr.Hierarchy) Hierarchy {
	supervised := map[string]Hierarchy{} // supervised[employee] = whom the employee supervise?
	topo := hier.Topology()
	if len(topo) == 0 {
		return Hierarchy{}
	}

	root := topo[0]
	for _, name := range hier.Topology() {
		supervised[name] = Hierarchy{}

		sup := supervised[hier.SupervisorOf(name)]
		if sup != nil {
			sup[name] = supervised[name]
		}
	}
	return Hierarchy{
		root: supervised[root],
	}
}
