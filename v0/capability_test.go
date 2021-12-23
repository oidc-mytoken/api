package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fail(t *testing.T, expected, got Capabilities) {
	t.Errorf("Expected '%v', got '%v'", expected, got)
}

func testTighten(t *testing.T, a, b, expected Capabilities) {
	intersect := TightenCapabilities(a, b)
	assert.Len(t, intersect, len(expected))
	assert.Equal(t, expected, intersect)
	// for i, ee := range expected {
	// 	if ee != intersect[i] {
	// 		fail(t, expected, intersect)
	// 	}
	// }
}

func TestTighten(t *testing.T) {
	tests := []struct {
		name     string
		a        Capabilities
		b        Capabilities
		expected Capabilities
	}{
		{
			name:     "All Empty",
			a:        Capabilities{},
			b:        Capabilities{},
			expected: nil,
		},
		{
			name:     "First Empty",
			a:        Capabilities{},
			b:        NewCapabilities([]string{"not", "empty"}),
			expected: nil,
		},
		{
			name:     "Second Empty",
			a:        NewCapabilities([]string{"not", "empty"}),
			b:        Capabilities{},
			expected: nil,
		},
		{
			name:     "Same",
			a:        NewCapabilities([]string{"some", "values"}),
			b:        NewCapabilities([]string{"some", "values"}),
			expected: NewCapabilities([]string{"some", "values"}),
		},
		{
			name:     "No Intersection",
			a:        NewCapabilities([]string{"some", "values"}),
			b:        NewCapabilities([]string{"completely", "different"}),
			expected: nil,
		},
		{
			name:     "Some Intersection",
			a:        NewCapabilities([]string{"some", "values"}),
			b:        NewCapabilities([]string{"some", "different"}),
			expected: NewCapabilities([]string{"some"}),
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				testTighten(t, test.a, test.b, test.expected)
			},
		)
	}
}

func TestCapabilities_Has(t *testing.T) {
	tests := []struct {
		name     string
		cap      Capabilities
		c        Capability
		expected bool
	}{
		{
			name:     "Empty wanted",
			cap:      AllCapabilities,
			c:        Capability{},
			expected: false,
		},
		{
			name:     "Empty slice",
			cap:      Capabilities{},
			c:        CapabilityAT,
			expected: false,
		},
		{
			name:     "normal",
			cap:      AllCapabilities,
			c:        CapabilityAT,
			expected: true,
		},
		{
			name: "scoped 1",
			cap: Capabilities{
				CapabilitySettings,
			},
			c:        CapabilityGrants,
			expected: true,
		},
		{
			name: "scoped 2",
			cap: Capabilities{
				CapabilityGrants,
			},
			c:        CapabilitySettings,
			expected: false,
		},
		{
			name: "scoped 3",
			cap: Capabilities{
				CapabilityGrants,
			},
			c:        CapabilityGrants,
			expected: true,
		},
		{
			name: "Read / Write 1",
			cap: Capabilities{
				CapabilityGrants,
			},
			c:        CapabilityGrantsRead,
			expected: true,
		},
		{
			name: "Read / Write 2",
			cap: Capabilities{
				CapabilityGrantsRead,
			},
			c:        CapabilityGrants,
			expected: false,
		},
		{
			name: "Scoped Read / Write",
			cap: Capabilities{
				CapabilitySettings,
			},
			c:        CapabilityGrantsRead,
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				assert.Equal(t, test.expected, test.cap.Has(test.c))
			},
		)
	}
}
