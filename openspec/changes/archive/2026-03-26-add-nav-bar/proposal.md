## Why

The TUI has no persistent indication of available keybindings. Users must remember shortcuts or refer to the README. A navigation bar at the bottom provides discoverability and gives the app a polished look.

## What Changes

- Add a single-line navigation bar at the bottom of the TUI
- Left side: keybindings with their actions (e.g., `q quit  s scan  e edit  tab switch  ↑↓ navigate`)
- Right side: application name (`dirty-repo-scanner`)
- The nav bar takes 1 line of height, reducing available panel space by 1 line

## Capabilities

### New Capabilities

- `nav-bar`: Bottom navigation bar showing keybindings and application name

### Modified Capabilities

(none)

## Impact

- **Code**: `src/ui/ui.go` — new render function for the nav bar, layout adjustment to reserve 1 line at the bottom
- **Layout**: All panel height calculations must account for the nav bar line
