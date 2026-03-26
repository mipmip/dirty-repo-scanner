## ADDED Requirements

### Requirement: VERSION file as single source of truth
The project SHALL have a `src/VERSION` file containing the current version string.

#### Scenario: VERSION file exists
- **WHEN** the project is checked out
- **THEN** `src/VERSION` SHALL contain a semver version string (e.g., `0.2.0`)

### Requirement: Version embedded in binary
The version from `src/VERSION` SHALL be embedded into the binary at compile time via `//go:embed`.

#### Scenario: Development build
- **WHEN** the binary is built with `make build`
- **THEN** `drs --version` SHALL display the version from `src/VERSION`

#### Scenario: Release build
- **WHEN** the binary is built by goreleaser with ldflags
- **THEN** the ldflags-injected version SHALL override the embedded version

### Requirement: Nix flake reads version from file
The `flake.nix` SHALL read the version from `src/VERSION` rather than hardcoding it.

#### Scenario: Nix version consistency
- **WHEN** `src/VERSION` is updated
- **THEN** `nix build` SHALL produce a package with the matching version
