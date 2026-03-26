## ADDED Requirements

### Requirement: Coverage profile generation
The Makefile SHALL include a `cover` target that generates a coverage profile and HTML report.

#### Scenario: Running make cover
- **WHEN** `make cover` is executed
- **THEN** it SHALL produce a `coverage.out` file and a `coverage.html` file in the project root

### Requirement: Coverage artifacts ignored by git
The `.gitignore` file SHALL exclude coverage output files.

#### Scenario: Coverage files not tracked
- **WHEN** `make cover` has been run and `git status` is checked
- **THEN** `coverage.out` and `coverage.html` SHALL NOT appear as untracked files
