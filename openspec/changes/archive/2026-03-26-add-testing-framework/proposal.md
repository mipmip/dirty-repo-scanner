## Why

The project has zero tests. The scanner package contains pure logic (config parsing, path exclusion, git status filtering) that is well-suited for unit testing but has no test coverage. Adding a testing foundation now makes future changes safer and catches regressions early.

## What Changes

- Add Go test files for the `scanner` package covering config parsing, exclusion logic, and git status filtering
- Add a `test` target to the Makefile
- Add test dependencies if needed (standard library `testing` should suffice)

## Capabilities

### New Capabilities

- `scanner-tests`: Unit tests for the scanner package — config parsing (`ParseConfigFile`), path exclusion (`Excluder.IsExcluded`), git status filtering (`Excluder.FilterGitStatus`), and the `skip` helper

### Modified Capabilities

(none)

## Impact

- **Code**: New `*_test.go` files in `scanner/`
- **Build**: New `test` target in Makefile
- **Dependencies**: No new dependencies expected (Go standard `testing` package)
- **CI**: Tests can be wired into CI later
