## 1. Panel Height Tests

- [x] 1.1 Create `src/ui/ui_test.go` with table-driven tests for `repoPanelHeight` (few repos, many repos, zero repos)
- [x] 1.2 Add tests for `logPanelHeight` (various terminal heights, bounded between 1 and 10)
- [x] 1.3 Add tests for `statusPanelHeight` (fills remaining space, returns positive value)

## 2. Repo List Rendering Tests

- [x] 2.1 Add tests for `renderRepoList` (all repos visible, cursor position reflected, empty list message)

## 3. Verification

- [x] 3.1 Run `make test` and verify all tests pass
