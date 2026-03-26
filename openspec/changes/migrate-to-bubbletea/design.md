## Context

dirtygit is a TUI tool that scans for dirty git repositories and displays them in a 3-panel interface. The current UI is built on `jroimartin/gocui` (unmaintained since ~2018) in a single 382-line file (`ui/ui.go`). The scanner layer (`scanner/`) is cleanly separated and has no UI dependencies. The entry point (`main.go`) calls `ui.Run(config, ignore_dir_errors)`.

## Goals / Non-Goals

**Goals:**
- Replace gocui with Bubbletea + Lipgloss + Bubbles
- Preserve identical user-facing behavior: same layout, same keybindings, same workflow
- Maintain clean separation between UI and scanner layers
- Bump minimum Go version to 1.18+

**Non-Goals:**
- Redesigning the layout or adding new panels
- Adding new features (mouse support, fuzzy search, etc.)
- Changing the scanner layer in any way
- Changing CLI argument parsing or config format

## Decisions

### 1. Bubbletea Model structure — single top-level model with embedded sub-models

The top-level `model` struct holds:
- A `list` sub-model (from `bubbles/list` or a simple custom list) for the repo panel
- A `viewport` sub-model (from `bubbles/viewport`) for the status panel
- A `viewport` sub-model for the log panel
- State fields: `repositories`, `scanning`, `scanProgress`, `err`, `activeView`

**Why not separate component models with a parent dispatcher?** The UI is simple enough (3 fixed panels, no dynamic composition) that a flat model is clearer and avoids message-routing boilerplate.

### 2. Panel layout — lipgloss.JoinVertical with fixed/dynamic sizing

The layout uses `lipgloss.JoinVertical` to stack three styled blocks:
1. Repo list: height = `min(len(repos)+2, termHeight/2)`
2. Status: fills remaining space
3. Log: fixed 10 lines at bottom

`lipgloss.Place` centers modal overlays (scanning spinner, error).

**Why not a flexbox/grid library?** The layout is a simple vertical stack with one dynamic region. lipgloss's join functions handle this directly without additional dependencies.

### 3. Background scanning — tea.Cmd returning a tea.Msg

Replace the goroutine + channel + atomic flag pattern with:
- A `scanMsg` type carrying `(MultiGitStatus, error)`
- A `tea.Cmd` function that calls `scanner.Scan()` and returns `scanMsg`
- A `tickMsg` from `tea.Tick` for the spinner animation during scanning

The model's `scanning bool` field controls whether the spinner overlay renders. No atomic operations needed — Bubbletea's Update loop is single-threaded.

**Why not keep channels?** Bubbletea's Cmd/Msg pattern is the idiomatic way to handle async work. It's simpler, avoids race conditions by design, and eliminates the need for `g.Update()` hacks.

### 4. Editor launch — tea.ExecProcess

Replace `exec.Command().Run()` (which blocks the gocui main loop) with `tea.ExecProcess`, which properly suspends and restores the terminal.

### 5. Log output — custom io.Writer that sends tea.Msg

The scanner uses `log.Println` to write progress. Create a custom `io.Writer` that captures log output and sends it as a `logMsg` to the Bubbletea program via `program.Send()`. The log viewport appends each message.

**Alternative considered**: Buffering logs and polling. Rejected because `program.Send()` integrates cleanly and updates the UI immediately.

### 6. Active view tracking — enum field on model

Replace gocui's `SetCurrentView`/`CurrentView` with an `activeView int` field (0=repo, 1=status, 2=log). Tab cycles this value. The active panel gets a highlighted border style via lipgloss.

### 7. Bubbles components used

| Panel | Component | Rationale |
|-------|-----------|-----------|
| Repo list | Custom list ([]string + cursor int) | `bubbles/list` is overkill for a simple string list with arrow navigation |
| Status | `bubbles/viewport` | Scrollable text content, already handles up/down |
| Log | `bubbles/viewport` | Autoscroll text, wrapping |
| Scanning modal | `bubbles/spinner` | Animated progress indicator |

### 8. File organization — keep single file

Keep all UI code in `ui/ui.go`. The current 382 lines will become roughly similar in size. If it grows significantly in the future, it can be split then.

## Risks / Trade-offs

- **[Behavioral parity]** The gocui cursor-bouncing scan animation is quirky — the Bubbletea version will use a standard spinner instead. This is a minor visual difference. → Accept; spinner is an improvement.
- **[Log writer threading]** `program.Send()` is thread-safe, but the log writer will be called from the scanner goroutine. → This is explicitly supported by Bubbletea's API.
- **[Go version bump]** Users on Go <1.18 can no longer build. → Acceptable; Go 1.18 is 4+ years old.
- **[Dependency size increase]** Adding 3 Charmbracelet packages vs 1 gocui package. → These are well-maintained, widely-used packages. Worth it.
