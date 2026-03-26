## ADDED Requirements

### Requirement: Nav bar position
The navigation bar SHALL be rendered as the last line at the bottom of the terminal, below all panels.

#### Scenario: Nav bar visible
- **WHEN** the TUI is running at any terminal size >= 20 lines
- **THEN** a single-line navigation bar SHALL be visible at the bottom of the screen

### Requirement: Keybinding hints on the left
The navigation bar SHALL display available keybindings with their actions on the left side.

#### Scenario: Keybindings displayed
- **WHEN** the TUI is running
- **THEN** the left side of the nav bar SHALL show key-action pairs for: quit, scan, edit, switch panels, and navigate

### Requirement: App name on the right
The navigation bar SHALL display the application name on the right side.

#### Scenario: App name displayed
- **WHEN** the TUI is running
- **THEN** the right side of the nav bar SHALL show `dirty-repo-scanner`

### Requirement: Layout adjustment
The panel layout SHALL account for the nav bar so that no content is hidden behind it.

#### Scenario: Panels fit with nav bar
- **WHEN** the TUI is rendered
- **THEN** the total height of all panels plus the nav bar SHALL NOT exceed the terminal height
