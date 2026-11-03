package output

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"

	"github.com/mipmip/dirty-repo-scanner/src/scanner"
)

func file(staging, worktree git.StatusCode) *git.FileStatus {
	return &git.FileStatus{Staging: staging, Worktree: worktree}
}

func TestList(t *testing.T) {
	t.Run("sorts absolute paths", func(t *testing.T) {
		repos := scanner.MultiGitStatus{
			"/tmp/z/repo": {VCS: scanner.VCSGit},
			"/tmp/a/repo": {VCS: scanner.VCSGit},
			"/tmp/m/repo": {VCS: scanner.VCSGit},
		}

		var buf bytes.Buffer
		if err := List(&buf, repos); err != nil {
			t.Fatalf("List: %v", err)
		}

		want := "/tmp/a/repo\n/tmp/m/repo\n/tmp/z/repo\n"
		if buf.String() != want {
			t.Errorf("List = %q, want %q", buf.String(), want)
		}
	})

	t.Run("empty result", func(t *testing.T) {
		var buf bytes.Buffer
		if err := List(&buf, scanner.MultiGitStatus{}); err != nil {
			t.Fatalf("List: %v", err)
		}
		if buf.String() != "" {
			t.Errorf("List = %q, want empty output", buf.String())
		}
	})

	t.Run("relative path becomes absolute", func(t *testing.T) {
		dir := t.TempDir()
		t.Chdir(dir)

		var buf bytes.Buffer
		if err := List(&buf, scanner.MultiGitStatus{"repo": {VCS: scanner.VCSGit}}); err != nil {
			t.Fatalf("List: %v", err)
		}

		want := filepath.Join(dir, "repo") + "\n"
		if buf.String() != want {
			t.Errorf("List = %q, want %q", buf.String(), want)
		}
	})
}

func TestJSON(t *testing.T) {
	t.Run("git repository", func(t *testing.T) {
		repos := scanner.MultiGitStatus{
			"/tmp/repo": {
				VCS: scanner.VCSGit,
				Status: git.Status{
					"staged.go":    file(git.Modified, git.Unmodified),
					"untracked.go": file(git.Untracked, git.Untracked),
				},
			},
		}

		var buf bytes.Buffer
		if err := JSON(&buf, repos); err != nil {
			t.Fatalf("JSON: %v", err)
		}

		want := `[{"path":"/tmp/repo","vcs":"git","status":{"staged":1,"unstaged":0,"untracked":1}}]` + "\n"
		if buf.String() != want {
			t.Errorf("JSON = %q, want %q", buf.String(), want)
		}
	})

	t.Run("jj repository", func(t *testing.T) {
		repos := scanner.MultiGitStatus{
			"/tmp/jj-repo": {
				VCS: scanner.VCSJJ,
				Status: git.Status{
					"modified.go": file(git.Unmodified, git.Modified),
				},
			},
		}

		var buf bytes.Buffer
		if err := JSON(&buf, repos); err != nil {
			t.Fatalf("JSON: %v", err)
		}

		var got []Repo
		if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
			t.Fatalf("unmarshal JSON: %v", err)
		}
		if len(got) != 1 {
			t.Fatalf("got %d repos, want 1", len(got))
		}
		if got[0].VCS != scanner.VCSJJ {
			t.Errorf("vcs = %q, want %q", got[0].VCS, scanner.VCSJJ)
		}
	})

	t.Run("empty result is an empty array", func(t *testing.T) {
		var buf bytes.Buffer
		if err := JSON(&buf, scanner.MultiGitStatus{}); err != nil {
			t.Fatalf("JSON: %v", err)
		}
		if buf.String() != "[]\n" {
			t.Errorf("JSON = %q, want %q", buf.String(), "[]\n")
		}
	})

	t.Run("count definitions", func(t *testing.T) {
		repos := scanner.MultiGitStatus{
			"/tmp/repo": {
				VCS: scanner.VCSGit,
				Status: git.Status{
					"staged.go":       file(git.Modified, git.Unmodified),
					"unstaged.go":     file(git.Unmodified, git.Modified),
					"untracked.go":    file(git.Untracked, git.Untracked),
					"staged-and-more": file(git.Modified, git.Modified),
				},
			},
		}

		var buf bytes.Buffer
		if err := JSON(&buf, repos); err != nil {
			t.Fatalf("JSON: %v", err)
		}

		var got []Repo
		if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
			t.Fatalf("unmarshal JSON: %v", err)
		}
		want := Status{Staged: 2, Unstaged: 2, Untracked: 1}
		if got[0].Status != want {
			t.Errorf("status = %+v, want %+v", got[0].Status, want)
		}
	})

	t.Run("shares ordering with list", func(t *testing.T) {
		repos := scanner.MultiGitStatus{
			"/tmp/z/repo": {VCS: scanner.VCSGit},
			"/tmp/a/repo": {VCS: scanner.VCSGit},
			"/tmp/m/repo": {VCS: scanner.VCSGit},
		}

		var listBuf bytes.Buffer
		if err := List(&listBuf, repos); err != nil {
			t.Fatalf("List: %v", err)
		}

		var jsonBuf bytes.Buffer
		if err := JSON(&jsonBuf, repos); err != nil {
			t.Fatalf("JSON: %v", err)
		}

		var got []Repo
		if err := json.Unmarshal(jsonBuf.Bytes(), &got); err != nil {
			t.Fatalf("unmarshal JSON: %v", err)
		}

		var jsonPaths []string
		for _, repo := range got {
			jsonPaths = append(jsonPaths, repo.Path)
		}
		if strings.Join(jsonPaths, "\n")+"\n" != listBuf.String() {
			t.Errorf("JSON paths %v do not match list output %q", jsonPaths, listBuf.String())
		}
	})
}
