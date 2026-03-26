## 1. Go Module and Imports

- [x] 1.1 Change module path in `go.mod` from `github.com/mipmip/dirtygit` to `github.com/mipmip/dirty-repo-scanner`
- [x] 1.2 Update imports in `main.go` (`dirtygit/scanner` → `dirty-repo-scanner/scanner`, `dirtygit/ui` → `dirty-repo-scanner/ui`)
- [x] 1.3 Update import in `ui/ui.go` (`dirtygit/scanner` → `dirty-repo-scanner/scanner`)

## 2. Config File Rename

- [x] 2.1 Rename `.dirtygit.yml` to `.dirty-repo-scanner.yml` in the project root
- [x] 2.2 Update `//go:embed` directive in `main.go` to reference `.dirty-repo-scanner.yml`
- [x] 2.3 Update default config path in `main.go` from `.dirtygit.yml` to `.dirty-repo-scanner.yml`

## 3. Binary and App Name

- [x] 3.1 Change `app.Name` in `main.go` from `"dirtygit"` to `"dirty-repo-scanner"`
- [x] 3.2 Update `Makefile` output binary name from `dirtygit` to `drs`

## 4. Nix Packaging

- [x] 4.1 Update `pname` in `package.nix` from `"dirtygit"` to `"dirty-repo-scanner"`
- [x] 4.2 Update `homepage` URL in `package.nix` to `github.com/mipmip/dirty-repo-scanner`
- [x] 4.3 Update package name references in `flake.nix` from `dirtygit` to `dirty-repo-scanner`

## 5. Build Artifacts and Docs

- [x] 5.1 Update `.gitignore` entries from `dirtygit` to `drs`
- [x] 5.2 Update `README.md` — project name, install command, binary name, config file references

## 6. Verification

- [x] 6.1 Run `go build -o drs ./...` to verify compilation
- [x] 6.2 Verify binary identifies as `dirty-repo-scanner` in `--help` output

## 7. GitHub Repo Rename (Manual)

- [ ] 7.1 Rename repository on GitHub: Settings → General → "dirty-repo-scanner"
