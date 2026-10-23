# Design

## Context

The current shell integration test asserts a fixed legacy bundle path after
`complyctl generate`. The test workspace config names the Ampel target
`complyctl-repo`. See `proposal.md` for motivation and the delta spec for the
required behavior.

## Goals / Non-Goals

**Goals:**

- Select the generated Ampel bundle using the provider's index when present.
- Keep the assertion compatible with providers that still use the legacy path.
- Preserve the current content and scan assertions after bundle selection.

**Non-Goals:**

- Change complyctl runtime generation, provider behavior, or the provider's
  index format.
- Add a new test framework or alter unrelated OPA assertions.

## Decisions

### Use the index as the layout discriminator

The generation test will check for `generate-index.json` beneath the Ampel
generated-artifact directory. Its presence selects indexed resolution; its
absence selects the existing legacy assertion. This distinguishes an older
provider from a broken new provider without masking provider errors.

The alternative of always falling back to the legacy path would let an invalid
index pass and would not exercise target-scoped isolation.

### Resolve the configured target with jq

The existing script already requires `jq`. The implementation will use it to
extract the entry for `complyctl-repo`, reject an absent or invalid result, and
construct the expected generated-bundle path. The extracted location will be
validated as remaining inside the generated-artifact root before it is used.

The alternative of grepping the index is rejected because it is not robust for
JSON escaping or future field ordering changes.

### Reuse existing assertion helpers

The resolver will produce the bundle path for `assert_file_exists`; the
existing non-empty-file and content assertions will remain the single source
of truth for bundle validation. This retains pass/fail accounting and avoids
duplicating assertion behavior.

## Risks / Trade-offs

- [Provider index schema differs from the issue description] → Confirm the
  schema from complytime-providers#207 before implementing the jq query.
- [An index value escapes the generated-artifact directory] → Reject the path
  before filesystem access.
- [Legacy provider behavior regresses] → Preserve the current assertion only
  when no index is present.

## Migration Plan

No deployment or data migration is required. The updated test will run against
both existing provider binaries and target-scoped provider binaries. Reverting
the single test-script change restores the previous assertion behavior.
