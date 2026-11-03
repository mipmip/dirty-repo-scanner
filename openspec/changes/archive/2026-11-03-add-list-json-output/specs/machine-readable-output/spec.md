# Spec Delta

## Purpose

Provides non-interactive output modes that expose the dirty repository scan
result to scripts and status bars as plain paths or structured JSON.

## ADDED Requirements

### Requirement: List output mode

The application SHALL accept a `--list` flag that writes each dirty repository
to stdout as its absolute path on its own line. Lines SHALL NOT contain prefixes,
annotations, or trailing whitespace. No other content SHALL be written to
stdout.

#### Scenario: Dirty repositories are listed

- **WHEN** the scan finds two dirty repositories
- **THEN** stdout contains exactly their two absolute paths, one per line

#### Scenario: Empty result

- **WHEN** no repository is dirty
- **THEN** stdout is empty and the process exits with status 0

#### Scenario: Positional directories

- **WHEN** the user runs `drs --list /some/dir`
- **THEN** only repositories under `/some/dir` are scanned, and the arguments
  notice does not appear on stdout

### Requirement: JSON output mode

The application SHALL accept a `--json` flag that writes a JSON array to stdout.
Each element SHALL be an object with `path` (absolute path string), `vcs`
(`"git"` or `"jj"`), and `status` containing non-negative integer `staged`,
`unstaged`, and `untracked` counts. No other content SHALL be written to stdout.

#### Scenario: Dirty repositories are emitted as JSON

- **WHEN** the scan finds a dirty git repository
- **THEN** stdout is a JSON array with one object whose `path` is the absolute
  repository path, `vcs` is `"git"`, and `status` contains the three counts

#### Scenario: jj repository is identified

- **WHEN** a dirty repository contains a `.jj` directory
- **THEN** its JSON object has `vcs` set to `"jj"`

#### Scenario: Empty result

- **WHEN** no repository is dirty
- **THEN** stdout is `[]` and the process exits with status 0

#### Scenario: Staged, unstaged, and untracked counts

- **WHEN** a repository has one staged modification and one untracked file
- **THEN** its `status` object has `staged` 1, `unstaged` 0, and `untracked` 1

#### Scenario: Count definitions

- **WHEN** a git status entry has a staging code other than space or `?`
- **THEN** it is counted as staged
- **AND** a worktree code other than space or `?` is counted as unstaged
- **AND** a worktree code of `?` is counted as untracked

### Requirement: Stable ordering

Repositories SHALL be ordered by absolute path in ascending byte order in both
`--list` and `--json` output.

#### Scenario: Unsorted scan results are sorted

- **WHEN** the scanner returns repositories in arbitrary order
- **THEN** `--list` prints them sorted by ascending path

#### Scenario: Both modes share the order

- **WHEN** the same scan result is rendered with `--list` and `--json`
- **THEN** the repository paths appear in the same order in both outputs

### Requirement: Output modes do not start the UI

The `--list` and `--json` flags SHALL suppress the interactive TUI. Exactly one
output mode SHALL be selected per invocation.

#### Scenario: List mode suppresses the UI

- **WHEN** the user runs `drs --list`
- **THEN** no TUI is started and the process terminates after writing stdout

#### Scenario: JSON mode suppresses the UI

- **WHEN** the user runs `drs --json`
- **THEN** no TUI is started and the process terminates after writing stdout

#### Scenario: Conflicting output flags

- **WHEN** the user combines `--list` with `--json`, or either with `--debug` or
  `--multiplex`
- **THEN** the application writes a clear error to stderr, writes nothing to
  stdout, and exits with a non-zero status

### Requirement: Diagnostics stay off stdout

In `--list` and `--json` modes, all diagnostics SHALL be written to stderr. Walk
errors SHALL NOT change stdout or the exit status. A scan failure other than a
walk error SHALL exit non-zero and SHALL NOT emit a partial stdout payload.

#### Scenario: Include glob matches nothing

- **WHEN** a configured include glob matches no directory
- **THEN** the walk error goes to stderr, stdout is unchanged, and the process
  exits with status 0

#### Scenario: Walk warning during a successful scan

- **WHEN** the walk logs a warning to stderr while other repositories scan
  successfully
- **THEN** stdout still contains only the valid list or JSON payload

#### Scenario: Fatal scan failure

- **WHEN** repository status cannot be determined
- **THEN** the application writes the error to stderr, writes no payload to
  stdout, and exits with a non-zero status
