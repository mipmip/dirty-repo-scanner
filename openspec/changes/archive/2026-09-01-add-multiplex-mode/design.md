## Context

dirty-repo-scanner is a single-binary Go/Bubble Tea TUI. The Enter key is
handled by `model.doEdit()` in `src/ui/ui.go`, which substitutes
`%WORKING_DIRECTORY` into `edit_command` and either opens a tmux popup (when
`$TMUX` is set) or execs the editor. Config is a flat `scanner.Config` struct
(`src/scanner/scan.go`) loaded from an XDG path with an embedded default
(`src/config.yml`). The CLI (`src/main.go`, `urfave/cli/v2`) has no subcommands.

This change ports huphop's "multiplex mode" — a configurable Enter action — but
drops everything in huphop that assumes remote repos (clone-if-needed, progress
popups, a `modes:` chrome map, clonepath template fields). Here every repo is
already on disk, so multiplex Enter reduces to: render a command against the
selected path, exec it without a shell, and quit.

## Goals / Non-Goals

**Goals:**
- A `--multiplex` flag that repurposes Enter to run a configurable
  `switch_command`, enabling a tmux popup that switches to the selected repo dir.
- Safe execution: shell-words splitting, no shell interpretation of repo paths.
- Backward compatibility: without the flag, nothing changes.

**Non-Goals:**
- No `modes:` map, `default_mode`, or in-TUI mode switching — one flag, one key.
- No chrome/layout changes (huphop's element registry is out of scope).
- No Go `text/template` engine — literal `%TOKEN` replacement matches the
  existing `edit_command` convention.
- No automatic sanitization of `%REPO_NAME` for tmux (documented caveat instead).

## Decisions

### Selection: a boolean `--multiplex` flag (not `--mode <name>`)
Because there is a single alternative behavior and no per-mode chrome, a boolean
flag is the honest surface. It is threaded from `main.go` into `ui.Run(...)` and
stored on `model`. Alternative considered: huphop's `--mode <name>` + `modes:`
map — rejected as overkill; it can be layered on later without breaking the flag.

### Config: one flat `switch_command` string key
Add `SwitchCommand string \`yaml:"switch_command"\`` to `scanner.Config`. No
nested map. This mirrors the existing flat `edit_command` and keeps the embedded
default readable. Validation of "multiplex requires switch_command" lives at
launch (see below), not in config parsing, so a config with `switch_command`
unset is still valid for normal use.

### Placeholders: literal `%WORKING_DIRECTORY` + `%REPO_NAME`
Reuse the existing `strings.Replace` substitution style for consistency with
`edit_command`. `%WORKING_DIRECTORY` → the repo's absolute path (already used);
`%REPO_NAME` → `filepath.Base(path)`, useful for tmux session names. Substitution
is done once against the rendered string before argv splitting.

### Execution: shell-words split, exec without a shell
Replace naive `strings.Fields` (in the multiplex path) with a POSIX shell-words
splitter — `github.com/mattn/go-shellwords` (small, widely used, MIT). The
rendered command is split into argv and run via `exec.Command(argv[0], argv[1:]...)`,
so a repo path with spaces or metacharacters stays a literal argument. This adds
one dependency and requires bumping `flake.nix` `vendorHash` (nix prints the
expected hash on the first failing build) plus `go.mod`/`go.sum`. The existing
`edit_command` path may keep `strings.Fields` or be migrated to the splitter too;
unifying is preferred but not required by the spec.

### Enter dispatch: branch inside the handler, quit on success
`doEdit()` grows a guard: when `m.multiplex` is set, call a new `doSwitch()`
that renders/splits/execs `switch_command` and, on success, returns `tea.Quit`.
On failure it sets an error status and stays open. Non-multiplex flow is
untouched. The switch command (e.g. `tmux switch-client`) is non-interactive, so
it runs via a plain exec wired to stdio rather than `tea.ExecProcess`.

### Fail-fast on empty switch_command
When `--multiplex` is set but `switch_command` is empty, `ui.Run` (or `main.go`)
returns an error before starting the program, printing a clear message and
exiting non-zero — consistent with the project's cron-friendly convention.

## Risks / Trade-offs

- **New dependency + vendorHash churn** → Mitigation: a single tiny, stable
  library; the flake hash bump is a known, mechanical step already documented in
  CLAUDE.md.
- **`%REPO_NAME` may contain characters tmux rejects in session names** (`.`, `:`)
  → Mitigation: document the caveat and let the user's `switch_command` handle
  naming; do not silently rewrite the value.
- **Two substitution/splitting code paths** (edit vs switch) could drift →
  Mitigation: factor a shared `renderCommand(template, repo)` + `splitArgs`
  helper and use it in both.
- **`tmux switch-client` from inside a popup targets the attached client** — a
  property of the user's tmux binding, not this code → Mitigation: ship a working
  example binding in the README.
