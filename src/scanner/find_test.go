package scanner

import (
	"context"
	"path/filepath"
	"testing"
)

func TestSkip(t *testing.T) {
	tests := []struct {
		name     string
		needle   string
		haystack []string
		expected bool
	}{
		{
			name:     "full path match",
			needle:   "/home/user/node_modules",
			haystack: []string{"/home/user/node_modules"},
			expected: true,
		},
		{
			name:     "basename match",
			needle:   "/home/user/project/node_modules",
			haystack: []string{"node_modules"},
			expected: true,
		},
		{
			name:     "no match",
			needle:   "/home/user/project/src",
			haystack: []string{"vendor"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := skip(tt.needle, tt.haystack)
			if got != tt.expected {
				t.Errorf("skip(%q, %v) = %v, want %v", tt.needle, tt.haystack, got, tt.expected)
			}
		})
	}
}

func TestWalkMissingGlobParent(t *testing.T) {
	config := &Config{}
	config.ScanDirs.Include = []string{filepath.Join(t.TempDir(), "missing", "*")}

	t.Run("ignored", func(t *testing.T) {
		err := Walk(context.Background(), config, make(chan string, 1), true)
		if err != nil {
			t.Errorf("Walk = %v, want nil when errors are ignored", err)
		}
	})

	t.Run("returned", func(t *testing.T) {
		err := Walk(context.Background(), config, make(chan string, 1), false)
		if err == nil {
			t.Error("Walk = nil, want error for missing glob parent")
		}
	})
}
