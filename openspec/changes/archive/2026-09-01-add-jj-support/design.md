## Context

dirty-repo-scanner is git-shaped end to end: `find.go` `Walk` discovers repos by
triggering on a `.git` directory; `scan.go` `Scan` runs `git status --porcelain`
per repo and keeps the non-clean ones as `RepoStatus` (embedding go-git's
`git.Status`, a `map[file]{Staging,Worktree}`); the TUI (`ui.go`) prints the bare
path in the list, renders the file list from that status map (`" %c%c  file"`),
and fetches per-file diffs with `git diff` (`fetchDiff`).

Jujutsu (jj) working copies come in three shapes: pure git (`.git` only),
**colocated** (`.git` + `.jj`), and native (`.jj` only). Only colocated repos are
in scope here — they already surface in the scan (because `.git` exists) but are
indistinguishable from git and use git for display. jj 0.41 output was verified
against a colocated repo: `jj diff -s --color never` yields clean `M path` /
`A path` lines, and `jj diff --git --color never` yields a standard git-format
unified diff (so the existing `colorizeDiff` needs no changes).

## Goals / Non-Goals

**Goals:**
- Mark colocated jj repos distinctly in the list.
- Show a jj repo's modified files and diffs via `jj` when the user opens it.
- Keep the fleet-wide scan strictly read-only and fresh.
- Degrade gracefully (warn + git fallback) when `jj` is missing.

**Non-Goals:**
- Native (non-colocated) jj discovery.
- A VCS-agnostic status model / interface refactor.
- jj concepts beyond the working-copy change set (descriptions, bookmarks,
  unpushed commits).

## Decisions

### Split detection (git) from display (jj)
The core decision. jj only snapshots the working copy when a jj command runs, so
`jj diff --ignore-working-copy` can report a freshly edited repo as **clean** —
fatal for a dirty scanner. Plain `jj diff` is fresh but **mutates** (snapshots
`@`, imports git refs); doing that across every repo in a sweep is an unacceptable
side effect. Because colocated repos always have `.git`, we sidestep both:

- **Detection (scan, all repos):** keep `git status --porcelain` — fresh,
  read-only, fast, already implemented. Preserves the scanner's no-mutation
  guarantee across the fleet.
- **Display (one repo, on open):** use `jj diff -s` (file list) and
  `jj diff --git` (diffs). Any snapshot side effect is scoped to the single repo
  the user actively opened — exactly when they would run jj themselves.

Alternatives rejected: `--ignore-working-copy` everywhere (stale → misses the
repos the tool exists to catch); plain `jj diff` for detection (mass mutation).

### Minimal data model: a `VCS` field, not an abstraction
Add a `VCS` kind (`git` | `jj`) to `RepoStatus`; keep embedding `git.Status` as
the transport. The scan still populates the git status map (used for the git
marker and the git display path). jj display is fetched lazily on open, mirroring
the existing lazy `fetchDiff`. This avoids a status-model rewrite; the only jj
branch points are the list marker, the file-list source, and the diff command.

### Detection point: stat `.jj` during the walk
`Walk` continues to trigger on `.git`; when it records a repo root it also stats
for a sibling `.jj` to set the kind. No change to which repos are discovered
(colocated repos already have `.git`), so native jj remains out of scope by
construction. The walk must not descend into `.jj`.

### jj invocation details (grounded in jj 0.41)
- File list: `jj diff -s --color never` → parse `^<KIND> <path>$` (KIND in
  `A M D R C`). Map to the file-list row; jj has no staging column, so render a
  single change char rather than git's two.
- Diff: `jj diff --git --color never -- <file>` → git-format, reuse `colorizeDiff`.
- Always pass `--color never` (jj colorizes by default); run with `cmd.Dir` set to
  the repo root, consistent with the existing git exec calls.

### Missing-jj fallback: warn, then git
When `VCS == jj` but `exec.LookPath("jj")` fails, log one warning and use the
existing git display path. Safe because colocated repos always have `.git`.

## Risks / Trade-offs

- **Detection (git) and display (jj) can disagree** — git compares worktree↔HEAD,
  jj compares `@`↔`@-`. For everyday working-copy edits they match; a repo mid-jj
  workflow could show a git dirty flag with a differing jj file set → Mitigation:
  document as a known colocated-only corner; do not abstract the model now.
- **On-open jj snapshot mutates one repo** → Mitigation: scoped to the repo the
  user explicitly opened; equivalent to them running `jj diff` by hand.
- **jj output format drift across versions** → Mitigation: parse the stable
  `-s` summary and rely on `--git` for diffs; pin behavior with tests using a
  stubbed jj command rather than a live binary.
- **`jj diff -s` file set differs from the git status map the list was built on**
  → Mitigation: the jj list is authoritative once a jj repo is opened; the scan's
  git-derived set only drives the dirty flag and marker.
