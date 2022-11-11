package app_test

import (
	"context"
	"personia/app"
	"personia/domain/hr"
	"personia/domain/hr/hrmock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHRUC_GetSupervisor(t *testing.T) {
	cases := []struct {
		name               string
		setupHierarchyRepo func(repo *hrmock.MockHierarchyRepo)
		req                string
		expected           app.GetSupervisorResp
	}{
		{
			name: "it returns supervisors",
			setupHierarchyRepo: func(repo *hrmock.MockHierarchyRepo) {
				repo.EXPECT().Get(gomock.Any()).Return(hr.Hierarchy{
					"Pete":    "Nick",
					"Barbara": "Nick",
					"Nick":    "Sophie",
					"Sophie":  "Jonas",
				}, nil)
			},
			req: "Pete",
			expected: app.GetSupervisorResp{
				Supervisor:            "Nick",
				SupervisorsSupervisor: "Sophie",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hierarchyRepo := hrmock.NewMockHierarchyRepo(ctrl)
			tc.setupHierarchyRepo(hierarchyRepo)

			hrUC := app.HRUC{
				HierarchyRepo: hierarchyRepo,
			}

			actual, err := hrUC.GetSupervisor(context.Background(), tc.req)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestHRUC_FormatHierarchy(t *testing.T) {
	cases := []struct {
		name      string
		hierarchy hr.Hierarchy
		expected  app.Hierarchy
	}{
		{
			name: "it puts the most senior employee at the top",
			hierarchy: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Nick":    "Sophie",
				"Sophie":  "Jonas",
			},
			expected: app.Hierarchy{
				"Jonas": app.Hierarchy{
					"Sophie": app.Hierarchy{
						"Nick": app.Hierarchy{
							"Pete":    app.Hierarchy{},
							"Barbara": app.Hierarchy{},
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			uc := app.HRUC{}

			actual := uc.FormatHierarchy(tc.hierarchy)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
