package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectVCS(t *testing.T) {
	t.Run("colocated jj", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Mkdir(filepath.Join(dir, ".git"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.Mkdir(filepath.Join(dir, ".jj"), 0o755); err != nil {
			t.Fatal(err)
		}
		if got := detectVCS(dir); got != VCSJJ {
			t.Errorf("detectVCS = %q, want %q", got, VCSJJ)
		}
	})

	t.Run("plain git", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.Mkdir(filepath.Join(dir, ".git"), 0o755); err != nil {
			t.Fatal(err)
		}
		if got := detectVCS(dir); got != VCSGit {
			t.Errorf("detectVCS = %q, want %q", got, VCSGit)
		}
	})

	t.Run("jj file (not a directory) is not jj", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, ".jj"), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
		if got := detectVCS(dir); got != VCSGit {
			t.Errorf("detectVCS = %q, want %q", got, VCSGit)
		}
	})
}
