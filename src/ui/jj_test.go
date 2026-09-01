package ui

import (
	"reflect"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"

	"github.com/mipmip/dirty-repo-scanner/src/scanner"
)

func TestParseJJSummary(t *testing.T) {
	t.Run("normal output", func(t *testing.T) {
		out := "A added.txt\nM tracked.txt\n"
		jl := parseJJSummary(out)
		wantPaths := []string{"added.txt", "tracked.txt"}
		if !reflect.DeepEqual(jl.paths, wantPaths) {
			t.Errorf("paths = %#v, want %#v", jl.paths, wantPaths)
		}
		if jl.kinds["added.txt"] != 'A' || jl.kinds["tracked.txt"] != 'M' {
			t.Errorf("kinds = %#v", jl.kinds)
		}
	})

	t.Run("skips informational banner", func(t *testing.T) {
		out := "Done importing changes from the underlying Git repo.\nM foo.go\n"
		jl := parseJJSummary(out)
		if !reflect.DeepEqual(jl.paths, []string{"foo.go"}) {
			t.Errorf("paths = %#v, want [foo.go]", jl.paths)
		}
	})

	t.Run("empty output", func(t *testing.T) {
		jl := parseJJSummary("")
		if len(jl.paths) != 0 {
			t.Errorf("paths = %#v, want empty", jl.paths)
		}
	})
}

func TestVcsTag(t *testing.T) {
	if got := vcsTag(scanner.VCSJJ); got != "jj " {
		t.Errorf("vcsTag(jj) = %q, want %q", got, "jj ")
	}
	if got := vcsTag(scanner.VCSGit); got != "git" {
		t.Errorf("vcsTag(git) = %q, want %q", got, "git")
	}
}

func TestRenderRepoListShowsVCSMarker(t *testing.T) {
	m := model{
		height:    40,
		repoPaths: []string{"/home/u/alpha", "/home/u/beta"},
		cursor:    0,
		repositories: scanner.MultiGitStatus{
			"/home/u/alpha": {VCS: scanner.VCSJJ},
			"/home/u/beta":  {VCS: scanner.VCSGit},
		},
	}
	got := m.renderRepoList(80)
	if !strings.Contains(got, "jj") {
		t.Errorf("expected a jj marker in output: %q", got)
	}
	if !strings.Contains(got, "git") {
		t.Errorf("expected a git marker in output: %q", got)
	}
}

func TestRenderFileListJJKinds(t *testing.T) {
	// A jj repo whose git status map does NOT contain the jj paths must still
	// render (single change char) without a nil dereference.
	m := model{
		height:     40,
		activeView: viewStatus,
		repoPaths:  []string{"/home/u/repo"},
		cursor:     0,
		filePaths:  []string{"changed.go"},
		fileCursor: 0,
		repositories: scanner.MultiGitStatus{
			"/home/u/repo": {VCS: scanner.VCSJJ},
		},
		jjFileKinds: map[string]byte{"changed.go": 'M'},
	}
	got := m.renderFileList(40, 10)
	if !strings.Contains(got, "changed.go") {
		t.Errorf("output missing file: %q", got)
	}
	if !strings.Contains(got, "M") {
		t.Errorf("output missing jj change kind: %q", got)
	}
}

func diffModel(vcs string, jjOK bool, worktree git.StatusCode) model {
	return model{
		repoPaths:  []string{"/nonexistent/repo"},
		filePaths:  []string{"foo.txt"},
		fileCursor: 0,
		cursor:     0,
		jjOK:       jjOK,
		repositories: scanner.MultiGitStatus{
			"/nonexistent/repo": {
				Status: git.Status{"foo.txt": &git.FileStatus{Staging: git.Unmodified, Worktree: worktree}},
				VCS:    vcs,
			},
		},
	}
}

func TestFetchDiffPathSelection(t *testing.T) {
	// git repo + untracked file → git-specific "Untracked file" message.
	gitCmd := diffModel(scanner.VCSGit, false, git.Untracked).fetchDiff()
	if gitCmd == nil {
		t.Fatal("git fetchDiff returned nil")
	}
	if msg, ok := gitCmd().(diffMsg); !ok || msg.content != "Untracked file" {
		t.Errorf("git path: got %#v, want diffMsg{Untracked file}", gitCmd())
	}

	// jj repo bypasses the git untracked branch entirely (runs jj diff, which on
	// a nonexistent repo yields "No diff available").
	jjCmd := diffModel(scanner.VCSJJ, true, git.Untracked).fetchDiff()
	if jjCmd == nil {
		t.Fatal("jj fetchDiff returned nil")
	}
	if msg, ok := jjCmd().(diffMsg); !ok || msg.content == "Untracked file" {
		t.Errorf("jj path should not use the git untracked branch, got %#v", jjCmd())
	}
}

func TestJJEnsureFallbackWhenMissing(t *testing.T) {
	// With PATH cleared, jj cannot be found; jjEnsure reports false (git fallback).
	t.Setenv("PATH", "")
	m := &model{}
	if m.jjEnsure() {
		t.Error("jjEnsure should be false when jj is not on PATH")
	}
	if !m.jjChecked {
		t.Error("jjEnsure should memoize the check")
	}
}
