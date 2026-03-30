## Context

The TUI handles keybindings in the `Update()` method's `tea.KeyMsg` switch. The nav bar renders key hints in `renderNavBar()`. The README has a keybinding table.

Current bindings: `q` quit, `s` scan, `e` edit, `tab` switch panels, `up`/`down` navigate.

## Goals / Non-Goals

**Goals:**
- `enter` opens terminal (replaces `e`)
- `j`/`k` for down/up navigation (in addition to arrow keys)
- Nav bar and README updated

**Non-Goals:**
- Adding `h`/`l` navigation (only `j`/`k` requested for up/down)
- Configurable keybindings

## Decisions

### 1. Map `enter` to the same action as old `e`

In bubbletea, `enter` is `"enter"` in `msg.String()`. Replace the `"e"` case with `"enter"`.

### 2. Add `j`/`k` alongside existing arrow keys

Add `"j"` to the `"down"` case and `"k"` to the `"up"` case. Keep arrow keys working.

### 3. Update nav bar text

Change `e:edit` to `enter:open` and `↑↓:navigate` to `jk/↑↓:navigate`.

## Risks / Trade-offs

- **[enter key conflict]** `enter` might conflict with future features (e.g., expanding a repo). → Acceptable for now; can be changed later.
