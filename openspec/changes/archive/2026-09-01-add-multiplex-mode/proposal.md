## Why

dirty-repo-scanner is a great picker for "which of my repos has uncommitted
work" — but pressing Enter only ever opens an editor. Users who live in tmux
want to run it inside a popup and, on Enter, **switch their outer tmux client
to the selected repo's directory** instead. Porting huphop's "multiplex mode"
(a configurable Enter action) turns the scanner into a fast tmux repo-switcher.

## What Changes

- Add a `--multiplex` CLI flag. When set, the TUI runs in multiplex mode.
- Add a `switch_command` config key (a command template) executed by Enter in
  multiplex mode, in place of the normal editor action.
- Support two placeholders in `switch_command`, substituted per selected repo:
  - `%WORKING_DIRECTORY` — the repo's absolute path (same token `edit_command`
    already uses).
  - `%REPO_NAME` — the repo directory basename (new; handy for tmux session
    names).
- Execute the rendered command **without a shell**, splitting it into argv with
  POSIX shell-quoting rules so quoted arguments and paths containing spaces
  survive. This adds a small shell-words dependency.
- After a successful switch, the TUI quits with exit code 0 so the surrounding
  tmux popup closes.
- In multiplex mode with an empty `switch_command`, fail fast with a clear error
  and a non-zero exit (cron/script-friendly).
- Document an example `switch_command` and a tmux popup binding.

Non-goals: no `modes:` map or in-TUI mode switching (a single flag + single key
is enough); no chrome/layout changes; management-mode behavior is unchanged.

## Capabilities

### New Capabilities
- `multiplex-mode`: a configurable Enter action selected by `--multiplex` that
  renders a `switch_command` template against the selected repo, executes it
  without a shell, and quits — enabling tmux repo-switching from a popup.

### Modified Capabilities
<!-- None. Management-mode Enter/editor behavior (tmux-popup-editor) is unchanged. -->

## Impact

- **CLI** (`src/main.go`): new `--multiplex` boolean flag, threaded into `ui.Run`.
- **Config** (`src/scanner/scan.go`): new `switch_command` field on `Config`;
  embedded `src/config.yml` gains a documented example.
- **TUI** (`src/ui/ui.go`): `model` carries the multiplex flag; the Enter handler
  branches to a switch action (`%WORKING_DIRECTORY` + `%REPO_NAME` substitution,
  shell-words split, exec, quit).
- **Dependencies**: a shell-words splitter (e.g. `github.com/mattn/go-shellwords`)
  added to `go.mod`/`go.sum`; `flake.nix` `vendorHash` updated.
- **Docs**: `README.md` gains a multiplex/tmux-popup usage section.
- **Tests** (`src/ui/ui_test.go`): substitution, shell-words splitting,
  empty-command error, and mode-branch coverage.
