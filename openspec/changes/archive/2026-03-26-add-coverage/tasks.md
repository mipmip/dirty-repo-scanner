## 1. Makefile

- [x] 1.1 Add `cover` target to Makefile that runs `go test -coverprofile=coverage.out ./...` then `go tool cover -html=coverage.out -o coverage.html`

## 2. Git Ignore

- [x] 2.1 Add `coverage.out` and `coverage.html` to `.gitignore`

## 3. Verification

- [x] 3.1 Run `make cover` and verify `coverage.out` and `coverage.html` are generated
- [x] 3.2 Verify coverage files do not appear in `git status`
