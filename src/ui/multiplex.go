package ui

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// renderCommand substitutes the per-repo placeholders in a command template:
// %WORKING_DIRECTORY becomes the repo's path and %REPO_NAME its basename.
func renderCommand(tmpl, repoPath string) string {
	out := strings.ReplaceAll(tmpl, "%WORKING_DIRECTORY", repoPath)
	out = strings.ReplaceAll(out, "%REPO_NAME", filepath.Base(repoPath))
	return out
}

// splitArgs splits a command string into argv using POSIX-ish shell-quoting
// rules (single quotes, double quotes, backslash escapes) WITHOUT invoking a
// shell, so substituted repo paths cannot inject shell behaviour.
func splitArgs(s string) ([]string, error) {
	var args []string
	var cur strings.Builder
	inArg := false

	flush := func() {
		if inArg {
			args = append(args, cur.String())
			cur.Reset()
			inArg = false
		}
	}

	for i := 0; i < len(s); {
		c := s[i]
		switch c {
		case ' ', '\t', '\n', '\r':
			flush()
			i++
		case '\'':
			inArg = true
			i++
			for i < len(s) && s[i] != '\'' {
				cur.WriteByte(s[i])
				i++
			}
			if i >= len(s) {
				return nil, errors.New("unterminated single quote")
			}
			i++ // closing quote
		case '"':
			inArg = true
			i++
			for i < len(s) && s[i] != '"' {
				if s[i] == '\\' && i+1 < len(s) {
					switch s[i+1] {
					case '"', '\\', '$', '`':
						cur.WriteByte(s[i+1])
						i += 2
						continue
					}
				}
				cur.WriteByte(s[i])
				i++
			}
			if i >= len(s) {
				return nil, errors.New("unterminated double quote")
			}
			i++ // closing quote
		case '\\':
			if i+1 >= len(s) {
				return nil, errors.New("trailing backslash")
			}
			inArg = true
			cur.WriteByte(s[i+1])
			i += 2
		default:
			inArg = true
			cur.WriteByte(c)
			i++
		}
	}
	flush()
	return args, nil
}

// runArgv executes argv directly (no shell), wired to the user's terminal. It
// is the default Model.runCommand; tests override it with a recorder.
func runArgv(argv []string) error {
	if len(argv) == 0 {
		return errors.New("empty command")
	}
	c := exec.Command(argv[0], argv[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
