# Proposal

## Why

The cross-repo integration test (`ci_cross_repo_integration.yml`) fails at the
"Validate EvaluationLog schema" step because `validate.sh` pins the Gemara CUE
schema to `v0.23.0`, which lacks the `evidence` field on `#AssessmentLog`. This
field was added in Gemara `v1.3.0` and is now emitted by `complytime-providers`
after PR #192 merged. All PRs triggering the cross-repo integration test are
blocked until the schema version is updated. Beyond the immediate fix, the
pinned-version approach requires manual bumps and creates surprise failures
when upstream changes land in providers before the schema pin is updated.

Fixes: https://github.com/complytime/complyctl/issues/924

## What Changes

- **Bump Gemara CUE schema pin** in `tests/schema-validation/validate.sh` from
  `v0.23.0` to `v1.6.0` (latest stable) and update the tracking comment
- **Add Renovate custom manager** via a repo-local `renovate.json` that uses a
  regex manager to track the `GEMARA_CUE_MODULE` version in `validate.sh`
  against `gemaraproj/gemara` GitHub releases, so future schema releases
  automatically open update PRs

## Capabilities

### New Capabilities

_None. This change modifies CI test infrastructure and adds dependency
automation configuration. No user-facing capabilities are introduced._

### Modified Capabilities

_None. No spec-level behavior changes. `skip_specs: true` is set._

## Impact

- **CI**: Unblocks all PRs and `main` pushes that trigger the cross-repo
  integration test
- **Tests**: `tests/schema-validation/validate.sh` validates against the
  current Gemara CUE schema instead of a stale pin
- **Dependencies**: New `renovate.json` in repo root adds Renovate custom
  manager config; interacts with the org-wide Renovate runner in
  `org-infra/ci_renovate.yml` which already auto-discovers `complytime/*` repos
- **No code changes**: No Go source, protobuf, or runtime behavior affected
