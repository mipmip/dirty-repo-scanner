## ADDED Requirements

### Requirement: Jump to top with gg
The user SHALL be able to press `g` followed by `g` to jump to the top of the active viewport.

#### Scenario: gg in status panel
- **WHEN** the status panel is active and the user presses `g` then `g`
- **THEN** the viewport scrolls to the very top of the content

#### Scenario: gg in log panel
- **WHEN** the log panel is active and the user presses `g` then `g`
- **THEN** the log viewport scrolls to the very top

#### Scenario: g followed by non-g key
- **WHEN** the user presses `g` followed by any key other than `g`
- **THEN** the pending `g` is discarded and the second key is processed normally

### Requirement: Jump to bottom with G
The user SHALL be able to press `G` (shift+g) to jump to the bottom of the active viewport.

#### Scenario: G in status panel
- **WHEN** the status panel is active and the user presses `G`
- **THEN** the viewport scrolls to the very bottom of the content

#### Scenario: G in log panel
- **WHEN** the log panel is active and the user presses `G`
- **THEN** the log viewport scrolls to the very bottom

### Requirement: gg in repo panel moves cursor
In the repo panel, `gg` SHALL move the cursor to the first repository and `G` SHALL move the cursor to the last repository.

#### Scenario: gg in repo panel
- **WHEN** the repo panel is active and the user presses `g` then `g`
- **THEN** the cursor moves to the first repository in the list

#### Scenario: G in repo panel
- **WHEN** the repo panel is active and the user presses `G`
- **THEN** the cursor moves to the last repository in the list
