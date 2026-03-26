## Context

The project has unit tests in `scanner/*_test.go` and a `make test` target. There is no coverage reporting.

## Goals / Non-Goals

**Goals:**
- `make cover` generates a coverage profile and opens an HTML report

**Non-Goals:**
- Coverage thresholds or CI enforcement
- Coverage for the `ui` package (no tests there yet)

## Decisions

### 1. Use `go tool cover -html` for the report

Generate `coverage.out` with `go test -coverprofile`, then convert to HTML with `go tool cover -html`.

**Why**: Built-in tooling, no dependencies. The HTML report is interactive and shows line-by-line coverage.

### 2. Don't auto-open the browser

Just generate the HTML file. The user can open it manually.

**Why**: Avoids platform-specific `open`/`xdg-open` commands. Simpler Makefile target.

## Risks / Trade-offs

- **[Low coverage numbers]** Only the `scanner` package has tests, so overall coverage will be low. → Expected; coverage will grow as more tests are added.
