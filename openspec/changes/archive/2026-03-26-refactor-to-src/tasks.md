## 1. Move Files

- [x] 1.1 Create `src/` directory
- [x] 1.2 Move `main.go` to `src/main.go`
- [x] 1.3 Move `scanner/` to `src/scanner/`
- [x] 1.4 Move `ui/` to `src/ui/`
- [x] 1.5 Move `.dirty-repo-scanner.yml` to `src/.dirty-repo-scanner.yml`

## 2. Update Import Paths

- [x] 2.1 Update imports in `src/main.go` to `dirty-repo-scanner/src/scanner` and `dirty-repo-scanner/src/ui`
- [x] 2.2 Update import in `src/ui/ui.go` to `dirty-repo-scanner/src/scanner`

## 3. Update Build Config

- [x] 3.1 Update Makefile `build` target to `go build -o drs ./src`
- [x] 3.2 Update Makefile `test` and `cover` targets if needed (they use `./...` so may work as-is)

## 4. Verification

- [x] 4.1 Run `make build` and verify the binary works
- [x] 4.2 Run `make test` and verify all tests pass
