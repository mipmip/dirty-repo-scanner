## 1. Thread Version to UI

- [x] 1.1 Add `version string` field to model struct
- [x] 1.2 Update `newModel()` to accept and store version
- [x] 1.3 Update `Run()` signature to accept version and pass to `newModel()`
- [x] 1.4 Update `ui.Run()` call in `src/main.go` to pass `version`

## 2. Display in Nav Bar

- [x] 2.1 Update `renderNavBar()` right side from `"dirty-repo-scanner"` to `"dirty-repo-scanner " + m.version`

## 3. Verification

- [x] 3.1 Run `make build` and `make test`
