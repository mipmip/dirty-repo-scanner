## Why

The nav bar shows the app name (`dirty-repo-scanner`) on the right but not the version. Displaying the version helps users confirm which build they're running.

## What Changes

- Pass the version string from `main.go` to `ui.Run()`
- Store it in the model
- Append version to the app name in the nav bar: `dirty-repo-scanner 0.2.1`

## Capabilities

### New Capabilities

- `navbar-version`: Display version number in the nav bar

### Modified Capabilities

(none)

## Impact

- **Code**: `src/main.go` — pass `version` to `ui.Run()`. `src/ui/ui.go` — update `Run()` signature, `newModel()`, `renderNavBar()`
