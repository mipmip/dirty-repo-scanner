## 1. Version Embedding

- [x] 1.1 Create `src/VERSION` file containing `0.2.0`
- [x] 1.2 Add `//go:embed VERSION` and `version` variable to `src/main.go`
- [x] 1.3 Set `app.Version = version` in `src/main.go`

## 2. Changelog and Docs

- [x] 2.1 Create `CHANGELOG.md` with `[Unreleased]`, `[0.2.0]`, and virtual `[0.1.0]` sections
- [x] 2.2 Create `RELEASING.md` documenting the release process

## 3. Release Script

- [x] 3.1 Create `scripts/release.sh` — interactive release script with gum, supporting git/jj-colocated repos

## 4. GoReleaser

- [x] 4.1 Create `.goreleaser-linux.yaml` for linux/amd64 and linux/arm64 (CGO_ENABLED=0)
- [x] 4.2 Create `.goreleaser-darwin.yaml` for darwin/amd64 and darwin/arm64 (CGO_ENABLED=0)

## 5. GitHub Actions

- [x] 5.1 Create `.github/workflows/release.yml` triggered on `v*` tags

## 6. Nix Integration

- [x] 6.1 Update `flake.nix` to read version from `src/VERSION`

## 7. Verification

- [x] 7.1 Run `make build` and verify `drs --version` outputs `0.2.0`
- [x] 7.2 Run `make test` and verify all tests pass
