## 1. Shell-words splitting

- [x] 1.1 Implement an in-repo POSIX shell-quoting splitter (`splitArgs` in `src/ui/multiplex.go`) instead of adding an external dependency — no `go.mod`/`go.sum` churn
- [x] 1.2 No `flake.nix` `vendorHash` change needed (no new dependency)

## 2. Config

- [x] 2.1 Add `SwitchCommand string \`yaml:"switch_command"\`` to `scanner.Config` in `src/scanner/scan.go`
- [x] 2.2 Add a commented `switch_command` example to the embedded `src/config.yml` (tmux popup switcher)

## 3. Command rendering & execution helper

- [x] 3.1 Add a helper that substitutes `%WORKING_DIRECTORY` and `%REPO_NAME` (`filepath.Base`) into a template string for a given repo path
- [x] 3.2 Add a helper that splits a rendered command into argv via `shellwords` and returns an error on empty/unparseable input
- [x] 3.3 Refactor the existing `edit_command` path in `doEdit()` to reuse these helpers (keep current behavior)

## 4. CLI flag

- [x] 4.1 Add a `--multiplex` boolean flag in `src/main.go`
- [x] 4.2 Thread the flag value into `ui.Run(...)` and onto the `model` struct
- [x] 4.3 Fail fast: when `--multiplex` is set and `switch_command` is empty, print a clear error and exit non-zero before starting the TUI

## 5. Multiplex Enter action

- [x] 5.1 In the Enter handler, branch to a new `doSwitch()` when the model is in multiplex mode
- [x] 5.2 `doSwitch()` renders `switch_command`, splits argv, and execs without a shell (stdio wired)
- [x] 5.3 On success return `tea.Quit` (exit 0); on failure set an error status and keep the TUI open

## 6. Tests

- [x] 6.1 Test placeholder substitution for `%WORKING_DIRECTORY` and `%REPO_NAME`
- [x] 6.2 Test shell-words splitting: quoted args and paths with spaces yield correct argv
- [x] 6.3 Test that `--multiplex` with empty `switch_command` errors and exits non-zero
- [x] 6.4 Test the Enter dispatch branches to switch vs edit based on the multiplex flag
- [x] 6.5 Test that a failing switch keeps the TUI open (no quit)

## 7. Docs & verification

- [x] 7.1 Add a README section: multiplex usage plus an example tmux popup keybinding; note the tmux session-name caveat for `%REPO_NAME`
- [x] 7.2 Run `gofmt -l .` (empty), `go test ./...`, and `nix flake check`; confirm all pass
