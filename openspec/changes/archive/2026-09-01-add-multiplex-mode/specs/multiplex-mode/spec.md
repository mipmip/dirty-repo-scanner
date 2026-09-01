## ADDED Requirements

### Requirement: Multiplex mode flag
The application SHALL accept a `--multiplex` boolean CLI flag that selects
multiplex mode for the TUI. When the flag is absent, the TUI SHALL behave
exactly as today (management mode).

#### Scenario: Flag enables multiplex mode
- **WHEN** the user runs `drs --multiplex`
- **THEN** the TUI SHALL run in multiplex mode, where Enter runs the configured
  `switch_command` instead of the editor action

#### Scenario: No flag keeps management mode
- **WHEN** the user runs `drs` without `--multiplex`
- **THEN** the Enter and `e` keys SHALL retain their current editor behavior and
  `switch_command` SHALL NOT be executed

### Requirement: switch_command configuration
The configuration SHALL support a `switch_command` string key holding the
command template executed by Enter in multiplex mode.

#### Scenario: Configured switch command
- **WHEN** the config contains `switch_command: tmux new -As %REPO_NAME -c %WORKING_DIRECTORY`
- **AND** the TUI runs in multiplex mode
- **THEN** activating a repository SHALL execute that command with placeholders
  substituted for the selected repository

#### Scenario: Missing switch command in multiplex mode
- **WHEN** the TUI is launched with `--multiplex` and `switch_command` is empty
- **THEN** the application SHALL report a clear error and exit with a non-zero
  status without entering the interactive UI

### Requirement: Placeholder substitution
The application SHALL substitute placeholders in `switch_command` against the
selected repository before execution, using literal string replacement.

#### Scenario: Working directory placeholder
- **WHEN** `switch_command` contains `%WORKING_DIRECTORY`
- **THEN** every occurrence SHALL be replaced with the selected repository's
  absolute path

#### Scenario: Repository name placeholder
- **WHEN** `switch_command` contains `%REPO_NAME`
- **THEN** every occurrence SHALL be replaced with the base name of the selected
  repository's directory

### Requirement: Shell-free command execution
The application SHALL split the rendered `switch_command` into arguments using
POSIX shell-quoting rules and execute it directly without invoking a shell, so
repository names or paths cannot inject shell behavior.

#### Scenario: Quoted argument with spaces
- **WHEN** the rendered command is `tmux new -As 'my repo' -c '/path/with spaces'`
- **THEN** the application SHALL execute the program `tmux` with the arguments
  `new`, `-As`, `my repo`, `-c`, `/path/with spaces` as distinct argv entries

#### Scenario: No shell interpretation
- **WHEN** a substituted value contains shell metacharacters such as `;` or `$()`
- **THEN** those characters SHALL be passed as literal argument text and SHALL
  NOT be interpreted as shell operations

### Requirement: Quit after successful switch
The TUI SHALL quit with exit code 0 after a `switch_command` completes
successfully in multiplex mode, so a surrounding tmux popup closes.

#### Scenario: Successful switch quits the TUI
- **WHEN** the `switch_command` runs and returns success
- **THEN** the TUI SHALL exit with status 0

#### Scenario: Failed switch keeps the TUI open
- **WHEN** the `switch_command` fails to run or returns a non-zero status
- **THEN** the TUI SHALL remain open and surface the error rather than quitting
