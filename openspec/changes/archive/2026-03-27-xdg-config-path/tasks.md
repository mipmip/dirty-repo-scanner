## 1. Config File Rename

- [x] 1.1 Rename `src/.dirty-repo-scanner.yml` to `src/config.yml`
- [x] 1.2 Update `//go:embed` directive in `src/main.go` to reference `config.yml`

## 2. Config Path Resolution

- [x] 2.1 Replace `getDefaultConfigPath()` to use `os.UserConfigDir()` with `dirty-repo-scanner/config.yml`
- [x] 2.2 Remove `go-homedir` import from `src/main.go`
- [x] 2.3 Run `go mod tidy` to clean up unused dependency

## 3. Documentation

- [x] 3.1 Update README config section with new path
- [x] 3.2 Add migration note to CHANGELOG.md under [Unreleased]

## 4. Verification

- [x] 4.1 Run `make build` and verify the binary compiles
- [x] 4.2 Run `make test` and verify all tests pass
- [x] 4.3 Verify `drs --help` shows the new default config path
