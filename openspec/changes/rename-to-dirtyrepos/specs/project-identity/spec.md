## ADDED Requirements

### Requirement: Go module name
The Go module path SHALL be `github.com/mipmip/dirty-repo-scanner`.

#### Scenario: Module declaration
- **WHEN** inspecting `go.mod`
- **THEN** the module line reads `module github.com/mipmip/dirty-repo-scanner`

### Requirement: Binary name
The compiled binary SHALL be named `drs`. The application SHALL identify itself as `dirty-repo-scanner` in CLI output.

#### Scenario: Binary file name
- **WHEN** the project is built with `go build`
- **THEN** the output binary is named `drs`

#### Scenario: App name in help
- **WHEN** the user runs `drs --help`
- **THEN** the output shows `dirty-repo-scanner` as the application name

### Requirement: Default config file path
The application SHALL look for its default config at `~/.dirty-repo-scanner.yml`.

#### Scenario: Default config location
- **WHEN** the user runs the application without `--config`
- **THEN** the application reads from `~/.dirty-repo-scanner.yml`

### Requirement: Embedded default config
The application SHALL embed a default config file named `.dirty-repo-scanner.yml` from the project root.

#### Scenario: Embedded config exists
- **WHEN** no config file is found at the default path
- **THEN** the application falls back to the embedded `.dirty-repo-scanner.yml` defaults

### Requirement: Internal imports use new module path
All Go source files SHALL import internal packages using the `github.com/mipmip/dirty-repo-scanner` module path.

#### Scenario: Scanner import
- **WHEN** inspecting `main.go` imports
- **THEN** the scanner import reads `github.com/mipmip/dirty-repo-scanner/scanner`

#### Scenario: UI import
- **WHEN** inspecting `main.go` imports
- **THEN** the UI import reads `github.com/mipmip/dirty-repo-scanner/ui`

### Requirement: Nix package name
The Nix package SHALL be named `dirty-repo-scanner`.

#### Scenario: Flake output
- **WHEN** inspecting `flake.nix` packages output
- **THEN** the package is named `dirty-repo-scanner`

#### Scenario: Package derivation
- **WHEN** inspecting `package.nix`
- **THEN** `pname` is set to `"dirty-repo-scanner"`
