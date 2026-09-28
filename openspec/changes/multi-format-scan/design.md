# Design

## Context

See proposal.md for motivation. The current `--format` flag is a
`StringVarP` (single string) threaded through five functions in
`cmd/complyctl/cli/scan.go`:

```
scanOptions.format (string)
  -> runScanAndReport(format string, ...)
    -> processScanOutput(format string, ...)
      -> writeScanReports(format string, ...)
        -> writeFormatReport(format string, ...)
```

The dispatch in `writeFormatReport` is a `switch` over the single
format value. The three formatters (OSCAL, SARIF, Markdown) are
independent — they all consume the same `gemara.EvaluationLog` via
`eval.GemaraLog()` and produce separate files. No shared state, no
ordering dependency.

## Goals / Non-Goals

**Goals:**

- Accept multiple `--format` values in one scan invocation
- Add `all` convenience shorthand
- Partial-success semantics for formatter failures
- Full backward compatibility with existing single-format usage

**Non-Goals:**

- Changing output file locations or naming conventions
- Modifying individual formatter behavior or output content
- Adding new output formats (e.g., CSV, HTML)
- Changing EvaluationLog production (always written, unchanged)
- Changing `doctor --format` (mutually exclusive rendering modes,
  semantically different from scan's additive formats)

## Decisions

### D1: Flag primitive — `StringSliceVarP`

**Choice**: Cobra `StringSliceVarP` over `StringArrayVarP`.

**Rationale**: `StringSliceVarP` accepts both comma-separated
(`--format oscal,sarif`) and repeated (`--format oscal --format
sarif`) syntax. `StringArrayVarP` only accepts the repeated form.
Since the valid values are fixed constants that never contain
commas, the comma-splitting behavior is safe and more convenient
for interactive use.

**Alternative considered**: `StringArrayVarP` — rejected because it
is strictly less flexible with no offsetting benefit for this use
case.

### D2: `all` expansion — pre-validation constant list

**Choice**: Expand `all` to `[]string{OutputFormatOSCAL,
OutputFormatPretty, OutputFormatSARIF}` during validation, before
the dispatch loop. Add `OutputFormatAll = "all"` constant to
`internal/complytime/consts.go`.

**Rationale**: Expanding early means the dispatch loop sees only
concrete format values. No special-case logic downstream.
Validation rejects `all` combined with specific values to prevent
redundancy and user confusion.

### D3: Partial-success semantics — warn and continue

**Choice**: When a secondary formatter fails, collect the error
as a warning, print to stderr, and continue with remaining
formats. The scan exit code is determined only by the scan itself
and EvaluationLog write (the critical path), not by secondary
formatter failures.

**Rationale**: Secondary reports are convenience outputs derived
from the EvaluationLog. A failure in one (e.g., an OSCAL
serialization edge case) should not prevent the user from
receiving other reports or the EvaluationLog. This matches the
existing pattern where `reportOperationalWarnings` prints
provider errors to stderr without aborting the scan.

**Alternative considered**: Fail-fast (abort on first formatter
error) — rejected because it would regress single-format
behavior where the user gets no report at all, and because the
EvaluationLog (the authoritative record) is already written
before any formatter runs.

### D4: Duplicate detection — reject at validation

**Choice**: Detect duplicate format values during `validate()`
and return an error. Applied after `all` expansion.

**Rationale**: Duplicates are always user error (typo or
misunderstanding). Producing the same report twice would create
filename collisions (timestamps could differ by milliseconds,
but the intent is clearly wrong). Failing early with a clear
message is better than silent double-write.

## Risks / Trade-offs

- **[Risk] Flag type change is technically breaking for
  programmatic Cobra flag inspection** — Mitigation: No known
  consumers inspect flag types programmatically. The CLI surface
  (accepted values, flag name, shorthand) is unchanged for
  single-value usage.

- **[Risk] `StringSliceVarP` splits on commas unconditionally**
  — Mitigation: All valid format values (`oscal`, `pretty`,
  `sarif`, `all`) are comma-free. No future format value should
  contain a comma. Document this constraint in the constant
  block comment.

- **[Trade-off] Partial-success means the exit code may not
  reflect formatter failures** — Accepted: the EvaluationLog is
  the authoritative scan artifact. Users who need to detect
  formatter failures can check stderr warnings. A dedicated
  `--strict` flag could be added later if needed.
