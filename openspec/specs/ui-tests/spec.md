# ui-tests Specification

## Purpose
TBD - created by archiving change add-ui-tests. Update Purpose after archive.
## Requirements
### Requirement: Panel height calculations are correct
The panel height functions SHALL return sensible values for various terminal sizes and repo counts.

#### Scenario: repoPanelHeight with few repos
- **WHEN** terminal height is 40 and there are 3 repos
- **THEN** `repoPanelHeight` SHALL return 3 (one line per repo)

#### Scenario: repoPanelHeight capped at max
- **WHEN** terminal height is 40 and there are 100 repos
- **THEN** `repoPanelHeight` SHALL return a value less than half the terminal height

#### Scenario: repoPanelHeight with zero repos
- **WHEN** there are 0 repos
- **THEN** `repoPanelHeight` SHALL return 1 (placeholder line)

#### Scenario: logPanelHeight bounded
- **WHEN** terminal height is 40
- **THEN** `logPanelHeight` SHALL return a value between 1 and 10

#### Scenario: statusPanelHeight fills remaining space
- **WHEN** terminal height is 40
- **THEN** `statusPanelHeight` SHALL return a positive value that, combined with repo and log panel heights plus borders, does not exceed terminal height

### Requirement: Repo list rendering includes correct paths
The `renderRepoList` function SHALL render the visible repo paths based on cursor position.

#### Scenario: All repos visible
- **WHEN** there are 3 repos and terminal is large enough to show all
- **THEN** `renderRepoList` output SHALL contain all 3 repo paths

#### Scenario: Current repo highlighted
- **WHEN** cursor is on the second repo
- **THEN** `renderRepoList` output SHALL contain the second repo's path

#### Scenario: Empty repo list
- **WHEN** there are 0 repos
- **THEN** `renderRepoList` SHALL return a message indicating no dirty repos found

