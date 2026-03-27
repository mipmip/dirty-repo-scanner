## ADDED Requirements

### Requirement: XDG-compliant default config path
The application SHALL use an XDG-compliant path as the default config file location.

#### Scenario: Default on Linux
- **WHEN** `$XDG_CONFIG_HOME` is not set and the app runs on Linux
- **THEN** the default config path SHALL be `~/.config/dirty-repo-scanner/config.yml`

#### Scenario: Custom XDG_CONFIG_HOME
- **WHEN** `$XDG_CONFIG_HOME` is set to `/custom/config`
- **THEN** the default config path SHALL be `/custom/config/dirty-repo-scanner/config.yml`

#### Scenario: macOS default
- **WHEN** the app runs on macOS
- **THEN** the default config path SHALL be `~/Library/Application Support/dirty-repo-scanner/config.yml`

### Requirement: Config flag override still works
The `--config` / `-c` flag SHALL continue to override the default config path.

#### Scenario: Custom config path
- **WHEN** `drs --config /tmp/myconfig.yml` is run
- **THEN** the application SHALL use `/tmp/myconfig.yml` as the config file

### Requirement: Embedded default config named config.yml
The embedded default config file SHALL be named `config.yml`.

#### Scenario: Embedded config
- **WHEN** no config file exists at the default path
- **THEN** the application SHALL fall back to the embedded `config.yml` default
