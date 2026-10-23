# Proposal

## Why

After adopting the org-infra reusable release workflows (#699),
complyctl must pass a hardcoded `ci_checks` override to the release
preflight because its test jobs are split across three separate
workflow files (`unit_test.yml`, `e2e_test.yml`,
`integration_test.yml`) instead of living in the conventional
`ci_local.yml` that the preflight auto-discovers via `yq`. This
creates a maintenance risk: renaming a job or adding a test workflow
requires a manual update to `release.yml`.

Ref: https://github.com/complytime/complyctl/issues/723

## What Changes

- **Create** `.github/workflows/ci_local.yml` combining all three
  test workflows into a single workflow with four jobs: `buf-lint`,
  `unit-test`, `e2e-test`, `integration-test`
- **Remove** `.github/workflows/unit_test.yml`,
  `.github/workflows/e2e_test.yml`,
  `.github/workflows/integration_test.yml`
- **Remove** the `ci_checks` override from `.github/workflows/release.yml`
  (line 52), enabling preflight auto-discovery
- **Update** `AGENTS.md` CI Workflow Structure table to reflect the
  consolidated workflow

## Capabilities

### New Capabilities

None. This is a CI workflow consolidation with no new system
behavior.

### Modified Capabilities

None. No spec-level behavior changes. `skip_specs: true` is set
in `.openspec.yaml` since this is a pure tooling change.

## Impact

- **Workflows**: Three workflow files replaced by one; one override
  removed from `release.yml`
- **GitHub status checks**: Check names will change from bare job
  names (e.g., `unit-test`) to prefixed names
  (e.g., `Local CI / unit-test`). Branch protection rules
  referencing the old names will need updating.
- **Release gate**: `buf-lint` becomes a release-gated check via
  auto-discovery (previously ran in CI but was not checked by the
  release preflight)
- **Dependencies**: None
- **Code**: None -- no source code changes
