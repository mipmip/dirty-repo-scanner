package scanner

import (
	"testing"

	"github.com/go-git/go-git/v5"
)

func TestIsExcluded(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		dirs     []string
		path     string
		expected bool
	}{
		{
			name:     "file glob matches basename",
			files:    []string{"*.pyc"},
			path:     "src/main.pyc",
			expected: true,
		},
		{
			name:     "file glob does not match",
			files:    []string{"*.pyc"},
			path:     "src/main.go",
			expected: false,
		},
		{
			name:     "dir glob matches directory component",
			dirs:     []string{"vendor"},
			path:     "vendor/lib/file.go",
			expected: true,
		},
		{
			name:     "dir glob does not match",
			dirs:     []string{"vendor"},
			path:     "src/lib/file.go",
			expected: false,
		},
		{
			name:     "no patterns configured",
			path:     "anything/file.go",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, err := NewExcluder(tt.files, tt.dirs)
			if err != nil {
				t.Fatalf("NewExcluder: %v", err)
			}
			got := ex.IsExcluded(tt.path)
			if got != tt.expected {
				t.Errorf("IsExcluded(%q) = %v, want %v", tt.path, got, tt.expected)
			}
		})
	}
}

func TestFilterGitStatus(t *testing.T) {
	tests := []struct {
		name      string
		files     []string
		dirs      []string
		status    git.Status
		wantPaths []string
	}{
		{
			name:  "excluded files are removed",
			files: []string{"go.sum"},
			status: git.Status{
				"go.sum": &git.FileStatus{Staging: git.Modified},
				"main.go": &git.FileStatus{Staging: git.Modified},
			},
			wantPaths: []string{"main.go"},
		},
		{
			name: "no files excluded",
			status: git.Status{
				"main.go": &git.FileStatus{Staging: git.Modified},
				"lib.go":  &git.FileStatus{Staging: git.Added},
			},
			wantPaths: []string{"main.go", "lib.go"},
		},
		{
			name:  "all files excluded",
			files: []string{"*.sum", "*.pyc"},
			status: git.Status{
				"go.sum":  &git.FileStatus{Staging: git.Modified},
				"app.pyc": &git.FileStatus{Staging: git.Modified},
			},
			wantPaths: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex, err := NewExcluder(tt.files, tt.dirs)
			if err != nil {
				t.Fatalf("NewExcluder: %v", err)
			}
			got := ex.FilterGitStatus(tt.status)

			// Check that only expected paths remain
			for _, p := range tt.wantPaths {
				if _, ok := got[p]; !ok {
					t.Errorf("expected path %q in result, but not found", p)
				}
			}

			// Count non-nil entries (FilterGitStatus may leave nil values in the map)
			nonNil := 0
			for _, v := range got {
				if v != nil {
					nonNil++
				}
			}
			if nonNil != len(tt.wantPaths) {
				t.Errorf("got %d non-nil entries, want %d", nonNil, len(tt.wantPaths))
			}
		})
	}
}
