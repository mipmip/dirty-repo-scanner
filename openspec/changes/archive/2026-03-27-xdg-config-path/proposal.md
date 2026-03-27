## Why

The default config file currently lives at `~/.dirty-repo-scanner.yml`, a dotfile in the home directory root. The XDG Base Directory Specification recommends user configuration go under `$XDG_CONFIG_HOME` (defaulting to `~/.config/`). Moving the config there follows modern conventions and reduces home directory clutter.

## What Changes

- **BREAKING**: Default config path changes from `~/.dirty-repo-scanner.yml` to `~/.config/dirty-repo-scanner/config.yml`
- Respect `$XDG_CONFIG_HOME` if set (e.g., `$XDG_CONFIG_HOME/dirty-repo-scanner/config.yml`)
- Update embedded default config filename from `.dirty-repo-scanner.yml` to `config.yml`
- Rename the embedded config file in `src/` from `.dirty-repo-scanner.yml` to `config.yml`
- Update README and documentation references

## Capabilities

### New Capabilities

- `xdg-config`: XDG-compliant config file path resolution

### Modified Capabilities

(none)

## Impact

- **Code**: `src/main.go` — update `getDefaultConfigPath()` and `//go:embed` directive
- **Files**: Rename `src/.dirty-repo-scanner.yml` to `src/config.yml`
- **Users**: Must move their config file from `~/.dirty-repo-scanner.yml` to `~/.config/dirty-repo-scanner/config.yml`
- **Docs**: README config section updated
