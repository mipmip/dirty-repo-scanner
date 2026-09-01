## MODIFIED Requirements

### Requirement: Keybinding hints on the left
The navigation bar SHALL display available keybindings with their actions on the left side, including the log toggle, scroll, and jump bindings. The keybinding hints SHALL include: q (quit), s (scan), enter (open), tab (switch), jk/arrows (navigate), l (log), pgup/pgdn (scroll), gg/G (jump).

#### Scenario: Keybindings displayed
- **WHEN** the TUI is running
- **THEN** the left side of the nav bar SHALL show key-action pairs for: quit, scan, edit, switch panels, navigate, log toggle, scroll (pgup/pgdn), and jump (gg/G)
