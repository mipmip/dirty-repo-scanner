## Why

The `ui` package has no test coverage. Several functions contain pure logic (panel height calculations, repo list rendering) that can be unit tested without needing a running terminal or bubbletea program.

## What Changes

- Add `src/ui/ui_test.go` with unit tests for the testable pure functions in the UI package
- Tests cover panel height calculations and repo list content

## Capabilities

### New Capabilities

- `ui-tests`: Unit tests for pure logic in the `ui` package — panel height math and repo list rendering

### Modified Capabilities

(none)

## Impact

- **Code**: New `src/ui/ui_test.go`
- **Dependencies**: None (standard `testing` package)
