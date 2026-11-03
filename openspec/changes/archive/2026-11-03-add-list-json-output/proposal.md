# Proposal

## Why

mipbar needs to show the number of dirty repositories, a status overview, and a
popover list. It currently has no stable machine-readable source: the only
non-TUI output is `--debug`, which prints timings in a human-oriented format and
is not a contract.

## What Changes

- Add a `--list` flag that writes one absolute dirty repository path per line to
  stdout, sorted, and exits 0 even when nothing is dirty.
- Add a `--json` flag that writes a JSON array of objects to stdout, one object
  per dirty repository, including the absolute path, VCS kind, and staged,
  unstaged, and untracked counts.
- Keep stdout clean in both modes: all diagnostics, including walk errors, go to
  stderr and do not affect stdout or the exit status.
- Never start the TUI in either mode, including when directories are passed as
  arguments.
- Document both modes in `README.md`, add a `CHANGELOG.md` entry, and release a
  new minor version so the mipnix flake input can be bumped.

## Capabilities

### New Capabilities

- `machine-readable-output`: non-interactive CLI output modes that expose the
  dirty repository scan result as plain paths or structured JSON for scripts.

### Modified Capabilities

<!-- None. The interactive TUI and existing flags keep their behavior. -->

## Impact

- `src/main.go`: new flags, output-mode dispatch, conflict handling.
- `src/scanner`: a stable ordering and status summarization path for list and
  JSON output; walk errors must remain non-fatal in these modes.
- New tests for flags, output formats, ordering, empty results, and stderr
  separation.
- `README.md`, `CHANGELOG.md`, `src/VERSION`, and the release produced through
  `scripts/release.sh`.
