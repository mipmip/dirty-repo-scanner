package ui

import (
	"os/exec"
	"sort"
	"strings"
)

// jjList is the modified-file set of a jj working copy: sorted paths plus a
// per-path change kind (one of A/M/D/R/C).
type jjList struct {
	paths []string
	kinds map[string]byte
}

// parseJJSummary parses `jj diff -s` output. Each relevant line is
// "<KIND> <path>"; informational banners (e.g. "Done importing changes ...")
// are skipped because their second byte is not a space or their leading byte is
// not a known change kind.
func parseJJSummary(out string) jjList {
	jl := jjList{kinds: map[string]byte{}}
	for _, line := range strings.Split(out, "\n") {
		if len(line) < 3 || line[1] != ' ' {
			continue
		}
		kind := line[0]
		if !strings.ContainsRune("AMDRC", rune(kind)) {
			continue
		}
		path := strings.TrimSpace(line[2:])
		if path == "" {
			continue
		}
		jl.paths = append(jl.paths, path)
		jl.kinds[path] = kind
	}
	sort.Strings(jl.paths)
	return jl
}

// jjSummary runs `jj diff -s` in repo and returns the parsed modified-file set.
func jjSummary(repo string) jjList {
	cmd := exec.Command("jj", "diff", "-s", "--color", "never")
	cmd.Dir = repo
	out, err := cmd.Output()
	if err != nil {
		return jjList{kinds: map[string]byte{}}
	}
	return parseJJSummary(string(out))
}

// jjDiff runs `jj diff --git` for a single file. The --git flag yields
// git-format output so the existing diff colorizer renders it unchanged.
func jjDiff(repo, file string) string {
	cmd := exec.Command("jj", "diff", "--git", "--color", "never", "--", file)
	cmd.Dir = repo
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return ""
	}
	return string(out)
}
