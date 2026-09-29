package cmdutil

import (
	"snd-cli/pkg/cmd/urls"
	"testing"
)

// TestProcessURL tests the ProcessURL function.
func TestProcessURL(t *testing.T) {
	tests := []struct {
		url      string
		env      string
		expected string
	}{
		{"https://example.com", "production", "https://example.com"},
		{urls.NexusUrl, "production", "https://nexus-production.snd-production.io"},
		{urls.NexusUrl, "awsd", "https://nexus-dev-0.snd-awsd.io"},
		{urls.NexusUrl, "awsp", "https://nexus-production-0.snd-awsp.io"},
		{urls.BoxerURL, "awsd", "https://boxer-dev-0.snd-awsd.io/api/v1"},
		{urls.BoxerURL, "awsp", "https://boxer-production-0.snd-awsp.io/api/v1"},
	}

	for _, test := range tests {
		result := ProcessURL(test.url, test.env)
		if result != test.expected {
			t.Errorf("ProcessURL(%q, %q) = %q; want %q", test.url, test.env, result, test.expected)
		}
	}
}

func TestProcessBeastURL(t *testing.T) {
	tests := []struct {
		url      string
		env      string
		expected string
	}{
		{"https://beast.sneaksanddata.com", "production", "https://beast.sneaksanddata.com"},
		{urls.BeastURL, "production", "https://beast-production.snd-awsp.io"},
		{urls.BeastURL, "awsp", "https://beast-production-0.snd-awsp.io"},
		{urls.BeastURL, "awsd", "https://beast-dev-0.snd-awsp.io"},
	}

	for _, test := range tests {
		result := processBeastURL(test.url, test.env)
		if result != test.expected {
			t.Errorf("processBeastURL(%q, %q) = %q; want %q", test.url, test.env, result, test.expected)
		}
	}
}
