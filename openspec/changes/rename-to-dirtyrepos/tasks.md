## 1. Go Module and Imports

- [ ] 1.1 Change module path in `go.mod` from `github.com/mipmip/dirtygit` to `github.com/mipmip/dirty-repo-scanner`
- [ ] 1.2 Update imports in `main.go` (`dirtygit/scanner` → `dirty-repo-scanner/scanner`, `dirtygit/ui` → `dirty-repo-scanner/ui`)
- [ ] 1.3 Update import in `ui/ui.go` (`dirtygit/scanner` → `dirty-repo-scanner/scanner`)

## 2. Config File Rename

- [ ] 2.1 Rename `.dirtygit.yml` to `.dirty-repo-scanner.yml` in the project root
- [ ] 2.2 Update `//go:embed` directive in `main.go` to reference `.dirty-repo-scanner.yml`
- [ ] 2.3 Update default config path in `main.go` from `.dirtygit.yml` to `.dirty-repo-scanner.yml`

## 3. Binary and App Name

- [ ] 3.1 Change `app.Name` in `main.go` from `"dirtygit"` to `"dirty-repo-scanner"`
- [ ] 3.2 Update `Makefile` output binary name from `dirtygit` to `drs`

## 4. Nix Packaging

- [ ] 4.1 Update `pname` in `package.nix` from `"dirtygit"` to `"dirty-repo-scanner"`
- [ ] 4.2 Update `homepage` URL in `package.nix` to `github.com/mipmip/dirty-repo-scanner`
- [ ] 4.3 Update package name references in `flake.nix` from `dirtygit` to `dirty-repo-scanner`

## 5. Build Artifacts and Docs

- [ ] 5.1 Update `.gitignore` entries from `dirtygit` to `drs`
- [ ] 5.2 Update `README.md` — project name, install command, binary name, config file references

## 6. Verification

- [ ] 6.1 Run `go build -o drs ./...` to verify compilation
- [ ] 6.2 Verify binary identifies as `dirty-repo-scanner` in `--help` output

## 7. GitHub Repo Rename (Manual)

- [ ] 7.1 Rename repository on GitHub: Settings → General → "dirty-repo-scanner"
