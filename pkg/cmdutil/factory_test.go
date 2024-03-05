package cmdutil

import "testing"

// TestProcessURL tests the processURL function.
func TestProcessURL(t *testing.T) {
	tests := []struct {
		url      string
		env      string
		expected string
	}{
		{"http://example.com", "production", "http://example.com"},
		{"http://%s.example.com", "production", "http://production.example.com"},
	}

	for _, test := range tests {
		result := processURL(test.url, test.env)
		if result != test.expected {
			t.Errorf("processURL(%q, %q) = %q; want %q", test.url, test.env, result, test.expected)
		}
	}
}
