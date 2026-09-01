## 1. VCS detection in the scanner

- [x] 1.1 Add a `VCS` field (kind `git` | `jj`) to `RepoStatus` in `src/scanner/scan.go`
- [x] 1.2 In the walk/scan, stat for a sibling `.jj` in the repo root and set `VCS` accordingly; ensure the walker never descends into `.jj`
- [x] 1.3 Keep dirtiness detection on `git status --porcelain` for all repos (no jj command during the scan)

## 2. VCS marker in the list

- [x] 2.1 Add lipgloss styles for a jj and a git marker
- [x] 2.2 Prefix each row in `renderRepoList` with the colored VCS marker

## 3. jj display helpers

- [x] 3.1 Add a helper that runs `jj diff -s --color never` in a repo and parses `^<KIND> <path>$` into a file list (KIND in A/M/D/R/C)
- [x] 3.2 Add a helper that runs `jj diff --git --color never -- <file>` in a repo and returns the diff text
- [x] 3.3 Add a `jj` availability check (`exec.LookPath`) with a one-time warning and git fallback

## 4. Wire jj into the TUI display

- [x] 4.1 When opening a repo with `VCS == jj` (and jj available), source the file list from the jj summary helper instead of the git status map
- [x] 4.2 In `fetchDiff`, branch on `VCS == jj` to use the jj diff helper; keep the git path unchanged
- [x] 4.3 Render the single jj change char in the file row (no staging column) while leaving git's two-column rendering intact

## 5. Tests

- [x] 5.1 Scanner: `.jj` sibling → `VCS == jj`; `.git` only → `VCS == git`
- [x] 5.2 Scanner: dirtiness still comes from git; no jj invocation during scan
- [x] 5.3 UI: repo list renders the correct VCS marker per kind
- [x] 5.4 jj summary parsing: `A/M/D` lines map to the expected file list
- [x] 5.5 Diff path selection: jj repo uses jj diff, git repo uses git diff (stubbed exec)
- [x] 5.6 Missing-jj fallback: warning logged and git display used

## 6. Docs & verification

- [x] 6.1 README: document colocated-only jj support and the native-jj caveat
- [x] 6.2 Run `gofmt -l .` (empty), `go test ./...`, and `nix flake check`; confirm all pass
