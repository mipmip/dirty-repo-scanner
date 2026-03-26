## Why

The project was forked from `boyvinall/dirtygit` and has since diverged significantly. The name "dirtygit" no longer reflects ownership and the tool's identity should stand on its own. Renaming to "dirty-repo-scanner" (binary: `drs`) establishes an independent identity with a descriptive name and a short, memorable command.

## What Changes

- **BREAKING**: Go module path changes from `github.com/mipmip/dirtygit` to `github.com/mipmip/dirty-repo-scanner`
- **BREAKING**: Binary name changes from `dirtygit` to `drs`
- **BREAKING**: Default config file changes from `~/.dirtygit.yml` to `~/.dirty-repo-scanner.yml`
- **BREAKING**: Embedded default config file renamed from `.dirtygit.yml` to `.dirty-repo-scanner.yml`
- GitHub repository renamed from `dirtygit` to `dirty-repo-scanner` (GitHub provides automatic redirect)
- All internal import paths updated
- Nix packaging updated (flake.nix, package.nix)
- Documentation updated (README.md)
- `.gitignore` updated for new binary name

## Capabilities

### New Capabilities

- `project-identity`: The naming, module path, binary name, and config file naming of the project

### Modified Capabilities

(none — no existing specs)

## Impact

- **Go module**: `go.mod` module path changes — all internal imports must be updated
- **Binary**: Users must update scripts/aliases referencing `dirtygit` to `drs`
- **Config**: Users must rename `~/.dirtygit.yml` to `~/.dirty-repo-scanner.yml` (or the tool won't find it)
- **Nix**: `flake.nix` and `package.nix` updated with new name
- **GitHub**: Repository URL changes (old URL auto-redirects)
- **Files affected**: `go.mod`, `main.go`, `ui/ui.go`, `package.nix`, `flake.nix`, `.gitignore`, `README.md`, `.dirtygit.yml` → `.dirty-repo-scanner.yml`
