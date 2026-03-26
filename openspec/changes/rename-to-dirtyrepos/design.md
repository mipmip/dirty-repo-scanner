## Context

The project is a fork of `boyvinall/dirtygit` that has diverged significantly. All references to "dirtygit" need to change: the project becomes "dirty-repo-scanner" with a short binary name `drs`. Git history is preserved — this is a rename, not a new repo.

## Goals / Non-Goals

**Goals:**
- Rename all project identifiers from `dirtygit` to `dirty-repo-scanner`
- Use `drs` as the binary/command name
- Rename the GitHub repository (preserving history and automatic redirect)
- Ensure the project builds and runs under the new name

**Non-Goals:**
- Backward compatibility shims (no symlinks, no fallback config paths)
- Changing any functionality or behavior
- Updating external references (other repos, blog posts, etc.)

## Decisions

### 1. Separate project name and binary name

The project/module is `dirty-repo-scanner` (descriptive) while the binary is `drs` (short for daily use). The `app.Name` in CLI output uses the full name, but the compiled binary is `drs`.

**Why**: Descriptive names are good for discovery and documentation. Short names are good for typing. Both can coexist.

### 2. Clean rename, no backward compatibility

All references change in one shot. No fallback to `~/.dirtygit.yml` or dual-name support.

**Why**: This is a small personal project. Adding compatibility layers adds complexity for no real benefit.

### 3. Rename GitHub repo via GitHub settings

Use GitHub's built-in repo rename (Settings → General → Repository name) rather than creating a new repo.

**Why**: This preserves all git history, issues, stars, and creates an automatic redirect from the old URL.

### 4. Config file rename: `.dirtygit.yml` → `.dirty-repo-scanner.yml`

Both the embedded default and the default config path change. Uses the full project name for the config file (not the short binary name) for clarity.

**Why**: Consistent with the project name. Config files are rarely typed manually so the longer name is fine.

## Risks / Trade-offs

- **[Nix cache invalidation]** Users with the old flake in their system will need to update their flake input URL. → Acceptable; GitHub redirect helps with the transition.
- **[Config file not found]** Existing users won't find their config after rename. → Document in README that `~/.dirtygit.yml` should be renamed to `~/.dirty-repo-scanner.yml`.
