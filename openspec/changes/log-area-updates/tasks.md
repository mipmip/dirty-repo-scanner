## 1. Log Toggle

- [x] 1.1 Add `logVisible bool` field to model (default `false`)
- [x] 1.2 Handle `l` keypress to toggle `logVisible` and call `recalcLayout()`
- [x] 1.3 Update `recalcLayout()` to set log viewport height to 0 when hidden, giving space to status panel
- [x] 1.4 Update `View()` to skip rendering log panel when `logVisible` is false
- [x] 1.5 Update Tab cycling to skip log view when hidden; move focus to repo if log is active when hidden

## 2. Chord Input (gg)

- [x] 2.1 Add `pendingKey string` field to model
- [x] 2.2 Handle `g` keypress: if `pendingKey == "g"`, execute jump-to-top and clear; otherwise set `pendingKey = "g"`
- [x] 2.3 Clear `pendingKey` and process key normally when a non-`g` key follows a pending `g`

## 3. Viewport Jump (gg / G)

- [x] 3.1 Implement `gg` jump-to-top for status and log viewports (`GotoTop()`)
- [x] 3.2 Implement `gg` cursor-to-first for repo panel (`cursor = 0`)
- [x] 3.3 Implement `G` jump-to-bottom for status and log viewports (`GotoBottom()`)
- [x] 3.4 Implement `G` cursor-to-last for repo panel (`cursor = len(repoPaths)-1`)

## 4. Fast Scrolling (pgup/pgdn, ctrl+f/ctrl+b)

- [x] 4.1 Handle `pgdown` and `ctrl+f` to scroll active viewport down by half viewport height
- [x] 4.2 Handle `pgup` and `ctrl+b` to scroll active viewport up by half viewport height
- [x] 4.3 For repo panel, move cursor by half viewport height (clamped to bounds)

## 5. Nav Bar Update

- [x] 5.1 Add new keybinding hints to nav bar: `l` (log), `pgup/pgdn` (scroll), `gg/G` (jump)

## 6. Tests

- [x] 6.1 Test layout with log hidden: status panel fills freed space
- [x] 6.2 Test layout with log visible: same as current behavior
- [x] 6.3 Test Tab cycling skips hidden log panel
