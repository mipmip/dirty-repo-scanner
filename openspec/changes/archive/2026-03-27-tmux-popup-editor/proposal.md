## Why

When running inside tmux, pressing `e` to edit a repo currently suspends the entire TUI to run the editor process. tmux supports popup windows (`tmux display-popup`) which can overlay the editor on top of the TUI without suspending it. This provides a smoother editing experience — the TUI stays visible and the editor appears in a floating window.

## What Changes

- Detect whether the app is running inside tmux (check `$TMUX` environment variable)
- When inside tmux, open the editor via `tmux display-popup` instead of suspending the TUI
- When not inside tmux, keep the current behavior (suspend TUI, run editor inline)
- The tmux popup runs the configured `edit_command` in a centered popup window

## Capabilities

### New Capabilities

- `tmux-popup-editor`: Detect tmux and open editor in a tmux popup window instead of suspending the TUI

### Modified Capabilities

(none)

## Impact

- **Code**: `src/ui/ui.go` — modify `doEdit()` to branch on tmux detection
- **Dependencies**: None (tmux is detected via env var, invoked via `exec.Command`)
- **Behavior**: Only changes editor behavior when running inside tmux; outside tmux, behavior is unchanged
