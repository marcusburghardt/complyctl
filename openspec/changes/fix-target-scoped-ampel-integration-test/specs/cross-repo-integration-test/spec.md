# Spec Delta

## Purpose

Ensures cross-repository Ampel integration coverage remains valid as providers
move generated policy bundles from a shared path to target-scoped artifacts.

## ADDED Requirements

### Requirement: Ampel bundle layout compatibility
The cross-repo integration test SHALL verify the generated Ampel policy bundle
through the target's generated-artifact index when that index exists. When no
index exists, the test SHALL verify the legacy shared bundle path so it remains
compatible with supported provider versions that do not emit an index.

#### Scenario: Provider emits a generated-artifact index
- **WHEN** Ampel generation creates `generate-index.json` for the configured target
- **THEN** the test resolves the target's generated bundle from that index and verifies it is non-empty

#### Scenario: Provider uses the legacy generated-bundle layout
- **WHEN** Ampel generation does not create `generate-index.json`
- **THEN** the test verifies the non-empty bundle at the legacy shared path

### Requirement: Indexed generated artifacts fail closed
The cross-repo integration test SHALL fail its generation assertion when a
present generated-artifact index has no entry for the configured target or
references a missing or empty policy bundle. It SHALL NOT fall back to the
legacy shared path after detecting an index.

#### Scenario: Target entry is absent from a present index
- **WHEN** `generate-index.json` exists but does not identify a generated bundle for the configured target
- **THEN** the generation assertion fails

#### Scenario: Indexed bundle is unavailable
- **WHEN** `generate-index.json` identifies a bundle that is missing or empty
- **THEN** the generation assertion fails

### Requirement: Existing pipeline assertions remain intact
The cross-repo integration test SHALL continue to verify that the generated
bundle includes the `block-force-push` policy and that the subsequent Ampel
scan produces the existing attestation assertions.

#### Scenario: Target-scoped bundle supports the existing pipeline
- **WHEN** the indexed Ampel bundle is present and valid
- **THEN** the test verifies its policy content and executes the existing scan assertions
