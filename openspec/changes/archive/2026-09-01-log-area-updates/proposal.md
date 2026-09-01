## Why

The log panel is always visible and takes up screen space even when not needed. Users also lack efficient navigation within the log — only single-line scrolling with j/k is available, and there's no way to jump to the top or bottom.

## What Changes

- Log panel is hidden by default; a keybinding toggles its visibility
- Page-down/page-up and ctrl+f/ctrl+b (vim) bindings for half-page scrolling in all viewports
- `gg` chord to jump to top and `G` to jump to bottom of the active viewport

## Capabilities

### New Capabilities
- `log-toggle`: Toggle log panel visibility with a keybinding; log is hidden by default
- `viewport-fast-scroll`: Page-up/page-down and ctrl+f/ctrl+b for half-page scrolling in all viewports
- `viewport-jump`: gg to jump to top, G to jump to bottom in any viewport

### Modified Capabilities
- `nav-bar`: Nav bar must display the new keybindings (l for log toggle, pgup/pgdn, gg/G)

## Impact

- `src/ui/ui.go`: Layout logic changes (log panel height becomes 0 when hidden), new key handlers, chord state for `gg`
- `src/ui/ui_test.go`: Tests for layout with hidden log, new keybinding behavior
- Nav bar: Updated keybinding hints
