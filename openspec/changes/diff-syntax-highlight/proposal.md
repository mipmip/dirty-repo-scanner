## Why

The git diff output in the status panel is displayed as plain monochrome text, making it hard to quickly scan additions, deletions, and hunk boundaries. Standard diff coloring (green for additions, red for deletions, cyan for hunk headers) is expected by developers.

## What Changes

- Add line-level syntax coloring to the git diff output in the diff viewport
- Added lines (`+`) in green
- Deleted lines (`-`) in red
- Hunk headers (`@@`) in cyan
- Diff metadata (`diff --git`, `index`, `---`, `+++`) in dim/yellow
- Context lines unchanged
- No new dependencies — uses existing lipgloss styles

## Capabilities

### New Capabilities

- `diff-highlight`: Syntax coloring for unified diff output using lipgloss styles

### Modified Capabilities

(none)

## Impact

- **Code**: `src/ui/ui.go` — new `colorizeDiff()` function applied to diff content before setting viewport
- **Dependencies**: None (lipgloss already available)
