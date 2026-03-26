# bubbletea-tui Specification

## Purpose
TBD - created by archiving change migrate-to-bubbletea. Update Purpose after archive.
## Requirements
### Requirement: Three-panel layout
The TUI SHALL render three vertically stacked panels: a repository list panel at the top, a status panel in the middle, and a log panel at the bottom. The repository list panel height SHALL be `min(numRepos+2, terminalHeight/2)`. The log panel SHALL be fixed at 10 lines. The status panel SHALL fill the remaining vertical space. The terminal SHALL require a minimum height of 20 lines.

#### Scenario: Normal layout rendering
- **WHEN** the terminal has at least 20 lines of height and repositories are loaded
- **THEN** three panels are visible: repo list (top), status (middle), log (bottom)

#### Scenario: Repo list height adapts to content
- **WHEN** there are 5 repositories and the terminal is 40 lines tall
- **THEN** the repo list panel is 7 lines tall (5+2)

#### Scenario: Repo list height capped at half screen
- **WHEN** there are 30 repositories and the terminal is 40 lines tall
- **THEN** the repo list panel is 20 lines tall (40/2)

#### Scenario: Terminal too small
- **WHEN** the terminal height is less than 20 lines
- **THEN** the program displays an error or exits gracefully

### Requirement: Repository list navigation
The user SHALL navigate the repository list using Arrow Up and Arrow Down keys. The currently selected repository SHALL be visually highlighted with a green background and black foreground.

#### Scenario: Arrow down moves selection
- **WHEN** the user presses Arrow Down in the repo list
- **THEN** the selection moves to the next repository

#### Scenario: Arrow up moves selection
- **WHEN** the user presses Arrow Up in the repo list
- **THEN** the selection moves to the previous repository

#### Scenario: Selection highlighting
- **WHEN** a repository is selected
- **THEN** it is displayed with green background and black text

### Requirement: Status panel shows diff for selected repo
When a repository is selected in the repo list, the status panel SHALL display the git file status in `SW` format (staging code, worktree code, filename), sorted alphabetically by path.

#### Scenario: Repository with dirty files selected
- **WHEN** the user selects a repository that has modified files
- **THEN** the status panel shows a header line " SW" followed by "-----" and one line per file with staging code, worktree code, and path

#### Scenario: Empty repository selected
- **WHEN** the user selects a repository with no dirty files
- **THEN** the status panel is empty

### Requirement: View cycling with Tab
Pressing Tab SHALL cycle the active view in order: repo → status → log → repo. The active view SHALL be visually distinguishable.

#### Scenario: Tab from repo view
- **WHEN** the user presses Tab while the repo view is active
- **THEN** the status view becomes active

#### Scenario: Tab from log view
- **WHEN** the user presses Tab while the log view is active
- **THEN** the repo view becomes active

### Requirement: Scan trigger and progress
Pressing 's' SHALL trigger a repository scan. During scanning, a modal overlay with a spinner SHALL be displayed. After scanning completes, the repository list SHALL update with results showing only dirty repositories.

#### Scenario: Initial scan on startup
- **WHEN** the program starts
- **THEN** a scan is triggered automatically

#### Scenario: Manual scan trigger
- **WHEN** the user presses 's'
- **THEN** a new scan begins and a scanning modal with spinner is displayed

#### Scenario: Scan completes
- **WHEN** the scan finishes
- **THEN** the scanning modal disappears and the repo list shows dirty repositories sorted alphabetically

### Requirement: Error display
When an error occurs during scanning, a centered error modal SHALL be displayed with the error message. The modal SHALL have text wrapping.

#### Scenario: Scan error
- **WHEN** a scan produces an error
- **THEN** a centered error modal appears showing the error text

### Requirement: Log output
The log panel SHALL capture output from Go's `log` package. The log panel SHALL autoscroll to show the latest entries and SHALL wrap long lines.

#### Scenario: Scanner logs progress
- **WHEN** the scanner logs repository scan durations
- **THEN** the log panel displays them and scrolls to the latest entry

### Requirement: Edit command
Pressing 'e' SHALL launch the configured edit command with the selected repository path substituted for `%WORKING_DIRECTORY`. The TUI SHALL properly suspend and restore around the external process.

#### Scenario: Edit selected repository
- **WHEN** the user presses 'e' with a repository selected
- **THEN** the configured editor opens with that repository's path, and the TUI resumes after the editor exits

#### Scenario: Edit with no selection
- **WHEN** the user presses 'e' with no repository selected
- **THEN** nothing happens

### Requirement: Quit
Pressing 'q' or Ctrl+C SHALL exit the program cleanly.

#### Scenario: Quit with q
- **WHEN** the user presses 'q'
- **THEN** the program exits

#### Scenario: Quit with Ctrl+C
- **WHEN** the user presses Ctrl+C
- **THEN** the program exits

### Requirement: Program initialization
The `ui.Run` function SHALL accept `*scanner.Config` and `bool` (ignore_dir_errors) parameters and return an `error`. It SHALL create and run a Bubbletea program with alternate screen mode.

#### Scenario: Successful startup
- **WHEN** `ui.Run(config, true)` is called
- **THEN** a Bubbletea program starts in alternate screen mode and returns nil on clean exit

