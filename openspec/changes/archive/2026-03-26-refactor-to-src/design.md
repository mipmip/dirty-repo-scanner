## Context

Current layout has `main.go`, `scanner/`, and `ui/` at the project root alongside `go.mod`, `flake.nix`, `Makefile`, `README.md`, etc. The Go module path is `github.com/mipmip/dirty-repo-scanner`. The `main.go` uses `//go:embed .dirty-repo-scanner.yml` to embed the default config.

## Goals / Non-Goals

**Goals:**
- All Go source files live under `src/`
- Project root contains only config, build, and documentation files
- Build and tests continue to work

**Non-Goals:**
- Changing the Go module path (stays `github.com/mipmip/dirty-repo-scanner`)
- Restructuring within packages (scanner/ui internal layout unchanged)
- Changing any behavior

## Decisions

### 1. Use `src/` as a Go package parent, not a package itself

`src/` will not be a Go package — it's just a directory containing `main.go` and the sub-packages. The `main` package lives at `src/` (i.e., `package main` in `src/main.go`). Import paths become `dirty-repo-scanner/src/scanner` and `dirty-repo-scanner/src/ui`.

**Why**: Keeps the convention simple. Alternative was using `cmd/drs/main.go` (standard Go layout) but the issue specifically asks for `src/`.

### 2. Move `.dirty-repo-scanner.yml` into `src/`

The `//go:embed` directive resolves paths relative to the source file. Since `main.go` moves to `src/main.go`, the embedded config must also move to `src/.dirty-repo-scanner.yml`.

**Why**: `go:embed` cannot reference files outside its own package directory or parent module root using `..`.

### 3. Update Makefile to build from `src/`

Change build target from `go build -o drs main.go` to `go build -o drs ./src`.

**Why**: `main.go` is no longer at the root.

## Risks / Trade-offs

- **[go:embed path]** The embedded default config file must live alongside or below the source file. → Move it to `src/`.
- **[Nix vendorHash]** Moving files may invalidate the vendor hash in `package.nix`. → Will need to rebuild and update the hash.
