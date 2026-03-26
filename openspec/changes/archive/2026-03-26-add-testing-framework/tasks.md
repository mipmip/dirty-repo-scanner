## 1. Excluder Tests

- [x] 1.1 Create `scanner/excluded_test.go` with table-driven tests for `IsExcluded` (file glob match, no match, dir glob match, no match, no patterns)
- [x] 1.2 Add tests for `FilterGitStatus` (excluded files removed, no files excluded, all files excluded)

## 2. Config Parsing Tests

- [x] 2.1 Create `scanner/scan_test.go` with tests for `ParseConfigFile` (valid file, missing file with default, invalid YAML)

## 3. Skip Helper Tests

- [x] 3.1 Create `scanner/find_test.go` with tests for `skip` (full path match, basename match, no match)

## 4. Build Integration

- [x] 4.1 Add `test` target to Makefile that runs `go test ./...`

## 5. Verification

- [x] 5.1 Run `make test` and verify all tests pass
