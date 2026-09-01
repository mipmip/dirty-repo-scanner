# HupHop — agent build guide

HupHop is a multi-provider git portfolio manager (GitHub, Codeberg/Forgejo,
GitLab). It gives a uniform on-disk layout, a Bubble Tea TUI to browse/clone,
and a cron-friendly CLI to sync backups.

**Read [BRIEFING.md](./BRIEFING.md) first — it is the single source of truth**
for the design (config schema, clone-path template, cache, sync semantics,
scope, coverage gate). The condensed version is also in
`openspec/config.yaml` `context:`.

## Tooling

- **Language**: Go (module `github.com/mipmip/huphop`), Bubble Tea for the TUI.
- **Packaging**: Nix flakes, plain nix (no flake-utils), multi-arch. The PoC
  must build and test under nix from the start.
- **Tickets**: `beans` (run `beans prime` to learn it). Milestones are titled
  `01`…`05`; epics hang under milestones. Administer them as you work.
- **Specs**: OpenSpec (spec-driven schema). One change per epic.
- **VCS**: `jj`. Remote: `git@github.com:mipmip/huphop.git`.
  Commit as **Pim Snel**, **no self-promotion**. Commit after every archived
  OpenSpec change.

## Commands

```bash
# Nix (authoritative)
nix build              # build packages.default
nix run . -- version   # run the CLI
nix develop            # dev shell (go, gopls, golangci-lint, jj)
nix flake check        # build + tests (+ coverage gate from milestone 05)

# Go (inside dev shell)
go build ./...
go test ./...
go test -cover ./...
gofmt -l .             # must be empty
bash scripts/coverage.sh   # coverage gate: overall >=70%, core >=80%
```

The coverage gate (`scripts/coverage.sh`) is enforced by the
`checks.coverage` flake output, so `nix flake check` fails below threshold.

When you add third-party dependencies, update `flake.nix` `vendorHash` (nix
prints the expected hash on the first failing build).

## Autonomous build loop

Work milestone by milestone, epic by epic, in order. For each epic:

1. **Pick the next ready epic**: `beans list --json --ready` (lowest number
   first). Mark it in-progress: `beans update <id> -s in-progress`.
2. **Propose**: run `/opsx:propose "<epic goal>"` to create an OpenSpec change
   with proposal, design, specs and tasks. Ground it in BRIEFING.md.
3. **Apply**: run `/opsx:apply` — implement the tasks with thorough unit tests.
   Keep the epic's beans checklist current (`- [ ]` → `- [x]`).
4. **Verify**: `nix flake check` must pass (build + tests + coverage gate).
   `gofmt -l .` must be empty.
5. **Archive**: `openspec archive <change>` to fold specs into the main specs.
6. **Commit**: `jj` commit (code + beans + openspec) as Pim Snel, no
   self-promotion; then move the `main` bookmark and `jj git push`.
7. **Close the bean**: only when every checklist item is done,
   `beans update <id> -s completed` with a `## Summary of Changes` section.

Milestone 05 is the coverage/e2e milestone: unit coverage ≥80% on core-logic
packages (config, clonepath, provider, cache, syncer), overall ≥70%, enforced
in `nix flake check`.

## Layout

```
cmd/hup/             CLI entrypoint (subcommands: tui, sync, refresh, config)
internal/config/    config.yaml loading + validation
internal/clonepath/ clone-path template engine
internal/provider/  Provider interface, Repo model, auth resolver + clients
internal/cache/     per-provider JSON cache
internal/syncer/    clone-missing + ff-pull-existing engine
internal/tui/        Bubble Tea UI
```

## Conventions

- Default clone protocol `ssh`; `base_dir` default `~`; `include_archived`
  false, `include_forks` true.
- Sync never force-pulls and skips dirty trees with a warning (safe for cron).
- Non-interactive commands are cron-friendly: line-oriented logs, non-zero
  exit on failure.
- Auth resolution: configured CLI token (`gh`/`glab`) → env PAT
  (`HUPHOP_<PROVIDER>_TOKEN`) → clear error. Codeberg is env-PAT only.
- Keep `internal/` package boundaries; no network in unit tests (mock HTTP,
  use temp git repos and temp dirs).
