## ADDED Requirements

### Requirement: Enter key opens terminal
Pressing `enter` SHALL open a terminal in the selected repo's directory (replacing the old `e` binding).

#### Scenario: Press enter to open
- **WHEN** the user presses `enter` with a repo selected
- **THEN** a terminal SHALL open in that repo's directory

#### Scenario: Old e key no longer opens
- **WHEN** the user presses `e`
- **THEN** nothing SHALL happen

### Requirement: Vim-style jk navigation
Pressing `j` SHALL move the cursor down and `k` SHALL move the cursor up, in addition to arrow keys.

#### Scenario: j moves down
- **WHEN** the user presses `j` in the repo list view
- **THEN** the cursor SHALL move down one item

#### Scenario: k moves up
- **WHEN** the user presses `k` in the repo list view
- **THEN** the cursor SHALL move up one item

### Requirement: Nav bar reflects updated keybindings
The nav bar SHALL show the updated keybindings including `enter:open` and `jk/↑↓:navigate`.

#### Scenario: Nav bar content
- **WHEN** the TUI is displayed
- **THEN** the nav bar SHALL show `enter:open` instead of `e:edit` and `jk/↑↓:navigate` instead of `↑↓:navigate`

### Requirement: README updated
The README keybinding table SHALL reflect the new keybindings.

#### Scenario: README accuracy
- **WHEN** the README is read
- **THEN** the keybinding table SHALL list `enter` for opening and `j`/`k` for navigation
