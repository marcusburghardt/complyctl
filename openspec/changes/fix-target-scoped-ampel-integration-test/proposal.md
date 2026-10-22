# Proposal

## Why

The cross-repo integration test assumes the legacy shared Ampel policy-bundle
path. The provider's target-scoped artifact isolation moves the bundle beneath
an indexed generated-content directory, causing the test to fail although
generation and scanning succeed.

## What Changes

- Resolve the Ampel policy bundle through `generate-index.json` when that
  index is available.
- Keep the legacy shared bundle-path assertion only for provider versions that
  do not produce an index.
- Fail the test when a present index lacks the target entry or references a
  missing or empty bundle.
- Preserve the existing policy-content and scan assertions.

## Capabilities

### New Capabilities
- `cross-repo-integration-test`: Verifies the complyctl and provider pipeline
  across legacy and target-scoped Ampel generated-artifact layouts.

### Modified Capabilities
- None.

## Impact

- Affected code: `tests/cross-repo/cross_repo_integration_test.sh`.
- Affected systems: cross-repository CI using current and legacy
  `complytime-providers` binaries.
- No complyctl runtime APIs, dependencies, or user-facing behavior change.
