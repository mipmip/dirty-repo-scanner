## Why

The current keybindings use `e` for opening a terminal/editor, which isn't intuitive. `enter` is the standard "activate/open" key. Additionally, vim-style `j`/`k` navigation keys are missing, which is expected in TUI apps.

## What Changes

- Change the "open in terminal" keybinding from `e` to `enter`
- Rename the action from "edit" to "open" (it opens a terminal shell, not an editor)
- Add `j` (down) and `k` (up) as alternative navigation keys alongside arrow keys
- Update the nav bar to reflect the new keybindings
- Update README keybinding table

## Capabilities

### New Capabilities

- `keybinding-updates`: Updated key mappings and vim-style navigation

### Modified Capabilities

(none)

## Impact

- **Code**: `src/ui/ui.go` — Update() key handling and nav bar rendering
- **Docs**: README keybinding table
