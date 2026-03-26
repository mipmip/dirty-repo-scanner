# scanner-tests Specification

## Purpose
TBD - created by archiving change add-testing-framework. Update Purpose after archive.
## Requirements
### Requirement: Excluder path matching
The `Excluder.IsExcluded` function SHALL correctly identify files and directories that match configured exclusion patterns.

#### Scenario: File glob matches basename
- **WHEN** a file glob pattern `*.pyc` is configured and path `src/main.pyc` is checked
- **THEN** `IsExcluded` SHALL return `true`

#### Scenario: File glob does not match
- **WHEN** a file glob pattern `*.pyc` is configured and path `src/main.go` is checked
- **THEN** `IsExcluded` SHALL return `false`

#### Scenario: Dir glob matches directory component
- **WHEN** a dir glob pattern `vendor` is configured and path `vendor/lib/file.go` is checked
- **THEN** `IsExcluded` SHALL return `true`

#### Scenario: Dir glob does not match
- **WHEN** a dir glob pattern `vendor` is configured and path `src/lib/file.go` is checked
- **THEN** `IsExcluded` SHALL return `false`

#### Scenario: No patterns configured
- **WHEN** no file or dir glob patterns are configured
- **THEN** `IsExcluded` SHALL return `false` for any path

### Requirement: Git status filtering
The `Excluder.FilterGitStatus` function SHALL remove excluded files from a git status map.

#### Scenario: Excluded files are removed
- **WHEN** git status contains entries for `go.sum` and `main.go`, and `go.sum` matches a file glob
- **THEN** `FilterGitStatus` SHALL return a status containing only `main.go`

#### Scenario: No files excluded
- **WHEN** git status contains entries that match no exclusion patterns
- **THEN** `FilterGitStatus` SHALL return all entries unchanged

#### Scenario: All files excluded
- **WHEN** all git status entries match exclusion patterns
- **THEN** `FilterGitStatus` SHALL return an empty status

### Requirement: Config parsing
The `ParseConfigFile` function SHALL parse YAML config files and fall back to default config when the file does not exist.

#### Scenario: Valid config file
- **WHEN** a valid YAML config file is provided with `scandirs.include` entries
- **THEN** `ParseConfigFile` SHALL return a `Config` with those include paths populated

#### Scenario: File does not exist falls back to default
- **WHEN** the config file path does not exist and a default config string is provided
- **THEN** `ParseConfigFile` SHALL parse the default config string and return a valid `Config`

#### Scenario: Invalid YAML
- **WHEN** the config file contains invalid YAML
- **THEN** `ParseConfigFile` SHALL return an error

### Requirement: Skip helper
The `skip` function SHALL match paths against an exclusion list using full-path or basename comparison.

#### Scenario: Full path match
- **WHEN** the exclusion list contains `/home/user/node_modules` and the needle is `/home/user/node_modules`
- **THEN** `skip` SHALL return `true`

#### Scenario: Basename match
- **WHEN** the exclusion list contains `node_modules` (no leading `/`) and the needle is `/home/user/project/node_modules`
- **THEN** `skip` SHALL return `true`

#### Scenario: No match
- **WHEN** the exclusion list contains `vendor` and the needle is `/home/user/project/src`
- **THEN** `skip` SHALL return `false`

### Requirement: Make test target
The Makefile SHALL include a `test` target that runs Go tests.

#### Scenario: Running make test
- **WHEN** `make test` is executed
- **THEN** it SHALL run `go test ./...` and exit with the test result code

