## Context

The `doEdit()` method in `src/ui/ui.go` runs the configured `edit_command` by suspending the bubbletea TUI via `tea.ExecProcess`. This works but takes over the full terminal. tmux's `display-popup` command can run a command in a floating overlay window.

## Goals / Non-Goals

**Goals:**
- Detect tmux via `$TMUX` environment variable
- When in tmux: run editor via `tmux display-popup -E -w 80% -h 80% -- <edit_command>`
- When not in tmux: keep current `tea.ExecProcess` behavior
- The popup should be centered and take ~80% of the terminal

**Non-Goals:**
- Configurable popup size (hardcode 80% for now)
- Support for other terminal multiplexers (screen, zellij)
- Changing the edit_command config format

## Decisions

### 1. Use `tmux display-popup -E` for the editor

The `-E` flag closes the popup automatically when the command exits. The popup runs the full edit command string.

**Why**: Clean lifecycle — popup appears, user edits, popup disappears. No stale windows.

### 2. Detect tmux via `$TMUX` env var

Check `os.Getenv("TMUX") != ""` to determine if we're inside tmux.

**Why**: This is the standard way to detect tmux. It's set by tmux for all child processes.

### 3. Use `tea.ExecProcess` for tmux popup too

Even the tmux popup command should go through `tea.ExecProcess` so the TUI properly suspends and resumes. The difference is just what command is executed: `tmux display-popup ...` wrapping the edit command.

**Why**: Bubbletea needs to know about the external process to properly handle terminal state. Using `exec.Command` directly without `tea.ExecProcess` could leave the terminal in a bad state.

### 4. Popup size: 80% width and height

Use `-w 80% -h 80%` for the popup dimensions.

**Why**: Large enough to be useful for editing, small enough to see the TUI behind it.

## Risks / Trade-offs

- **[tmux version compatibility]** `display-popup` was added in tmux 3.2 (2021). Older versions will fail. → Acceptable; tmux 3.2+ is widely available.
- **[Non-terminal editors]** If `edit_command` is something like `code` (VS Code), the popup won't help since VS Code opens externally. → The popup will just briefly appear and close. No worse than current behavior.
