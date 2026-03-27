## Context

Currently `getDefaultConfigPath()` in `src/main.go` uses `go-homedir` to resolve `~/.dirty-repo-scanner.yml`. The embedded default config is at `src/.dirty-repo-scanner.yml`.

## Goals / Non-Goals

**Goals:**
- Default config path: `$XDG_CONFIG_HOME/dirty-repo-scanner/config.yml` (or `~/.config/dirty-repo-scanner/config.yml` if `$XDG_CONFIG_HOME` is unset)
- The `--config` flag continues to work for custom paths
- Embedded default config renamed to `config.yml`

**Non-Goals:**
- Backward compatibility fallback to old path (clean break, per project convention)
- Auto-migration of old config file
- XDG for cache or data directories

## Decisions

### 1. Use `os.UserConfigDir()` from the standard library

Go's `os.UserConfigDir()` returns `$XDG_CONFIG_HOME` on Linux (defaulting to `~/.config`), `~/Library/Application Support` on macOS. This is the idiomatic Go approach.

**Why**: No external dependency needed. Can drop `go-homedir` for this use case. Respects platform conventions.

### 2. Config file named `config.yml` inside an app directory

Path: `<config-dir>/dirty-repo-scanner/config.yml`. The directory is the app namespace, the file is simply `config.yml`.

**Why**: Standard XDG pattern. Allows adding more config files later without polluting the namespace.

### 3. Drop `go-homedir` dependency

Replace `homedir.Dir()` with `os.UserConfigDir()`. The `go-homedir` package is no longer needed.

**Why**: Standard library covers this. One fewer dependency.

## Risks / Trade-offs

- **[Config not found after upgrade]** Users must manually move their config. → Document in README and CHANGELOG.
- **[macOS path difference]** `os.UserConfigDir()` returns `~/Library/Application Support` on macOS, not `~/.config`. → This is correct platform behavior.
