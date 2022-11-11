package hr_test

import (
	"personia/domain/hr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHierarchy_Topology(t *testing.T) {
	cases := []struct {
		name     string
		hier     hr.Hierarchy
		expected []string
	}{
		{
			name: "it returns most senior employee first",
			hier: map[string]string{
				"Pete":   "Nick",
				"Nick":   "Sophie",
				"Sophie": "Jonas",
			},
			expected: []string{"Jonas", "Sophie", "Nick", "Pete"},
		},
		{
			name:     "when the hierarchy is empty, it returns nil slice",
			hier:     map[string]string{},
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.hier.Topology()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestHierarchy_SupervisorOf(t *testing.T) {
	cases := []struct {
		name     string
		hier     hr.Hierarchy
		req      string
		expected string
	}{
		{
			name: "it returns the supervisor",
			hier: hr.Hierarchy{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Nick":    "Sophie",
				"Sophie":  "Jonas",
			},
			req:      "Nick",
			expected: "Sophie",
		},
		{
			name:     "it returns empty string on not found",
			hier:     map[string]string{},
			req:      "John",
			expected: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.hier.SupervisorOf(tc.req)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewHierarchy(t *testing.T) {
	cases := []struct {
		name     string
		req      map[string]string
		err      error
		expected hr.Hierarchy
	}{
		{
			name: "it returns hierarchy",
			req: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Nick":    "Sophie",
				"Sophie":  "Jonas",
			},
			expected: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Nick":    "Sophie",
				"Sophie":  "Jonas",
			},
		},
		{
			name:     "when request is empty, it returns empty hierarchy",
			req:      map[string]string{},
			expected: map[string]string{},
		},
		{
			name: "when request contains loop, it fails",
			req: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Nick":    "Sophie",
				"Sophie":  "Pete",
			},
			err: hr.ErrHierarchyHasLoop,
		},
		{
			name: "when request contains multiple roots, it fails",
			req: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Sophie":  "Jonas",
			},
			err: hr.ErrHierarchyHasMultipleRoots,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hier, err := hr.NewHierarchy(tc.req)

			if tc.err != nil {
				assert.ErrorIs(t, err, tc.err, "match error")
			} else {
				assert.Equal(t, tc.expected, hier)
			}
		})
	}
}
