## 1. Dependencies and Go Version

- [x] 1.1 Bump `go.mod` to `go 1.18` (or higher)
- [x] 1.2 Add bubbletea, lipgloss, and bubbles dependencies; remove gocui
- [x] 1.3 Run `go mod tidy` and verify the project compiles

## 2. Model and Initialization

- [x] 2.1 Define the top-level `model` struct with fields: repositories, config, scanning, err, activeView, cursor, viewport sub-models, log content, spinner, window size
- [x] 2.2 Implement `Init()` returning a `tea.Batch` with initial scan command and spinner tick
- [x] 2.3 Implement `ui.Run()` function that creates a `tea.Program` with alternate screen and runs it

## 3. Update Loop and Messages

- [x] 3.1 Define message types: `scanMsg`, `logMsg`, `errMsg`
- [x] 3.2 Implement `Update()` handling for `tea.KeyMsg`: q/ctrl+c (quit), s (scan), e (edit), tab (cycle view), arrow up/down (navigate)
- [x] 3.3 Implement `Update()` handling for `scanMsg`: store results, update repo list, clear scanning state
- [x] 3.4 Implement `Update()` handling for `tea.WindowSizeMsg`: recalculate panel dimensions
- [x] 3.5 Implement `Update()` handling for `logMsg`: append to log viewport
- [x] 3.6 Implement `Update()` handling for spinner tick messages during scanning

## 4. View Rendering

- [x] 4.1 Implement repo list rendering with lipgloss styling (green bg/black fg for selected item, panel border with title)
- [x] 4.2 Implement status panel rendering showing SW header and file statuses for the selected repo
- [x] 4.3 Implement log panel rendering with viewport and autoscroll
- [x] 4.4 Implement scanning modal overlay with centered spinner
- [x] 4.5 Implement error modal overlay with centered, wrapped error text
- [x] 4.6 Compose all panels with `lipgloss.JoinVertical` and dynamic height calculation

## 5. Async Scanning and Log Capture

- [x] 5.1 Implement scan command as `tea.Cmd` that calls `scanner.Scan()` and returns `scanMsg`
- [x] 5.2 Implement custom `io.Writer` that sends `logMsg` via `program.Send()`
- [x] 5.3 Set `log.SetOutput` to the custom writer before starting the program

## 6. Editor Launch

- [x] 6.1 Implement edit command using `tea.ExecProcess` with the configured edit command and `%WORKING_DIRECTORY` substitution

## 7. Update main.go

- [x] 7.1 Update `main.go` to call the new `ui.Run()` (signature unchanged, so this is just verifying it still compiles)

## 8. Verification

- [x] 8.1 Verify the application compiles and starts
- [x] 8.2 Verify all keybindings work: q, ctrl+c, s, e, tab, arrow up/down
- [x] 8.3 Verify layout matches: 3 panels with correct sizing
- [x] 8.4 Verify scan results display correctly with status codes
- [x] 8.5 Verify log output appears and autoscrolls
