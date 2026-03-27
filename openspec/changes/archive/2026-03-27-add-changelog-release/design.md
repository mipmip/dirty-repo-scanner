## Context

The project is a Go CLI tool using urfave/cli, built with a Makefile and Nix flake. It uses both git and jj (colocated). The teejay project provides a proven pattern for versioning and releasing that should be adapted here.

Key differences from teejay: this project uses urfave/cli (has built-in `app.Version`), the binary is `drs`, and main.go is at `src/main.go`.

## Goals / Non-Goals

**Goals:**
- Embedded version from a VERSION file (`//go:embed`)
- `drs --version` displays the version
- CHANGELOG.md with Keep a Changelog format
- Interactive release script supporting git and jj-colocated repos
- GoReleaser configs for Linux and macOS cross-compilation
- GitHub Actions release workflow on `v*` tags
- Nix flake reads version from VERSION file

**Non-Goals:**
- Windows builds
- Automatic changelog generation
- Publishing to package registries beyond GitHub Releases

## Decisions

### 1. VERSION file at `src/VERSION`

Place the VERSION file next to `src/main.go` where `//go:embed` can reach it. Contains just the version string (e.g., `0.2.0`).

**Why**: Same pattern as teejay. The `//go:embed` directive resolves relative to the source file, so VERSION must be in `src/`.

### 2. Use urfave/cli's built-in `app.Version`

Set `app.Version = version` instead of manually handling `--version`. Urfave/cli provides `--version` and `-v` flags automatically.

**Why**: Less code than manual flag handling. Already using the framework.

### 3. No CGO for this project

Unlike teejay (which needs CGO for audio), drs is pure Go. GoReleaser configs use `CGO_ENABLED=0` for easy cross-compilation.

**Why**: No C dependencies. Static binaries are simpler to distribute.

### 4. Release script uses git commands (works with jj-colocated)

The release script uses `git diff`, `git add`, `git commit`, `git tag`, `git push` — which all work in jj-colocated repos. No jj-specific commands needed.

**Why**: jj-colocated repos have a real `.git` directory, so git commands work directly. This is the same approach teejay uses.

### 5. Nix vendorHash update in release script

The release script includes a `update_nix_vendor_hash` function that temporarily blanks the hash, runs `nix build` to get the correct hash from the error output, and updates flake.nix.

**Why**: Go module changes between versions may change the vendor hash. Automating this prevents broken Nix builds after release.

## Risks / Trade-offs

- **[gum dependency]** Release script requires `gum` for interactive prompts. → Document as a prerequisite. It's a charmbracelet tool, fitting the project's existing charmbracelet dependency.
- **[Nix hash update may fail]** If nix isn't available, the script warns and continues. → Acceptable; manual update is straightforward.
