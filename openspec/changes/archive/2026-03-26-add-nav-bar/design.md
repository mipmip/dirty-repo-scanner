## Context

The TUI uses bubbletea with three panels (Repositories, Status, Log) stacked vertically. The `View()` method in `src/ui/ui.go` joins these panels and optionally overlays modals. Panel heights are calculated in `repoPanelHeight`, `statusPanelHeight`, and `logPanelHeight` based on `m.height`.

## Goals / Non-Goals

**Goals:**
- A single-line bar at the very bottom of the terminal
- Left-aligned: key hints (key + action pairs)
- Right-aligned: app name `dirty-repo-scanner`
- Consistent styling (subtle, doesn't distract from content)

**Non-Goals:**
- Dynamic/contextual keybinding hints (showing different keys per view)
- Clickable elements

## Decisions

### 1. Render as a plain styled string, not a panel

The nav bar is a single line with no border. Use lipgloss to style it with a subtle background color and place left/right content using `lipgloss.PlaceHorizontal`.

**Why**: A bordered panel for 1 line wastes space. A flat bar is the standard convention (like vim's status line).

### 2. Reserve 1 line from available height

Subtract 1 from `m.height` in the `statusPanelHeight` calculation (which absorbs remaining space). The nav bar is appended after the log panel in `View()`.

**Why**: The status panel is the flexible one — it already takes whatever space remains. Reducing it by 1 is the least disruptive change.

### 3. Style keybindings as `key:action` pairs separated by spaces

Format: `q:quit  s:scan  e:edit  tab:switch  ↑↓:navigate`

**Why**: Compact, scannable, common convention in TUI apps.

## Risks / Trade-offs

- **[1 line less for panels]** Slightly reduces content area. → Negligible at typical terminal sizes (80x24+). The minimum height check (20 lines) already provides headroom.
