## Why

The project has no versioning, changelog, or release process. As the project matures past its fork origins, it needs a proper release workflow. Version 0.1.0 is considered virtual (covering all changes since the fork). This change establishes v0.2.0 as the first explicit release.

## What Changes

- Add `src/VERSION` file containing `0.2.0`, embedded into the binary via `//go:embed`
- Set `app.Version` in urfave/cli so `--version` works automatically
- Add `CHANGELOG.md` following Keep a Changelog format, with 0.2.0 and virtual 0.1.0 sections
- Add `RELEASING.md` documenting the release process
- Add `scripts/release.sh` — interactive release script using `gum` (supports both git and jj-colocated repos)
- Add `.goreleaser-linux.yaml` and `.goreleaser-darwin.yaml` for cross-platform binary builds
- Add `.github/workflows/release.yml` GitHub Actions workflow triggered by `v*` tags
- Update `flake.nix` to read version from `src/VERSION`

## Capabilities

### New Capabilities

- `versioning`: VERSION file, go:embed, CLI --version flag
- `release-workflow`: release script, goreleaser configs, GitHub Actions, changelog, releasing docs

### Modified Capabilities

(none)

## Impact

- **Code**: `src/main.go` — add version embed and `app.Version`
- **Build**: goreleaser configs for Linux/macOS builds
- **CI**: New GitHub Actions release workflow
- **Nix**: `flake.nix` reads version from `src/VERSION`
- **New files**: `src/VERSION`, `CHANGELOG.md`, `RELEASING.md`, `scripts/release.sh`, `.goreleaser-linux.yaml`, `.goreleaser-darwin.yaml`, `.github/workflows/release.yml`
