package hr_test

import (
	"personia/domain"
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
		err      *domain.Error
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
			err: hr.NewInvalidHierarchyError([]string{"Pete", "Sophie", "Nick", "Pete"}, nil),
		},
		{
			name: "when request contains multiple roots, it fails",
			req: map[string]string{
				"Pete":    "Nick",
				"Barbara": "Nick",
				"Sophie":  "Jonas",
			},
			err: hr.NewInvalidHierarchyError(nil, []string{"Nick", "Jonas"}),
		},
		{
			name: "when request contains both loop & multiple roots, it fails",
			req: map[string]string{
				"Nick":  "Jonas",
				"Jonas": "Nick",
				"Bell":  "Elsa",
			},
			err: hr.NewInvalidHierarchyError([]string{"Nick", "Jonas", "Nick"}, []string{"Elsa", "?"}),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hier, err := hr.NewHierarchy(tc.req)

			if tc.err != nil {
				derr := err.(*domain.Error)
				assert.Equal(t, tc.err.Code, derr.Code, "match error code")
				assert.Equal(t, tc.err.Message, derr.Message, "match error message")

				if loop := tc.err.Meta["loop"]; loop != nil {
					actual := derr.Meta["loop"].([]string)
					expected := loop.([]string)

					start := expected[0]
					offset := 0
					for i, e := range actual {
						if e == start {
							offset = i
						}
					}

					actualShifted := make([]string, len(actual))
					for i := range actual {
						pos := (offset + i) % (len(actual) - 1)
						actualShifted[i] = actual[pos]
					}
					assert.Equal(t, expected, actualShifted)
				} else {
					assert.NotContains(t, tc.err.Meta, "loop")
				}

				if roots := tc.err.Meta["roots"]; roots != nil {
					assert.ElementsMatch(t, roots, derr.Meta["roots"])
				} else {
					assert.NotContains(t, tc.err.Meta, "roots")
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, hier)
			}
		})
	}
}
