## 1. Nav Bar Style

- [x] 1.1 Add a `navBarStyle` variable in the styles section of `src/ui/ui.go` with subtle background color and full-width rendering

## 2. Render Function

- [x] 2.1 Add a `renderNavBar` method on model that returns a single-line string with keybindings left-aligned and app name right-aligned using `lipgloss.PlaceHorizontal`

## 3. Layout Adjustment

- [x] 3.1 Adjust `statusPanelHeight` to subtract 1 line for the nav bar
- [x] 3.2 Append the nav bar output after the log panel in `View()`

## 4. Verification

- [x] 4.1 Run `make build` and verify the binary compiles
- [x] 4.2 Run `make test` and verify all tests pass
