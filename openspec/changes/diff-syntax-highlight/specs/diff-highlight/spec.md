## ADDED Requirements

### Requirement: Added lines colored green
Lines starting with `+` (but not `+++`) in the diff output SHALL be rendered in green.

#### Scenario: Added line
- **WHEN** a diff line starts with `+` and is not a `+++` metadata line
- **THEN** the line SHALL be displayed in green

### Requirement: Deleted lines colored red
Lines starting with `-` (but not `---`) in the diff output SHALL be rendered in red.

#### Scenario: Deleted line
- **WHEN** a diff line starts with `-` and is not a `---` metadata line
- **THEN** the line SHALL be displayed in red

### Requirement: Hunk headers colored cyan
Lines starting with `@@` SHALL be rendered in cyan.

#### Scenario: Hunk header
- **WHEN** a diff line starts with `@@`
- **THEN** the line SHALL be displayed in cyan

### Requirement: Metadata lines colored yellow
Lines starting with `diff `, `index `, `---`, or `+++` SHALL be rendered in yellow.

#### Scenario: Metadata line
- **WHEN** a diff line starts with `diff --git`, `index`, `---`, or `+++`
- **THEN** the line SHALL be displayed in yellow

### Requirement: Context lines unchanged
Lines that don't match any prefix pattern SHALL be rendered in the default color.

#### Scenario: Context line
- **WHEN** a diff line starts with a space or has no recognized prefix
- **THEN** the line SHALL be displayed in the default terminal color
