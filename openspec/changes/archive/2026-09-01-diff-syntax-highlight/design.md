## Context

The `fetchDiff()` method returns raw `git diff` output as a `diffMsg`. The `Update()` handler sets this as the `diffViewport` content. The diff is in unified format with standard prefixes.

## Goals / Non-Goals

**Goals:**
- Color diff lines by type using lipgloss
- Apply coloring before setting viewport content

**Non-Goals:**
- Language-aware syntax highlighting within code (just line-level diff coloring)
- Configurable color schemes
- External dependencies (chroma, glamour, etc.)

## Decisions

### 1. Line-prefix-based coloring with lipgloss styles

Parse each line of the diff output and apply lipgloss styles based on the first character(s):

| Prefix | Color | Style |
|--------|-------|-------|
| `+` | Green (color 2) | Added line |
| `-` | Red (color 1) | Deleted line |
| `@@` | Cyan (color 6) | Hunk header |
| `diff `, `index `, `---`, `+++` | Yellow (color 3) | Diff metadata |
| other | Default | Context line |

### 2. Apply in diffMsg handler

Add a `colorizeDiff(content string) string` function. Call it in the `diffMsg` case before `SetContent()`.

**Why**: Single point of transformation. The raw diff is still fetched as plain text; coloring is a presentation concern.

## Risks / Trade-offs

- **[ANSI codes in viewport]** Lipgloss-styled strings contain ANSI escape codes which increase string length. → Viewport handles this correctly (it uses visible width calculations).
