## 1. Tmux Detection

- [x] 1.1 Add an `inTmux` field to the `model` struct in `src/ui/ui.go`, set in `newModel` based on `os.Getenv("TMUX")`

## 2. Popup Editor

- [x] 2.1 Modify `doEdit()` to branch: if `m.inTmux`, build a `tmux display-popup -E -w 80% -h 80% -- <edit_command>` command; otherwise keep current behavior

## 3. Verification

- [x] 3.1 Run `make build` and verify the binary compiles
- [x] 3.2 Run `make test` and verify all tests pass
