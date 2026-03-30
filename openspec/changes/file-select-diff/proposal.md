## Why

The status panel currently shows a read-only list of dirty files with their git status codes. There's no way to inspect individual file changes or open a specific file for editing. Users need to switch to a terminal to run `git diff` or open files manually.

## What Changes

- Make files selectable in the status panel when it has focus (cursor navigation with `j`/`k`/arrows)
- Add a `fileCursor` tracking which file is selected in the status panel
- When a file is selected, split the status panel vertically: file list on the left, `git diff` output for the selected file on the right
- Add `e` keybinding (shown in nav bar) that opens the selected file:
  - In tmux: open `$EDITOR <file>` in a tmux popup
  - Outside tmux with `xdg-open` available: open with `xdg-open`
  - Otherwise: open with `$EDITOR`
- Update nav bar to show `e:edit` when status panel is active

## Capabilities

### New Capabilities

- `file-select-diff`: Selectable files in status panel with inline git diff preview and file editing

### Modified Capabilities

(none)

## Impact

- **Code**: `src/ui/ui.go` — new `fileCursor` state, split status panel rendering, diff viewport, file edit action
- **Dependencies**: None (uses `git diff` via `exec.Command`)
- **Behavior**: Status panel gains interactive file selection; diff shown alongside file list
