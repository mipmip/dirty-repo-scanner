# release-workflow Specification

## Purpose
TBD - created by archiving change add-changelog-release. Update Purpose after archive.
## Requirements
### Requirement: Changelog follows Keep a Changelog format
The project SHALL have a `CHANGELOG.md` following the Keep a Changelog format with an `[Unreleased]` section at the top.

#### Scenario: Changelog structure
- **WHEN** `CHANGELOG.md` is read
- **THEN** it SHALL contain an `[Unreleased]` section and versioned sections with dates

### Requirement: Release script automates versioning
The `scripts/release.sh` script SHALL automate the release process including version bumping, changelog updating, and tag creation.

#### Scenario: Interactive release
- **WHEN** `scripts/release.sh` is run
- **THEN** it SHALL prompt for bump type (patch/minor/major), update VERSION, update CHANGELOG, commit, tag, and push

#### Scenario: Clean working tree required
- **WHEN** the working tree has uncommitted changes
- **THEN** the release script SHALL refuse to proceed

#### Scenario: Works with jj-colocated repos
- **WHEN** the project is managed with jj (colocated with git)
- **THEN** the release script SHALL work correctly using git commands

### Requirement: GoReleaser builds cross-platform binaries
GoReleaser configs SHALL produce binaries for Linux and macOS (amd64 and arm64).

#### Scenario: Linux build
- **WHEN** goreleaser runs with `.goreleaser-linux.yaml`
- **THEN** it SHALL produce `drs` binaries for linux/amd64 and linux/arm64

#### Scenario: macOS build
- **WHEN** goreleaser runs with `.goreleaser-darwin.yaml`
- **THEN** it SHALL produce `drs` binaries for darwin/amd64 and darwin/arm64

### Requirement: GitHub Actions release on tag push
A GitHub Actions workflow SHALL trigger on `v*` tag pushes and run goreleaser.

#### Scenario: Tag push triggers release
- **WHEN** a tag matching `v*` is pushed
- **THEN** the release workflow SHALL build binaries and create a GitHub Release

### Requirement: Release documentation
A `RELEASING.md` file SHALL document the release process and pre-release checklist.

#### Scenario: Documentation exists
- **WHEN** a developer needs to create a release
- **THEN** `RELEASING.md` SHALL describe the steps, prerequisites, and troubleshooting

