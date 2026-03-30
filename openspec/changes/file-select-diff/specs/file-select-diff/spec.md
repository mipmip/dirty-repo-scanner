## ADDED Requirements

### Requirement: Files are selectable in the status panel
When the status panel is active, files SHALL be individually selectable with a visible cursor.

#### Scenario: Navigate files with j/k
- **WHEN** the status panel is active and the user presses `j` or `down`
- **THEN** the file cursor SHALL move to the next file

#### Scenario: Navigate files with k/up
- **WHEN** the status panel is active and the user presses `k` or `up`
- **THEN** the file cursor SHALL move to the previous file

#### Scenario: File cursor highlight
- **WHEN** a file is selected in the status panel
- **THEN** it SHALL be visually highlighted (same style as repo cursor)

### Requirement: Git diff shown for selected file
When a file is selected in the status panel, the panel SHALL split vertically to show the git diff.

#### Scenario: Diff displayed on file selection
- **WHEN** a file is selected and the status panel is active
- **THEN** the status panel SHALL show the file list on the left and `git diff` output on the right

#### Scenario: Untracked file
- **WHEN** the selected file is untracked (not in git index)
- **THEN** the diff area SHALL show "Untracked file"

#### Scenario: Diff scrollable
- **WHEN** the diff output is taller than the viewport
- **THEN** the user SHALL be able to scroll the diff (when navigating within the status view)

### Requirement: Edit file with e keybinding
Pressing `e` when a file is selected SHALL open the file in an editor.

#### Scenario: Edit in tmux
- **WHEN** the user presses `e` with a file selected and the app is running in tmux
- **THEN** a tmux popup SHALL open with `$EDITOR` editing the full file path

#### Scenario: Edit outside tmux with xdg-open
- **WHEN** the user presses `e` outside tmux and `xdg-open` is available
- **THEN** the file SHALL be opened with `xdg-open`

#### Scenario: Edit outside tmux fallback to EDITOR
- **WHEN** the user presses `e` outside tmux and `xdg-open` is not available
- **THEN** the file SHALL be opened with `$EDITOR`

### Requirement: Nav bar shows edit keybinding
The nav bar SHALL include the `e:edit` keybinding.

#### Scenario: Nav bar content
- **WHEN** the TUI is displayed
- **THEN** the nav bar SHALL include `e:edit`
