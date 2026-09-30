# Proposal

## Why

`complyctl scan --format` accepts only a single secondary report format
(oscal, pretty, or sarif). When a user needs more than one format --
for example, both OSCAL for compliance tooling and SARIF for GitHub
Code Scanning -- they must run the scan twice. This is wasteful since
the scan itself and EvaluationLog are identical between runs; only the
output serialization differs. The formatters are already independent
consumers of the same `gemara.EvaluationLog` object, so the
single-format restriction is an artificial limitation of the flag type,
not an architectural constraint.

Upstream: complytime/complyctl#890

## What Changes

- `--format` / `-f` flag changes from `StringVarP` (single value) to
  `StringSliceVarP` (multi-value). Accepts comma-separated values
  (`--format oscal,sarif`) or repeated flags
  (`--format oscal --format sarif`).
- `--format all` convenience value expands to all three formats
  (oscal, pretty, sarif). Cannot be combined with specific formats.
- Validation updated to iterate over the slice and reject unknown or
  duplicate values.
- `writeFormatReport` dispatch changes from a single-value switch to
  a loop over requested formats with partial-success semantics:
  if one formatter fails, the others still run and a warning is
  emitted to stderr.
- Existing single-format usage (`--format oscal`) remains fully
  backward compatible.
- No changes to `internal/output/` formatters, proto definitions,
  constants, or provider interface.

## Capabilities

### New Capabilities

- `multi-format-scan`: Support for requesting multiple secondary
  report formats in a single `complyctl scan` invocation, including
  the `all` convenience shorthand and partial-success error handling.

### Modified Capabilities

(none -- existing formatter behavior and EvaluationLog production
are unchanged)

## Impact

- **Code**: `cmd/complyctl/cli/scan.go` -- flag type, validation,
  call chain signatures (`string` to `[]string`), dispatch loop.
  Estimated ~80-100 lines changed.
- **APIs**: No proto, gRPC, or provider interface changes.
- **Dependencies**: No new dependencies.
- **Tests**: Unit tests for multi-format validation, `all` expansion,
  duplicate rejection, and partial-success semantics. E2E test
  coverage for multi-format invocation.
- **Documentation**: `--format` flag help text update. Quickstart
  and user docs if they reference `--format` usage examples.
