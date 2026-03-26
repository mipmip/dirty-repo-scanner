## Context

The `ui` package (`src/ui/ui.go`) contains a bubbletea TUI. Most functions depend on terminal state, but several are pure logic operating on the `model` struct fields (`height`, `repoPaths`, `cursor`). These can be tested by constructing a model directly without running a bubbletea program.

Key testable functions:
- `repoPanelHeight()` — returns int based on `m.height` and `len(m.repoPaths)`
- `statusPanelHeight()` — returns int based on repo and log panel heights
- `logPanelHeight()` — returns int based on `m.height`
- `renderRepoList(width)` — returns string containing repo paths with cursor highlighting

## Goals / Non-Goals

**Goals:**
- Unit tests for panel height calculations across various terminal sizes and repo counts
- Tests for `renderRepoList` verifying correct repos appear in output and cursor position is reflected

**Non-Goals:**
- Testing `View()` output (contains ANSI escape codes from lipgloss, brittle to assert)
- Testing `Update()` message handling (requires bubbletea test infrastructure)
- Testing `doEdit` / `doScan` (side-effectful)

## Decisions

### 1. Construct model structs directly in tests

Create `model` values with specific `height`, `width`, `repoPaths`, and `cursor` values. Call methods directly.

**Why**: The model is a plain struct with exported-within-package fields. No constructor ceremony needed. This avoids needing a running bubbletea program.

### 2. Assert on string content for renderRepoList, not on exact styling

Check that `renderRepoList` output contains the expected repo paths using `strings.Contains`. Don't assert on ANSI codes or exact formatting.

**Why**: Lipgloss output includes terminal-dependent escape sequences. Content assertions are stable; style assertions are brittle.

### 3. Table-driven tests for height calculations

Test each height function with a matrix of terminal heights and repo counts.

**Why**: These functions have interesting edge cases (very small terminals, zero repos, more repos than screen space). Table-driven tests cover these efficiently.

## Risks / Trade-offs

- **[Internal API coupling]** Tests access unexported struct fields and methods, so they live in `package ui` (not `package ui_test`). → This is fine; they test internal logic, not public API.
