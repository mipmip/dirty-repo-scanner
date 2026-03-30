## Context

The version string is embedded in `src/main.go` via `//go:embed VERSION`. The `ui.Run()` function currently takes `config` and `ignoreDirErrors`. The nav bar's right side is hardcoded as `"dirty-repo-scanner"`.

## Goals / Non-Goals

**Goals:**
- Pass version to the UI and display it in the nav bar

**Non-Goals:**
- Changing the nav bar layout or styling

## Decisions

### 1. Add version parameter to Run() and model

Add `version string` to `Run()`, store in model, use in `renderNavBar()`.

**Why**: Simplest threading. The version is known at startup and never changes.

## Risks / Trade-offs

None — trivial change.
