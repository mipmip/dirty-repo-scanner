## Why

More developers drive their git repositories with [Jujutsu (jj)](https://github.com/jj-vcs/jj).
A jj working copy that is *colocated* with git (`.git` + `.jj`) already appears in
the scan, but the tool silently treats it as plain git — the list gives no hint
it is a jj repo, and the file list / diff use `git` even though the user thinks
in jj terms. This change makes colocated jj repos first-class: mark them in the
list and show their changes with `jj`.

## What Changes

- Detect whether a scanned repo is a jj working copy by checking for a sibling
  `.jj` directory next to `.git`, and record a `VCS` kind (`git` | `jj`) on each
  repo's status.
- Show a colored VCS marker in the repository list (e.g. `◆ jj` / `● git`) so
  the kind is visible at a glance.
- Keep dirtiness **detection** on `git status --porcelain` for every repo,
  including colocated jj repos. It is fresh, read-only, and fleet-safe — the
  scanner never snapshots or mutates working copies during a sweep.
- For **display** of a jj repo (only when the user opens it):
  - build the modified-file list from `jj diff -s` instead of the git status map;
  - render per-file diffs with `jj diff --git -- <file>` (git-format output, so
    the existing diff colorizer is unchanged).
- If a repo has `.jj` but the `jj` binary is not installed, log a **warning** and
  fall back to git display (safe, since colocated repos always have `.git`).
- Document the scope: only **colocated** jj repos are supported; native-only jj
  working copies (no `.git`) are not scanned yet.

Non-goals: native (non-colocated) jj discovery; a VCS-agnostic status model;
jj-specific concepts beyond the working-copy change set (descriptions, bookmarks,
unpushed commits).

## Capabilities

### New Capabilities
- `jj-support`: recognize colocated jj working copies, mark them in the list, and
  render their modified files and diffs via `jj` while keeping dirtiness
  detection on read-only `git status`.

### Modified Capabilities
<!-- None. git detection, the diff colorizer, and tmux-popup-editor behavior are unchanged. -->

## Impact

- **Scanner** (`src/scanner/find.go`, `src/scanner/scan.go`): detect a sibling
  `.jj` during the walk; add a `VCS` field to `RepoStatus`; detection stays on
  `git status --porcelain`.
- **TUI** (`src/ui/ui.go`): colored VCS marker in `renderRepoList`; when
  `VCS == jj`, source the file list from `jj diff -s` and diffs from
  `jj diff --git`; warn-and-fallback when `jj` is absent.
- **Docs** (`README.md`): note colocated-only support and the native-jj caveat.
- **Tests** (`src/scanner`, `src/ui`): `.jj` detection, VCS marker rendering,
  jj file-list parsing, jj diff path selection, and missing-binary fallback
  (using temp repos / a stubbed jj command).
- No new Go module dependency; jj is an external runtime tool invoked via `exec`.
