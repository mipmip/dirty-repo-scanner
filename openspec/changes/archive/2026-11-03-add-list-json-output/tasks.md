# Tasks

## 1. Output encoders

- [x] 1.1 Add `src/output` with `List(io.Writer, scanner.MultiGitStatus)` and
  `JSON(io.Writer, scanner.MultiGitStatus)` that sort repositories by absolute
  path and write only the payload. Verify `go test ./src/output` passes.
- [x] 1.2 Add table tests for list output covering sorted paths, empty output,
  and absence of annotations. Verify the tests fail if ordering or extra text is
  introduced.
- [x] 1.3 Add JSON tests covering `path`, `vcs` (`git` and `jj`), the three
  counts, empty result as `[]`, valid JSON parsing, and the shared ordering with
  list output. Verify `go test ./src/output` passes.

## 2. CLI wiring

- [x] 2.1 Add `--list` and `--json` boolean flags in `src/main.go`, dispatch to
  the encoders, buffer the payload, and keep the TUI from starting in either
  mode. Verify `go run ./src --list` and `go run ./src --json` print only the
  payload.
- [x] 2.2 Move the positional-arguments notice off stdout and keep all
  diagnostics on stderr in output modes. Verify `go run ./src --list <dir>`
  produces parseable output with no notice on stdout.
- [x] 2.3 Detect conflicting modes among `--list`, `--json`, `--debug`, and
  `--multiplex`, and fail with a stderr message, empty stdout, and a non-zero
  exit. Verify `drs --list --json` exits non-zero with empty stdout.
- [x] 2.4 Make `main` report errors to stderr and exit non-zero. Verify a command
  with an invalid flag or missing `switch_command` returns a non-zero status.

## 3. Walker error handling

- [x] 3.1 Replace the `log.Fatal` for an unreadable include-glob parent in
  `src/scanner/find.go` with a returned error so it flows through the existing
  `ignore_dir_errors` path. Verify `go test ./src/scanner` passes and a missing
  glob parent does not terminate the test process.
- [x] 3.2 In `--list` and `--json` modes, ignore walk errors regardless of
  `--ignore_dir_errors` and keep the exit status at 0. Verify an unmatched
  include glob logs to stderr while stdout stays valid and the exit status is 0.

## 4. Documentation

- [x] 4.1 Add a README section documenting `--list`, `--json`, the JSON object
  shape, stderr behavior, exit statuses, and that `--debug` is not a contract.
  Verify the documented commands produce the documented shapes.
- [x] 4.2 Add a `CHANGELOG.md` entry under `[Unreleased]` describing both modes
  and the release version. Verify `grep` finds the entry.

## 5. Integration verification

- [x] 5.1 Run `go test ./...` and `go build ./...`. Verify both succeed.
- [x] 5.2 Run an end-to-end check against the current config: `drs --list` is
  sorted and contains absolute paths, `drs --json | jq` succeeds, and scanning an
  empty directory returns empty output with exit status 0.
- [x] 5.3 Confirm `drs --list --json`, `drs --list --debug`, and
  `drs --json --multiplex` all exit non-zero, write nothing to stdout, and
  explain the conflict on stderr.

## Workflow follow-up

- Archive the change after review and verification.
- After the implementation is committed on a clean tree, run
  `scripts/release.sh`, choose the minor bump to `0.4.0`, and verify the tag,
  `CHANGELOG.md` header, and `drs --version`.
- Bump the mipnix flake input to the new release and verify mipbar sees the dirty
  repository count.
