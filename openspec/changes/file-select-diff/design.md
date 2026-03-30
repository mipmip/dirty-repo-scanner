## Context

The status panel (`viewStatus`) currently renders a sorted list of dirty files using `updateStatusContent()` into a `viewport.Model`. Navigation within the status panel only scrolls the viewport. There's no concept of a selected file.

The model already tracks `activeView` (repo, status, log) and handles `up`/`down`/`j`/`k` per active view.

## Goals / Non-Goals

**Goals:**
- File cursor in the status panel (separate from viewport scroll)
- Split the status panel: left = file list with cursor, right = git diff output for selected file
- `e` keybinding to edit the selected file (tmux popup / xdg-open / $EDITOR fallback)
- Nav bar shows `e:edit` contextually when status panel is active

**Non-Goals:**
- Staging/unstaging files from the TUI
- Inline diff editing
- Syntax highlighting for diff output

## Decisions

### 1. Add fileCursor and filePaths to the model

New fields: `fileCursor int`, `filePaths []string` (sorted file paths for the current repo). Updated when `updateStatusContent()` runs.

### 2. Split status panel rendering into two columns

When a file is selected (status panel active and files exist), render:
- Left column (~40% width): file list with cursor highlighting (reuse `selectedStyle`)
- Right column (~60% width): a new `diffViewport` showing `git diff` output for the selected file

Use `lipgloss.JoinHorizontal` to combine them.

### 3. Diff viewport for git diff output

Add a `diffViewport viewport.Model` to the model. When `fileCursor` changes, run `git diff -- <file>` in the repo directory and set the viewport content. For untracked files, show the file contents or a "new file" message.

### 4. Run git diff as a tea.Cmd

Fetch diff asynchronously via a `tea.Cmd` to avoid blocking the UI. Return a `diffMsg` with the output.

### 5. File edit with `e` keybinding

When status panel is active and a file is selected:
- Build full file path: `repoPath + "/" + filePath`
- In tmux: `tmux display-popup -E -w 80% -h 80% -- $EDITOR <fullpath>`
- Outside tmux: try `xdg-open <fullpath>`, fall back to `$EDITOR <fullpath>`
- Use `tea.ExecProcess` for non-tmux editors (they need terminal control)

### 6. Nav bar context

Add `e:edit` to the nav bar keys. It's always shown (applicable in both repo and status views, with different behavior).

## Risks / Trade-offs

- **[Large diffs]** Very large git diffs could be slow to render. → The viewport handles scrolling; we just set the content. Acceptable for now.
- **[Untracked files]** `git diff` shows nothing for untracked files. → Show "Untracked file" as the diff content.
- **[$EDITOR not set]** If `$EDITOR` is empty, the edit action does nothing. → Acceptable; standard Unix expectation.
