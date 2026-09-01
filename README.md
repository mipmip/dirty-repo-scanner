# dirty-repo-scanner

Do you find yourself context-switching between a bunch of different git repos?

Have you ever accidentally discovered that changes you've made locally have not
been committed or pushed to your git server?

`drs` (dirty-repo-scanner) is a text-mode UI tool to find git repos that have uncommitted files or which have not
been pushed to a remote.

## Source-mode installation

```bash
go install github.com/mipmip/dirty-repo-scanner@master
```

## Configuration

Copy [config.yml](src/config.yml) to `~/.config/dirty-repo-scanner/config.yml` and edit to your needs.

The config path follows the XDG Base Directory Specification. If `$XDG_CONFIG_HOME` is set, the config is read from `$XDG_CONFIG_HOME/dirty-repo-scanner/config.yml`.

## Running

```bash
drs [ <directories...> ]
```

If one/more directories are specified as `<directories>`, then this will override the
`scandirs.include` from your config file.

![demo](demo.gif)

## UI

Simple key navigation in the UI as follows:

| Key                        | Action                                           |
| -------------------------- | ------------------------------------------------ |
| `j`/`k` or `<up>`/`<down>` | Navigation inside repositories or diff views     |
| `<tab>`                    | switch focus between repositories and diff views |
| `<enter>`                  | Open terminal in selected repo directory         |
| `s`                        | Rescan directories                               |
| `q` / `ctrl-C`             | quit                                             |

Inside the "diff" view, a list of dirty files is shown, with the git status
for both staged changes (`S`) and working directory (`W`).

## Multiplex mode

Run `drs --multiplex` to turn the scanner into a quick repo switcher. In this
mode, pressing `<enter>` on a repo runs the configured `switch_command` and then
quits, instead of opening the editor. This is meant to be launched inside a tmux
popup so selecting a repo switches your session to it.

Configure the command in `config.yml`:

```yaml
switch_command: tmux switch-client -t %REPO_NAME
```

Two placeholders are substituted for the selected repo before the command runs:

| Placeholder          | Value                              |
| -------------------- | ---------------------------------- |
| `%WORKING_DIRECTORY` | the repository's absolute path     |
| `%REPO_NAME`         | the repository directory base name |

The command is split with shell-quoting rules and executed without a shell, so
quoted arguments and paths containing spaces are preserved. `--multiplex`
requires `switch_command` to be set, otherwise `drs` exits with an error.

Example tmux binding that opens the scanner in a popup:

```tmux
bind-key C-d display-popup -E "drs --multiplex"
```

Note: tmux rejects `.` and `:` in session names, so choose a `switch_command`
that suits your repository names when using `%REPO_NAME`.

## Development

```bash
make lint
```

