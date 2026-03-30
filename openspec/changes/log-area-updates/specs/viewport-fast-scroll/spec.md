## ADDED Requirements

### Requirement: Page-down and page-up scrolling
The user SHALL be able to scroll the active viewport by half a page using page-down and page-up keys.

#### Scenario: Page down in status panel
- **WHEN** the status panel is active and the user presses page-down
- **THEN** the viewport scrolls down by half the viewport height

#### Scenario: Page up in status panel
- **WHEN** the status panel is active and the user presses page-up
- **THEN** the viewport scrolls up by half the viewport height

#### Scenario: Page scrolling in log panel
- **WHEN** the log panel is active and the user presses page-down or page-up
- **THEN** the log viewport scrolls by half the viewport height in the corresponding direction

### Requirement: Vim-style half-page scroll with ctrl+f and ctrl+b
The user SHALL be able to scroll the active viewport by half a page using ctrl+f (forward) and ctrl+b (backward).

#### Scenario: Ctrl+f scrolls forward
- **WHEN** the active viewport has scrollable content and the user presses ctrl+f
- **THEN** the viewport scrolls down by half the viewport height

#### Scenario: Ctrl+b scrolls backward
- **WHEN** the active viewport has scrollable content and the user presses ctrl+b
- **THEN** the viewport scrolls up by half the viewport height

### Requirement: Scroll bounds
Scrolling SHALL NOT move the viewport beyond the top or bottom of the content.

#### Scenario: Page down at end of content
- **WHEN** the viewport is at or near the bottom of content and the user presses page-down
- **THEN** the viewport scrolls to the bottom and does not exceed it
