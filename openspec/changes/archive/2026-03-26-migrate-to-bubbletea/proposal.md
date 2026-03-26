## Why

The TUI is built on `jroimartin/gocui`, which has been unmaintained since ~2018. It lacks modern terminal features (true color, rich styling, component ecosystem) and uses a coordinate-based layout model that makes UI changes cumbersome. Migrating to Bubbletea + Lipgloss gives us an actively maintained framework with the Elm architecture, a rich component library (bubbles), and string-based composable layouts — setting the foundation for future UI improvements.

## What Changes

- **BREAKING**: Minimum Go version bumped from 1.16 to 1.18+
- Replace `gocui` with `bubbletea` (TUI framework), `lipgloss` (styling), and `bubbles` (components)
- Rewrite `ui/ui.go` using Bubbletea's Model/Update/View pattern
- Update `main.go` to launch a `tea.Program` instead of `gocui.Gui`
- Preserve the existing 3-panel layout (repo list, status/diff, log) and all current keybindings
- Scanner layer (`scanner/`) remains completely untouched

## Capabilities

### New Capabilities
- `bubbletea-tui`: The core TUI built on Bubbletea, replacing gocui. Covers the Model, Update loop, View rendering, keybindings, and panel layout.

### Modified Capabilities

(none — no existing specs)

## Impact

- **Dependencies**: Remove `github.com/jroimartin/gocui`. Add `github.com/charmbracelet/bubbletea`, `github.com/charmbracelet/lipgloss`, `github.com/charmbracelet/bubbles`.
- **Code**: `ui/ui.go` is fully rewritten. `main.go` has minor changes to program initialization.
- **Build**: `go.mod` bumps to `go 1.18` minimum.
- **Scanner**: No changes. The `scanner` package interface remains the same.
- **User-facing**: Identical behavior and layout. Same keybindings, same panels, same workflow.
