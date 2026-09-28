# Spec Delta

## Purpose

Allow `complyctl scan` to produce multiple secondary report formats
(OSCAL, SARIF, Markdown) in a single invocation, avoiding redundant
scan executions when more than one format is needed.

## ADDED Requirements

### Requirement: Multiple format values accepted

The `--format` / `-f` flag SHALL accept multiple values via
comma-separated syntax (`--format oscal,sarif`) or repeated flags
(`--format oscal --format sarif`). Each requested format SHALL
produce its own report file in the output directory.

#### Scenario: Comma-separated formats

- **WHEN** user runs `complyctl scan --format oscal,sarif`
- **THEN** the system writes an EvaluationLog, an OSCAL report,
  and a SARIF report to the output directory

#### Scenario: Repeated flag formats

- **WHEN** user runs `complyctl scan --format oscal --format sarif`
- **THEN** the system writes an EvaluationLog, an OSCAL report,
  and a SARIF report to the output directory

#### Scenario: Single format backward compatibility

- **WHEN** user runs `complyctl scan --format oscal`
- **THEN** the system writes an EvaluationLog and an OSCAL report,
  behaving identically to the pre-change behavior

#### Scenario: No format specified

- **WHEN** user runs `complyctl scan` without `--format`
- **THEN** the system writes only the EvaluationLog, behaving
  identically to the pre-change behavior

### Requirement: All formats shorthand

The `--format all` value SHALL expand to all available secondary
report formats (oscal, pretty, sarif). The `all` value SHALL NOT
be combined with specific format values.

#### Scenario: All formats requested

- **WHEN** user runs `complyctl scan --format all`
- **THEN** the system writes an EvaluationLog, an OSCAL report,
  a SARIF report, and a Markdown report to the output directory

#### Scenario: All combined with specific format rejected

- **WHEN** user runs `complyctl scan --format all,oscal`
- **THEN** the system exits with an error indicating that `all`
  cannot be combined with specific format values

### Requirement: Duplicate format values rejected

The system SHALL reject duplicate format values with an error
message identifying the duplicated value.

#### Scenario: Duplicate format specified

- **WHEN** user runs `complyctl scan --format oscal,oscal`
- **THEN** the system exits with an error indicating the
  duplicate format value

### Requirement: Partial success on formatter failure

When multiple formats are requested and a secondary formatter
fails, the system SHALL continue producing the remaining
formats. Formatter failures SHALL be reported as warnings to
stderr. The scan exit code SHALL NOT be affected by secondary
formatter failures.

#### Scenario: One formatter fails

- **WHEN** user runs `complyctl scan --format oscal,sarif` and
  the SARIF formatter encounters an error
- **THEN** the system writes the EvaluationLog and the OSCAL
  report, prints a warning about the SARIF failure to stderr,
  and exits with the same code it would without `--format`

### Requirement: Shell completion includes all

Shell completion for the `--format` flag SHALL include `all` in
addition to the existing format values (oscal, pretty, sarif).

#### Scenario: Shell completion values

- **WHEN** user triggers shell completion for `--format`
- **THEN** the completion list includes oscal, pretty, sarif,
  and all
