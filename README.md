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

| Key               | Action                                           |
| ----------------- | ------------------------------------------------ |
| `<up>` / `<down>` | Navigation inside repositories or diff views     |
| `<tab>`           | switch focus between repositories and diff views |
| `e`               | Open selected repo in editor (current vscode)    |
| `s`               | Rescan directories                               |
| `q` / `ctrl-C`    | quit                                             |

Inside the "diff" view, a list of dirty files is shown, with the git status
for both staged changes (`S`) and working directory (`W`).

## Development

```bash
make lint
```

