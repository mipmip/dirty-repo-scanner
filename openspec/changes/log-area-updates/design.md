## Context

The TUI has three panels: repo list (top), status (middle), and log (bottom). The log panel is always visible with a fixed height of `min(10, (height-6)/3)`. Navigation within viewports is limited to single-line scrolling via j/k or arrow keys. There is no chord-based input (multi-key sequences like `gg`).

## Goals / Non-Goals

**Goals:**
- Hide log panel by default to give more space to repo and status panels
- Provide fast scrolling via page-up/page-down and ctrl+f/ctrl+b
- Provide jump-to-top (`gg`) and jump-to-bottom (`G`) navigation
- Update nav bar to reflect new keybindings

**Non-Goals:**
- Changing the log capture mechanism or log formatting
- Adding filtering or search within the log
- Persisting log visibility state across sessions

## Decisions

### 1. Log visibility toggle key: `l`

Use `l` to toggle log panel visibility. It's mnemonic (l for log), not currently bound, and reachable from home row.

**Alternative**: `L` (shift+l) — rejected because simple single key is faster and consistent with other bindings.

### 2. Hidden log still captures output

When hidden, the `logContent` string and `logWriter` continue to accumulate messages. The viewport is simply not rendered and its height is set to 0. This way, toggling the log back on shows full history.

**Alternative**: Stop capturing when hidden — rejected because users would lose log context.

### 3. Layout recalculation on toggle

When log is hidden, `recalcLayout()` sets log height to 0 and redistributes the space to the status panel. The repo panel height stays the same (it's based on repo count).

### 4. Chord input for `gg`

Add a `pendingKey` field to the model. When `g` is pressed, set `pendingKey = "g"`. On the next keypress, if it's also `g`, jump to top; otherwise, clear `pendingKey` and process the key normally. This is a minimal chord implementation — no timeout needed since we handle it on next keypress.

**Alternative**: Timer-based chord detection — rejected as overengineered for a single chord.

### 5. Half-page scrolling

Page-up/page-down and ctrl+f/ctrl+b scroll by half the viewport height, matching vim's ctrl+u/ctrl+d behavior. This applies to whichever viewport is currently active (status or log).

### 6. Active view cycling skips log when hidden

When log is hidden, Tab cycles only between repo and status views. If the log was the active view when hidden, focus moves to repo.

## Risks / Trade-offs

- **Chord state complexity**: The `pendingKey` field adds state to the model. Mitigation: it's a single string field cleared on every non-matching keypress — minimal complexity.
- **Hidden log may surprise users**: New users won't see log output. Mitigation: nav bar shows `l: log` hint, and log errors could still surface via the error modal.
