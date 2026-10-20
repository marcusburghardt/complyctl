# Design

## Context

See proposal.md for motivation. The current state:

- `tests/schema-validation/validate.sh` line 30 pins the Gemara CUE schema to
  `v0.23.0` via the `GEMARA_CUE_MODULE` env var default
- The org-wide Renovate runner (`org-infra/ci_renovate.yml`) auto-discovers all
  `complytime/*` repos and applies `go-toolchain-patches.json` via
  `globalExtends`, which sets `enabledManagers: ["gomod"]` and restricts updates
  to Go toolchain patch versions only
- complyctl has no local `renovate.json` today
- The CUE schema version (`github.com/gemaraproj/gemara`) and the Go library
  version (`github.com/gemaraproj/go-gemara` in `go.mod`) are from the same
  project but have independent version lines

## Goals / Non-Goals

**Goals:**

- Unblock CI immediately by aligning the CUE schema pin to the current Gemara
  release
- Automate future CUE schema version tracking so upstream releases produce
  reviewable PRs instead of surprise CI failures

**Non-Goals:**

- Changing the CUE tool version (`cue@v0.17.1` in the CI workflow)
- Enabling Renovate for general Go dependency management (the org preset
  intentionally restricts to Go toolchain patches; broader dep management is
  handled by Dependabot)
- Correlating `go-gemara` Go module versions with CUE schema versions (they
  are independent and tracked through separate mechanisms)

## Decisions

### 1. Target CUE schema version: `v1.6.0`

**Choice**: Bump directly to `v1.6.0` (latest stable release, confirmed via
GitHub releases API).

**Rationale**: The script comment said it tracked `go-gemara v0.7.0`, but
`go.mod` already has `go-gemara v0.11.0` -- the pin was already stale. Jumping
to the latest stable aligns with the current state and includes the `evidence`
field (added in `v1.3.0`) that triggered the failure.

**Alternative**: Bump to `v1.3.0` (minimum version with `evidence`). Rejected
because it would leave the pin stale for other fields added in `v1.4.0`-`v1.6.0`
and Renovate would immediately open a PR to bump further anyway.

### 2. Repo-local `renovate.json` with custom regex manager

**Choice**: Add a `renovate.json` in complyctl's repo root with a
`customManagers` entry using `customType: "regex"` to extract the version from
`GEMARA_CUE_MODULE` in `validate.sh`.

**Rationale**: Renovate's regex custom managers can match version strings in
arbitrary files, which is exactly what's needed for a shell variable that
Dependabot cannot see. The org Renovate runner already auto-discovers
`complytime/*` repos and respects repo-level config (`requireConfig: 'optional'`).

**Alternative considered**: Modify `org-infra/go-toolchain-patches.json` to add
the custom manager. Rejected because this dependency is specific to complyctl's
test infrastructure -- no other repo in the org has this pattern.

### 3. Override `enabledManagers` in repo config

**Choice**: The repo-local `renovate.json` sets
`enabledManagers: ["gomod", "custom.gemara-cue"]` to enable both the org
preset's Go toolchain manager and the new custom manager.

**Rationale**: The org preset restricts `enabledManagers` to `["gomod"]` via
`globalExtends`. Renovate's config merging allows repo-level config to override
`globalExtends` for this setting. Without this override, the custom manager
would be defined but never executed.

**Risk**: If `org-infra` later adds more enabled managers to the global preset,
this repo-level override would shadow them. Mitigated by the Renovate custom
manager name being explicit -- a reviewer seeing `enabledManagers` in the repo
config would understand why it exists.

### 4. Use `github-releases` datasource

**Choice**: Track `gemaraproj/gemara` via Renovate's `github-releases`
datasource.

**Rationale**: Gemara publishes releases on GitHub with semver tags. This is
the standard Renovate datasource for GitHub projects and requires no special
registry configuration.

## Risks / Trade-offs

- **[Repo overrides org `enabledManagers`]** -> The repo config lists its own
  `enabledManagers`, which shadows the org preset. If the org adds new managers
  later, this repo won't inherit them automatically. Mitigation: the override is
  documented in the config and carries a comment explaining why it exists.
- **[Independent version lines]** -> Bumping the CUE schema version does not
  automatically bump `go-gemara` in `go.mod` or vice versa. If Gemara
  introduces a breaking schema change that requires a corresponding `go-gemara`
  update, the Renovate PR for the CUE schema bump would fail CI (cross-repo
  integration test), signaling the need for a coordinated update. This is the
  desired behavior: CI-gated, not silent.
