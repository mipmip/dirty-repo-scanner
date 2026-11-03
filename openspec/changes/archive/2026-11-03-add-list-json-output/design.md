# Design

## Context

See `proposal.md` for motivation and `specs/machine-readable-output/spec.md` for
the behavior contract.

Current CLI handling lives in `src/main.go` with `urfave/cli`. `scanner.Scan`
returns `MultiGitStatus`, a `map[string]RepoStatus` keyed by absolute repository
path. `RepoStatus` embeds `git.Status`, a map of file path to `FileStatus` with
`Staging` and `Worktree` codes, plus `VCS` and `ScanTime`.

Two existing behaviors conflict with clean stdout:

- When positional directories are given, `src/main.go` prints
  `Arguments given, skipping config` with `fmt.Println`, which writes to stdout.
- `scanner.Walk` calls `log.Fatal` when the parent of an include glob cannot be
  read, which terminates the process instead of reporting a walk error.

`log` output already goes to stderr. Walk errors are otherwise controlled by
`ignore_dir_errors`, which defaults to true.

## Goals / Non-Goals

**Goals:**

- A stable, documented machine-readable contract with deterministic output.
- stdout contains only the payload in `--list` and `--json` modes.
- Output encoding is testable without running the binary or a TUI.
- Existing interactive behavior is unchanged.

**Non-Goals:**

- Ahead/behind counts. The scanner does not compute them today, so the JSON
  contract does not invent them.
- Native-only jj working copies without `.git`. Scanning support stays as it is.
- Changing the `--debug` format or the TUI.
- Changing how dirtiness is detected: the same `git status` result is used.

## Decisions

### Output encoding lives in a new `src/output` package

Add pure functions that accept `scanner.MultiGitStatus` plus an `io.Writer` and
encode either the list or the JSON payload. `main.go` only selects the mode and
calls the encoder.

Alternative considered: encode inline in `main.go`. Rejected because asserting
exact stdout and ordering through `app.Run` requires capturing process output,
which is slower and more brittle than testing a writer directly.

### JSON schema: path, vcs, and three counts

Each object is:

```json
{
  "path": "/absolute/repo",
  "vcs": "git",
  "status": {"staged": 1, "unstaged": 0, "untracked": 2}
}
```

Counts are derived from `git.FileStatus` codes:

- `staged`: staging code is neither space nor `?`
- `unstaged`: worktree code is neither space nor `?`
- `untracked`: worktree code is `?`

The categories can overlap for a file that is both staged and modified, which is
why they are three independent counts instead of a partition.

`ScanTime` is deliberately omitted because it makes output nondeterministic and
is not status information.

Alternative considered: include every file path and its two status codes.
Rejected as a larger contract than mipbar needs and expensive to keep stable;
fields can be added later without breaking consumers.

### One order for both modes

Collect map keys, apply `sort.Strings`, and use that slice for both encoders.
This matches the TUI's ordering and makes `--list` and `--json` agree.

### Buffer the payload before writing

Encode into a buffer and write it to stdout only after the scan and encoding
succeed. This guarantees no partial payload when a status lookup fails, as the
spec requires.

### Route diagnostics to stderr

Move the positional-arguments notice to stderr, or suppress it in list and JSON
modes. Keep `log` at its default stderr destination. In list and JSON modes,
call the scanner with walk errors ignored even when `--ignore_dir_errors=false`,
because the spec makes walk errors non-fatal for these modes.

Replace `log.Fatal` in `scanner.Walk` with a returned error so a missing glob
parent becomes a normal walk error that the existing `ignore_dir_errors` path
can report to stderr. This also removes a hard process exit from the TUI path.

### Resolve mode conflicts explicitly

Count the selected modes among `--list`, `--json`, `--debug`, and `--multiplex`.
More than one is an error with no stdout payload and a non-zero exit. `--list`
and `--json` otherwise bypass TUI startup.

### Make errors exit non-zero

`main` currently prints the returned error and exits 0. Change it to print to
stderr and call `os.Exit(1)`. The multiplex spec already promises a non-zero
exit for its configuration error, so this aligns the implementation with the
existing contract.

### Release as a minor version

The change is additive and backward compatible, so the release is `0.4.0`.
`CHANGELOG.md` entries go under `[Unreleased]` first; `scripts/release.sh`
moves them under the new version header, updates `src/VERSION`, updates the
`package.nix` vendor hash, commits, tags, and pushes.

## Risks / Trade-offs

- **JSON becomes a public contract** → Document the shape in `README.md` and
  allow additive fields; changing or removing fields is a breaking change.
- **Counts come from git status, even for jj repos** → This is consistent with
  how dirtiness is detected today. The README already explains that limitation,
  and the `vcs` field lets consumers distinguish the cases.
- **Forcing walk errors to be ignored in output modes** → A misconfigured include
  could silently produce an empty list. Mitigation: the error is still printed to
  stderr, and the exit status stays 0 as the contract requires.
- **Changing all errors to exit non-zero** → Scripts relying on the old exit 0
  behavior would notice. The old behavior was inconsistent with documented
  multiplex behavior and with normal CLI expectations, so the fix is intentional.
- **`json.Marshal` of a nil slice yields `null`** → Initialize the result slice
  with `make(..., 0, len(...))` so an empty scan emits `[]`.

## Migration Plan

No migration is required. The new flags are opt-in and the TUI remains the
default. Rollback is removing the flags and the output package.

## Open Questions

None.
