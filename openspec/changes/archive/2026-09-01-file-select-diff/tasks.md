## 1. Model Changes

- [x] 1.1 Add `fileCursor int`, `filePaths []string`, and `diffViewport viewport.Model` fields to the model struct
- [x] 1.2 Initialize `diffViewport` in `newModel()`
- [x] 1.3 Add `diffMsg` message type for async diff results
- [x] 1.4 Update `updateStatusContent()` to populate `filePaths` from the sorted file list and reset `fileCursor`

## 2. File Navigation

- [x] 2.1 In `Update()`, when `activeView == viewStatus`, handle `j`/`k`/`up`/`down` to move `fileCursor` instead of scrolling the viewport
- [x] 2.2 When `fileCursor` changes, trigger a `tea.Cmd` that runs `git diff -- <file>` and returns a `diffMsg`
- [x] 2.3 Handle `diffMsg` in `Update()` to set `diffViewport` content (or "Untracked file" for `?` status)

## 3. Split Status Panel Rendering

- [x] 3.1 Replace `updateStatusContent()` viewport rendering with a custom `renderStatusPanel()` that draws the file list with cursor highlighting on the left
- [x] 3.2 Render the `diffViewport` on the right side using `lipgloss.JoinHorizontal`
- [x] 3.3 Update `recalcLayout()` to set `diffViewport` dimensions
- [x] 3.4 Update `View()` to use the new split rendering for the status panel

## 4. File Edit Action

- [x] 4.1 Add `doEditFile()` method that opens the selected file: tmux popup with `$EDITOR`, else `xdg-open`, else `$EDITOR` via `tea.ExecProcess`
- [x] 4.2 Wire `"e"` keybinding in `Update()` to call `doEditFile()` when status panel is active

## 5. Nav Bar

- [x] 5.1 Add `e:edit` to the nav bar key hints

## 6. Verification

- [x] 6.1 Run `make build` and `make test`
