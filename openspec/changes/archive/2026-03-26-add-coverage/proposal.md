## Why

The project now has unit tests (from the add-testing-framework change) but no way to measure or report test coverage. Adding coverage support lets developers see which code paths are tested and identify gaps.

## What Changes

- Add a `cover` target to the Makefile that runs tests with coverage and generates an HTML report
- Add coverage output files to `.gitignore`

## Capabilities

### New Capabilities

- `test-coverage`: Makefile target for generating and viewing Go test coverage reports

### Modified Capabilities

(none)

## Impact

- **Build**: New `cover` target in Makefile
- **Files**: `.gitignore` updated with coverage artifacts (`coverage.out`, `coverage.html`)
- **Dependencies**: None — uses built-in `go test -coverprofile`
