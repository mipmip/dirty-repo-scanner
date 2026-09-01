---
# dirty-repo-scanner-8yog
title: multiplex mode
status: completed
type: feature
priority: normal
created_at: 2026-09-01T20:21:21Z
updated_at: 2026-09-01T21:43:54Z
---

I want to implement simular functionality described in /home/pim/gh.mipmip/huphop/openspec/changes/archive/2026-08-19-add-tui-modes/

so a multiplex mode with configurable enter action, have a look at the current configuration of huphop in which the enter command is defined

After this change I'll configure a tmux popup witch switches to the dirty repo dir

## Summary of Changes

Added a `--multiplex` mode ported (and simplified) from huphop's configurable
enter action. In multiplex mode, `<enter>` runs a configurable `switch_command`
for the selected repo and quits, so `drs` can drive a tmux popup that switches to
the chosen repo.

- New `switch_command` config key with `%WORKING_DIRECTORY` and `%REPO_NAME`
  placeholders.
- Commands are split with in-repo POSIX shell-quoting rules and run without a
  shell (no external dependency, so no `vendorHash` change).
- `--multiplex` requires `switch_command`, else exits non-zero before the UI.
- Success quits (exit 0); a failed switch stays open with an error overlay.
- README documents usage + an example tmux popup binding; CHANGELOG updated.
- OpenSpec change `add-multiplex-mode` archived as `2026-09-01-add-multiplex-mode`.
