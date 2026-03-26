## Context

The `scanner` package has three files (`scan.go`, `find.go`, `excluded.go`) with no tests. The package contains a mix of pure logic (config parsing, exclusion matching, status filtering) and I/O-heavy code (filesystem walking, git operations). The pure logic is straightforward to test; the I/O code requires real git repos or mocks.

## Goals / Non-Goals

**Goals:**
- Unit tests for pure logic in the scanner package
- A `make test` target for running tests
- Tests that run fast and don't require network or special setup

**Non-Goals:**
- Integration tests that create real git repos on disk (future work)
- Testing the `ui` package (TUI testing is complex, low ROI for now)
- CI pipeline setup
- Code coverage targets

## Decisions

### 1. Test only pure functions first

Test `Excluder.IsExcluded`, `Excluder.FilterGitStatus`, `ParseConfigFile`, and `skip`. These are pure functions with clear inputs/outputs. Skip testing `Scan`, `Walk`, `GitStatus`, `GoGitStatus` — they require filesystem and git setup.

**Why**: Maximum test value with minimum complexity. The I/O functions can be tested later with integration tests.

### 2. Use standard library `testing` only

No external test frameworks (testify, gomega, etc.). Use table-driven tests with `t.Run`.

**Why**: Zero dependencies, idiomatic Go, sufficient for these tests.

### 3. Test `ParseConfigFile` with temp files

Use `os.CreateTemp` to write YAML to disk, then parse it. Also test the default-config fallback path (when file doesn't exist).

**Why**: `ParseConfigFile` reads from disk, so we need real files. Temp files are cheap and cleaned up automatically.

## Risks / Trade-offs

- **[Limited coverage]** Only testing pure logic means git interaction bugs won't be caught. → Acceptable for a first pass; integration tests are a natural follow-up.
- **[FilterGitStatus depends on go-git types]** Tests need to construct `git.Status` and `git.FileStatus` values. → These are simple map/struct types, easy to construct in tests.
