## ADDED Requirements

### Requirement: Log panel hidden by default
The log panel SHALL be hidden when the application starts. The status panel SHALL use the space that the log panel would have occupied.

#### Scenario: Application startup
- **WHEN** the application starts
- **THEN** the log panel is not visible and the status panel fills the remaining vertical space below the repo panel

### Requirement: Toggle log visibility with l key
The user SHALL be able to press `l` to toggle the log panel visibility on and off.

#### Scenario: Show log panel
- **WHEN** the log panel is hidden and the user presses `l`
- **THEN** the log panel becomes visible with its standard height and the status panel shrinks to accommodate it

#### Scenario: Hide log panel
- **WHEN** the log panel is visible and the user presses `l`
- **THEN** the log panel is hidden and the status panel expands to fill the freed space

### Requirement: Log continues capturing when hidden
The log writer SHALL continue capturing log output even when the log panel is not visible.

#### Scenario: Logs captured while hidden
- **WHEN** the log panel is hidden and log messages are written
- **AND** the user presses `l` to show the log panel
- **THEN** all log messages written while hidden are visible in the log panel

### Requirement: Tab cycling skips hidden log
When the log panel is hidden, Tab SHALL cycle only between repo and status views.

#### Scenario: Tab with hidden log
- **WHEN** the log panel is hidden and the user presses Tab
- **THEN** focus cycles between repo and status views only, skipping the log view

#### Scenario: Active log view when hiding
- **WHEN** the log panel is the active view and the user presses `l` to hide it
- **THEN** focus moves to the repo view
