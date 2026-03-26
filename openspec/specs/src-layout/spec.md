# src-layout Specification

## Purpose
TBD - created by archiving change refactor-to-src. Update Purpose after archive.
## Requirements
### Requirement: Go source under src directory
All Go source files (`main.go`, `scanner/`, `ui/`) SHALL reside under the `src/` directory.

#### Scenario: Directory structure after refactor
- **WHEN** the project root is listed
- **THEN** there SHALL be a `src/` directory containing `main.go`, `scanner/`, `ui/`, and `.dirty-repo-scanner.yml`
- **THEN** there SHALL NOT be a `main.go`, `scanner/`, or `ui/` at the project root

### Requirement: Build from src
The Makefile `build` target SHALL compile the binary from `src/`.

#### Scenario: Running make build
- **WHEN** `make build` is executed
- **THEN** it SHALL produce a working `drs` binary at the project root

### Requirement: Tests pass after move
All existing tests SHALL continue to pass after the directory restructure.

#### Scenario: Running make test
- **WHEN** `make test` is executed
- **THEN** all tests SHALL pass

### Requirement: Import paths updated
All Go import paths SHALL use the `src/` segment (e.g., `dirty-repo-scanner/src/scanner`).

#### Scenario: No references to old import paths
- **WHEN** the codebase is searched for `dirty-repo-scanner/scanner` or `dirty-repo-scanner/ui` (without `src/`)
- **THEN** no matches SHALL be found in Go source files

