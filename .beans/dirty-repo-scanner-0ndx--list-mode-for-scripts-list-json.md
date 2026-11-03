---
# dirty-repo-scanner-0ndx
title: list mode for scripts (--list / --json)
status: completed
type: feature
created_at: 2026-11-03T13:45:13Z
updated_at: 2026-11-03T17:40:30Z
---

mipbar (in mipnix, bean mipnix-gq1b) gets a bar icon with the number of dirty
repos, a tooltip with a status overview and a popover listing the repos. It
scans once an hour and again when the popover opens. It needs machine-readable
output from drs.

Today the only non-TUI output is `--debug` (src/main.go): it prints
`%-40s %v` per dirty repo, path plus scan time, and logs walk errors and
timings to stderr. That is a debug format and should not be a contract.

Wanted:

- `--list`: one absolute repo path per line on stdout, nothing else, sorted.
  Exit 0 also when there are no dirty repos (empty output).
- `--json`: an array of objects, at least `path`, plus whatever status drs
  already knows per repo (for example modified/untracked/ahead counts, jj or
  git), so the mipbar tooltip and popover can show more than a name later.
- Walk errors (for example a glob in `include` that matches nothing) go to
  stderr and do not change stdout or the exit code, the same as
  `ignore_dir_errors` does now.
- No UI starts in either mode.
- Document both in README and CHANGELOG, and release a new version so the
  mipnix flake input can be bumped.

Context: a full scan of the current config takes about 1 s
(`walkDuration` ~950 ms, `statusDuration` ~60 ms) and finds 5 dirty repos.

## Summary of Changes

Added two non-interactive output modes backed by a new `src/output` package.

- `--list` prints sorted absolute paths, one per line, with no other stdout
  content and exit 0 for an empty result.
- `--json` prints an array of `{path, vcs, status}` objects, where `status`
  contains staged, unstaged, and untracked counts.
- Both modes bypass the TUI, buffer output, send diagnostics and walk errors to
  stderr, and keep walk errors non-fatal.
- Conflicting modes exit non-zero with nothing on stdout; all CLI errors now
  exit non-zero.
- README and CHANGELOG document the modes; OpenSpec change
  `add-list-json-output` archived as `2026-11-03-add-list-json-output`.
