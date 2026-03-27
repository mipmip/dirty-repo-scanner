## ADDED Requirements

### Requirement: Detect tmux environment
The application SHALL detect whether it is running inside tmux by checking the `$TMUX` environment variable.

#### Scenario: Inside tmux
- **WHEN** the `$TMUX` environment variable is set and non-empty
- **THEN** the application SHALL consider itself running inside tmux

#### Scenario: Outside tmux
- **WHEN** the `$TMUX` environment variable is empty or unset
- **THEN** the application SHALL consider itself running outside tmux

### Requirement: Editor in tmux popup
When running inside tmux, pressing `e` SHALL open the editor in a tmux popup window.

#### Scenario: Edit in tmux popup
- **WHEN** the user presses `e` while inside tmux
- **THEN** the application SHALL execute `tmux display-popup -E -w 80% -h 80% -- <edit_command>` with the selected repo's working directory substituted

#### Scenario: Popup closes on exit
- **WHEN** the editor command in the popup exits
- **THEN** the popup SHALL close automatically and the TUI SHALL resume

### Requirement: Fallback for non-tmux
When not running inside tmux, pressing `e` SHALL continue to use the current inline editor behavior.

#### Scenario: Edit outside tmux
- **WHEN** the user presses `e` while not inside tmux
- **THEN** the application SHALL suspend the TUI and run the editor inline (current behavior)
