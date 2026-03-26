## Why

The project root is cluttered with source packages (`scanner/`, `ui/`) mixed alongside config files, Nix files, and documentation. Moving application code under a `src/` directory separates source from project metadata, making the repo easier to navigate.

## What Changes

- Move `scanner/` to `src/scanner/`
- Move `ui/` to `src/ui/`
- Move `main.go` to `src/main.go`
- Update all Go import paths from `dirty-repo-scanner/scanner` → `dirty-repo-scanner/src/scanner`, `dirty-repo-scanner/ui` → `dirty-repo-scanner/src/ui`
- Update Makefile build command to target `src/main.go`
- Move embedded config file `.dirty-repo-scanner.yml` to `src/` (required by `//go:embed` which is relative to the source file)

## Capabilities

### New Capabilities

- `src-layout`: Directory structure convention placing all Go source under `src/`

### Modified Capabilities

(none)

## Impact

- **Go imports**: All internal import paths gain a `src/` segment
- **Build**: Makefile and `go:embed` directive updated for new paths
- **Nix**: `package.nix` may need no changes (builds from module root)
- **Files moved**: `main.go`, `scanner/`, `ui/`, `.dirty-repo-scanner.yml`
