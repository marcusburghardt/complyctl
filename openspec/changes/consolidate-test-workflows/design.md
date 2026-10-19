# Design

## Context

See proposal.md for motivation. The org-infra reusable release
preflight auto-discovers required CI checks by parsing
`ci_local.yml` with `yq`. It reads the workflow-level `name:`
field and each job key (or job `name:` field), constructing
check names as `<workflow name> / <job name>`. complyctl
currently lacks `ci_local.yml`, so the preflight falls back
to a hardcoded `ci_checks` override in `release.yml`.

Current workflow state:

- `unit_test.yml`: jobs `buf-lint`, `unit-test`
  (triggers: push main, PR all default types)
- `e2e_test.yml`: job `e2e-test`
  (triggers: push main, PR opened+synchronize only)
- `integration_test.yml`: job `integration-test`
  (triggers: push main, PR all default types)

## Goals / Non-Goals

**Goals:**
- Enable preflight auto-discovery by providing `ci_local.yml`
- Remove the `ci_checks` override from `release.yml`
- Preserve existing job names for continuity

**Non-Goals:**
- Changing test commands or test infrastructure
- Modifying org-infra reusable workflows
- Adding new test jobs

## Decisions

### D1: Workflow `name:` field

**Decision**: Use `Local CI` as the workflow name.

**Rationale**: Must not collide with `ci_checks.yml` which uses
`name: CI`. `Local CI` is short, descriptive, and mirrors the
filename. The preflight constructs check names as
`Local CI / <job>`.

**Alternatives considered**:
- `CI` -- collides with `ci_checks.yml`
- `Tests` -- ambiguous, does not convey CI context
- `Local Tests` -- longer without added clarity

### D2: Job `name:` fields

**Decision**: Do not set explicit `name:` fields on jobs. Let
the preflight use the job keys directly.

**Rationale**: The existing workflows do not set job `name:`
fields. The keys (`unit-test`, `e2e-test`, etc.) are already
descriptive. Omitting `name:` eliminates a divergence point
between the YAML key and the displayed check name. The `yq`
expression falls back to the key when `name:` is absent.

**Resulting check names**:
- `Local CI / buf-lint`
- `Local CI / unit-test`
- `Local CI / e2e-test`
- `Local CI / integration-test`

### D3: PR trigger types

**Decision**: Use default PR trigger types for the consolidated
workflow (no explicit `types:` filter).

**Rationale**: `unit_test.yml` and `integration_test.yml` already
use the default types (`opened`, `synchronize`, `reopened`).
`e2e_test.yml` restricts to `opened` and `synchronize`, skipping
`reopened`. The only practical difference is an extra e2e run
when a PR is reopened -- an uncommon event with negligible cost.
Normalizing to the default simplifies the workflow and matches
the majority behavior.

### D4: Include `buf-lint` in `ci_local.yml`

**Decision**: Include `buf-lint` as a job in `ci_local.yml`.

**Rationale**: `buf-lint` currently lives in `unit_test.yml` and
runs on every PR. Moving it to `ci_local.yml` means the preflight
will auto-discover it as a release-gated check. This is correct
behavior -- proto lint violations should block a release. Keeping
it in a separate non-`ci_`-prefixed workflow would hide it from
the preflight, adding complexity without benefit.

## Risks / Trade-offs

- **Check name change** -- GitHub status check names will change
  from `unit-test` to `Local CI / unit-test` (and similarly for
  other jobs). Branch protection rules referencing the old names
  will need a one-time admin update.
  Mitigation: Document in the PR description; update rules as
  part of the merge process.

- **buf-lint as release gate** -- `buf-lint` was not previously
  checked by the release preflight. If it has flaky behavior, it
  could block releases.
  Mitigation: `buf lint` is deterministic (no network calls, no
  flaky tests). The proto file (`api/plugin/plugin.proto`) changes
  rarely. Risk is negligible.
