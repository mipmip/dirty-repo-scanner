## 1. Diff Styles

- [x] 1.1 Add lipgloss style variables for diff line types: `diffAddedStyle` (green), `diffDeletedStyle` (red), `diffHunkStyle` (cyan), `diffMetaStyle` (yellow)

## 2. Colorize Function

- [x] 2.1 Add `colorizeDiff(content string) string` function that iterates lines and applies styles based on prefix
- [x] 2.2 Apply `colorizeDiff()` in the `diffMsg` handler before `SetContent()`

## 3. Verification

- [x] 3.1 Run `make build` and `make test`
