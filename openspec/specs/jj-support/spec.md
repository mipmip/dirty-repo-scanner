# jj-support Specification

## Purpose
TBD - created by archiving change add-jj-support. Update Purpose after archive.
## Requirements
### Requirement: Detect colocated jj working copies
The scanner SHALL classify each discovered repository as a jj working copy when a
`.jj` directory exists in the repository root, otherwise as a git repository, and
SHALL record this VCS kind on the repository's status.

#### Scenario: Colocated jj repo
- **WHEN** a scanned repository root contains both `.git` and `.jj`
- **THEN** the repository SHALL be recorded with VCS kind `jj`

#### Scenario: Plain git repo
- **WHEN** a scanned repository root contains `.git` but no `.jj`
- **THEN** the repository SHALL be recorded with VCS kind `git`

#### Scenario: Native jj is out of scope
- **WHEN** a directory contains `.jj` but no `.git`
- **THEN** the scanner SHALL NOT discover it (native-only jj is unsupported)

### Requirement: Read-only dirtiness detection
The scanner SHALL determine whether a repository is dirty using
`git status --porcelain` for every repository, including colocated jj
repositories, and SHALL NOT run any command that snapshots or mutates a working
copy during the scan.

#### Scenario: jj repo dirtiness via git
- **WHEN** a colocated jj repository is scanned
- **THEN** its dirty state SHALL be computed from `git status --porcelain`
- **AND** no `jj` command that snapshots the working copy SHALL run during the scan

#### Scenario: Clean repo excluded
- **WHEN** `git status --porcelain` reports no changes for a repository
- **THEN** the repository SHALL be excluded from the results regardless of VCS kind

### Requirement: VCS marker in the repository list
The repository list SHALL display a visible marker indicating each repository's
VCS kind.

#### Scenario: jj marker
- **WHEN** the list renders a repository with VCS kind `jj`
- **THEN** the row SHALL include a distinct jj marker alongside the repository path

#### Scenario: git marker
- **WHEN** the list renders a repository with VCS kind `git`
- **THEN** the row SHALL include a distinct git marker alongside the repository path

### Requirement: jj file list on open
The modified-file list SHALL be sourced from `jj diff -s` when the user opens a
repository whose VCS kind is `jj`, rather than from the git status map.

#### Scenario: Modified files from jj
- **WHEN** a jj repository's status view is opened
- **THEN** the file list SHALL reflect the paths and change kinds reported by
  `jj diff -s` for that repository

### Requirement: jj per-file diff
When the user views a file in a repository whose VCS kind is `jj`, the diff SHALL
be produced by `jj diff --git -- <file>` so the output is git-format and renders
with the existing diff colorizer.

#### Scenario: jj diff rendering
- **WHEN** a file is selected in a jj repository's status view
- **THEN** the diff panel SHALL show the output of `jj diff --git -- <file>` for
  that file

#### Scenario: git repo unchanged
- **WHEN** a file is selected in a git repository's status view
- **THEN** the diff SHALL be produced by `git diff` exactly as before

### Requirement: Fallback when jj is unavailable
The application SHALL log a warning and fall back to git-based display when a
repository is classified as `jj` but the `jj` executable is not available.

#### Scenario: Missing jj binary
- **WHEN** a colocated jj repository is opened and no `jj` executable is found on
  the PATH
- **THEN** the application SHALL log a warning
- **AND** SHALL display the repository's files and diffs using git

