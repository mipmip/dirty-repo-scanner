package output

import (
	"encoding/json"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"

	"github.com/mipmip/dirty-repo-scanner/src/scanner"
)

type Status struct {
	Staged    int `json:"staged"`
	Unstaged  int `json:"unstaged"`
	Untracked int `json:"untracked"`
}

type Repo struct {
	Path   string `json:"path"`
	VCS    string `json:"vcs"`
	Status Status `json:"status"`
}

type entry struct {
	key  string
	path string
}

func entries(repos scanner.MultiGitStatus) []entry {
	result := make([]entry, 0, len(repos))
	for key := range repos {
		path := key
		if abs, err := filepath.Abs(key); err == nil {
			path = abs
		}
		result = append(result, entry{key: key, path: path})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].path < result[j].path
	})
	return result
}

func counts(status git.Status) Status {
	var result Status
	for _, file := range status {
		if file == nil {
			continue
		}
		if file.Staging != git.Unmodified && file.Staging != git.Untracked {
			result.Staged++
		}
		if file.Worktree == git.Untracked {
			result.Untracked++
			continue
		}
		if file.Worktree != git.Unmodified {
			result.Unstaged++
		}
	}
	return result
}

func List(w io.Writer, repos scanner.MultiGitStatus) error {
	var out strings.Builder
	for _, e := range entries(repos) {
		out.WriteString(e.path)
		out.WriteByte('\n')
	}
	_, err := io.WriteString(w, out.String())
	return err
}

func JSON(w io.Writer, repos scanner.MultiGitStatus) error {
	out := make([]Repo, 0, len(repos))
	for _, e := range entries(repos) {
		status := repos[e.key]
		vcs := status.VCS
		if vcs == "" {
			vcs = scanner.VCSGit
		}
		out = append(out, Repo{
			Path:   e.path,
			VCS:    vcs,
			Status: counts(status.Status),
		})
	}

	b, err := json.Marshal(out)
	if err != nil {
		return err
	}
	b = append(b, '\n')
	_, err = w.Write(b)
	return err
}
