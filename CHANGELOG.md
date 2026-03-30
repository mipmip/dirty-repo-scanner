# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.2] - 2026-03-30
- fix: flake
- fix: tmux popup should not block tui
- feat: log panel hidden by default, toggle with `l` key
- feat: page-up/page-down and ctrl+f/ctrl+b for half-page scrolling
- feat: `gg` to jump to top, `G` to jump to bottom
- feat: updated nav bar with new keybinding hints
- feat: selectable files in status panel with cursor navigation
- feat: inline git diff preview for selected file (split view)
- feat: `e` key to edit selected file (tmux popup / xdg-open / $EDITOR)
- feat: `enter` opens terminal, `j`/`k` for vim-style navigation
- feat: syntax highlighting for diff output (green/red/cyan/yellow)
- feat: show version number in nav bar

## [0.2.1] - 2026-03-27

- **BREAKING**: Config path moved from `~/.dirty-repo-scanner.yml` to `~/.config/dirty-repo-scanner/config.yml` (XDG compliant). Move your config file manually.
- Drop `go-homedir` dependency in favor of `os.UserConfigDir()`
- Detect running in tmux and then run e tmux popup

## [0.2.0] - 2026-03-27

- Rename project from dirtygit to dirty-repo-scanner (binary: drs)
- Add testing framework with scanner and UI tests
- Add test coverage reporting (make cover)
- Refactor source code under src/ directory
- Add navigation bar with keybindings and app name
- Add changelog, versioning, and release workflow
- Migrate to bubbletea TUI framework

## [0.1.0] - 2026-03-26 (virtual)

- Fork from boyvinall/dirtygit
- Add Nix flake and package.nix
- Ignore directory errors by default
- simple build task
- dir wildcards
- exclude part of paths
- custom edit command
- Allow configuration of editor
- flake
